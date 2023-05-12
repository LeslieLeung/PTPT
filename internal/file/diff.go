package file

import (
	log "github.com/sirupsen/logrus"
	"os/exec"
	"strings"
)

func DiffFiles(args []string) ([]string, error) {
	cmd := append([]string{"diff", "--find-renames", "--name-only"}, args...)
	log.Debugf("cmd: %v", cmd)
	bytes, err := exec.Command("git", cmd...).Output()
	if err != nil {
		return nil, err
	}
	return strings.Split(strings.Trim(string(bytes), "\n"), "\n"), nil
}
