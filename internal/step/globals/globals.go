package globals

import (
	"github.com/szumel/rusher/internal/platform/schema"
	"regexp"
	"strings"
)

const (
	wrapLeft  = "{"
	wrapRight = "}"

	varPattern = wrapLeft + ".+" + wrapRight
)

func IsGlobal(p string) (bool, error) {
	regex, err := regexp.Compile(varPattern)
	if err != nil {
		return false, err
	}

	return regex.MatchString(p), nil
}

func Parse(globals []schema.Var, p string) string {
	var val string
	for _, g := range globals {
		if strings.Contains(p, wrap(g.Name)) {
			val = strings.Replace(p, wrap(g.Name), g.Value, 1)
		}
	}

	return val
}

func wrap(s string) string {
	return wrapLeft + s + wrapRight
}
