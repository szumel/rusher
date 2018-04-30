//package step provides tools for steps executing
package step

import (
	"fmt"

	"github.com/szumel/rusher/internal/platform/schema"
	"github.com/szumel/rusher/internal/step/globals"
	"github.com/szumel/rusher/internal/step/macro"
)

const (
	typeMacro = "macro"
	typeStep  = "step"
)

//Steps pool holds all available steps
var StepsPool = &Pool{Steps: []Step{
	&ChangeCwd{},
	&ComposerInstall{},
	&copyFiles{},
	&GitClone{},
	&Magento2SetupUpgrade{},
	&Magento2EnableModules{},
	&Magento2Compile{},
	&Move{},
	&PrintPwd{},
	&PrintString{},
	&RemoveDir{},
	&Symlink{},
	&openLink{},
}}

type ErrInvalidStep struct{}

func (e *ErrInvalidStep) Error() string {
	return "invalid step code. You can ensure that step code is correct by calling listSteps command."
}

func Rush(c *schema.Config, env string) error {
	r := rusher{pool: StepsPool, config: c}

	return r.rush()
}

type rusher struct {
	pool   *Pool
	config *schema.Config
}

func (r *rusher) rush() error {
	scs, err := r.extract()
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

func (r *rusher) extract() ([]stepCtx, error) {
	var scs []stepCtx
	for _, sequenceElem := range r.config.Sequence.SequenceElems {
		switch sequenceElem.XMLName.Local {
		case typeMacro:
			mscs, err := r.macroScs(scs, toMacro(sequenceElem))
			scs = mscs
			if err != nil {
				return nil, err
			}
			break
		case typeStep:
			sc, err := r.stepSc(sequenceElem)
			scs = append(scs, sc)
			if err != nil {
				return nil, err
			}
			break
		}
	}

	return scs, nil
}

func (r *rusher) stepSc(step schema.SequenceElem) (stepCtx, error) {
	var sc stepCtx
	for _, stepCmd := range r.pool.Steps {
		if stepCmd.Code() == step.Code {
			ctx, err := r.createContext(step, r.config)
			if err != nil {
				return stepCtx{}, err
			}
			sc = stepCtx{step: stepCmd, ctx: ctx}
			break
		}
	}

	return sc, nil
}

func (r *rusher) macroScs(scs []stepCtx, macroElem schema.MacroElem) ([]stepCtx, error) {
	loader := macro.CreateLoader()
	macroSchema, err := loader.Load(macroElem.Source)
	if err != nil {
		return nil, err
	}

	m, err := macro.Create(macroSchema)
	if err != nil {
		return nil, err
	}

	for _, step := range m.Steps {
		for _, stepCmd := range r.pool.Steps {
			if stepCmd.Code() == step.Code {
				ctx, err := r.createContext(schema.SequenceElem{Code: step.Code, Params: step.Params}, r.config)
				if err != nil {
					return nil, err
				}
				scs = append(scs, stepCtx{step: stepCmd, ctx: ctx})
				continue
			}
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

//@todo schema.SequenceElem change to schema.StepElem
func (r *rusher) createContext(step schema.SequenceElem, config *schema.Config) (Context, error) {
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

func toMacro(seqElem schema.SequenceElem) schema.MacroElem {
	var m schema.MacroElem
	m.Params = seqElem.Params
	for _, param := range m.Params {
		if param.Name.Local == "v" {
			m.Version = param.Value
			continue
		}

		if param.Name.Local == "source" {
			m.Source = param.Value
			continue
		}
	}

	return m
}
