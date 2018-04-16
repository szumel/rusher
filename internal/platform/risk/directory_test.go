package risk

import "testing"

func TestDirCheck(t *testing.T) {
	cases := map[string]error{
		"/":          ErrRiskyDir,
		"/home/test": nil,
		"/etc":       ErrRiskyDir,
		"/etc/xxx":   ErrRiskyDir,
		"/root":      ErrRiskyDir,
		"/var":       ErrRiskyDir,
		"/var/www":   nil,
		"/bin":       ErrRiskyDir,
	}

	for path, expected := range cases {
		err := DirCheck(path)
		if err != expected {
			t.Fatalf("Checking risky directories failed. Expected %T have %T.", expected, err)
		}
	}
}
