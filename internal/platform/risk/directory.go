package risk

import (
	"errors"
)

type wholeTree bool

var riskyDirs = map[string]wholeTree{
	"/":      false,
	"/bin":   true,
	"/boot":  true,
	"/cdrom": true,
	"/dev":   true,
	"/etc":   true,
	"/home":  false,
	"/lib":   true,
	"/lib64": true,
	"/media": true,
	"/mnt":   false,
	"/opt":   true,
	"/proc":  true,
	"/root":  true,
	"/run":   true,
	"/sbin":  true,
	"/snap":  true,
	"/srv":   true,
	"/sys":   true,
	"/usr":   true,
	"/var":   false,
}

var (
	ErrRiskyDir = errors.New("risky dir detected")
)

//DirCheck detects if given directories are treated as risky
func DirCheck(directories ...string) error {
	for _, dir := range directories {
		for riskyDir, wholeTree := range riskyDirs {
			if wholeTree == true && len(dir) >= len(riskyDir) {
				if dir[:len(riskyDir)] == riskyDir {
					return ErrRiskyDir
				}
			} else if riskyDir == dir {
				return ErrRiskyDir
			}
		}
	}

	return nil
}
