//package step provides tools for steps executing
package step

import (
	"fmt"
	"github.com/szumel/rusher/internal/platform/container"
	"github.com/szumel/rusher/internal/platform/schema"
	"github.com/szumel/rusher/internal/step/globals"
)

const AliasPool = "step.Pool"

type ErrInvalidStep struct{}

func (e *ErrInvalidStep) Error() string {
	return "invalid step code. You can ensure that step code is correct by calling listSteps command."
}

func init() {
	pool := Pool{Steps: []Step{}}
	container.Set(AliasPool, &pool)
}

func Rush(c *schema.Config, env string) error {
	pool, err := container.Get(AliasPool)
	if err != nil {
		return err
	}

	r := rusher{pool: pool.(*Pool), config: c}

	return r.rush(c, env)
}

type rusher struct {
	pool   *Pool
	config *schema.Config
}

func (r *rusher) rush(s *schema.Config, env string) error {
	scs, err := r.extractSteps()
	if err != nil {
		return err
	}

	err = r.validate(scs)
	if err != nil {
		return err
	}

	err = r.execute(scs)
	if err != nil {
		return err
	}

	return nil
}

//Step represents one command which will be executed during deploy
type Step interface {
	Execute(ctx Context) error
	Code() string
	Name() string
	Description() string
	Params() map[string]string
	Validate(ctx Context) error
}

//Pool holds all steps in actual context
type Pool struct {
	Steps []Step
}

//Register adds new step to steps pool
func (p *Pool) Register(step Step) error {
	p.Steps = append(p.Steps, step)

	return nil
}

func Register(step Step) error {
	pool, err := container.Get(AliasPool)
	if err != nil {
		return err
	}

	err = pool.(*Pool).Register(step)
	if err != nil {
		return err
	}

	return nil
}

func (r *rusher) extractSteps() ([]stepCtx, error) {
	scs := []stepCtx{}
	for _, configStep := range r.config.Steps {
		valid := false
		for _, step := range r.pool.Steps {
			if step.Code() == configStep.Code {
				ctx, err := r.createContext(configStep, r.config)
				if err != nil {
					return scs, err
				}
				stepCtx := stepCtx{step: step, ctx: ctx}
				scs = append(scs, stepCtx)
				valid = true
				continue
			}
		}

		if valid == false {
			return scs, &ErrInvalidStep{}
		}
	}

	return scs, nil
}

type stepCtx struct {
	step Step
	ctx  Context
}

func (r *rusher) validate(scs []stepCtx) error {
	for _, sc := range scs {
		err := sc.step.Validate(sc.ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *rusher) execute(scs []stepCtx) error {
	for _, sc := range scs {
		err := sc.step.Execute(sc.ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

type Context interface {
	ProjectPath() string
	Params() map[string]string
	Globals() map[string]string
}

func (r *rusher) createContext(step schema.Step, config *schema.Config) (Context, error) {
	ctx := ContextImpl{params: map[string]string{}, globals: map[string]string{}}
	ctx.projectPath = r.config.ProjectPath

	for _, param := range step.Params {
		isG, err := globals.IsGlobal(param.Value)
		if err != nil {
			return &ContextImpl{}, err
		}

		if isG {
			parsed := globals.Parse(config.Globals, param.Value)
			if parsed == "" {
				return &ContextImpl{}, NewError(step.Code, fmt.Sprintf("Global %s is required", param.Value))
			}
			ctx.params[param.Name.Local] = parsed
		} else {
			ctx.params[param.Name.Local] = param.Value
		}
	}

	for _, global := range config.Globals {
		ctx.globals[global.Name] = global.Value
	}

	return &ctx, nil
}

type ContextImpl struct {
	projectPath string
	params      map[string]string
	globals     map[string]string
}

func (c *ContextImpl) Globals() map[string]string {
	return c.globals
}

func (c *ContextImpl) ProjectPath() string {
	return c.projectPath
}

func (c *ContextImpl) Params() map[string]string {
	return c.params
}
