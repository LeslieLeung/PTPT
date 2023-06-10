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

func DiffFileContents(args []string) (string, error) {
	cmd := append([]string{"diff", "--find-renames"}, args...)
	log.Debugf("cmd: %v", cmd)
	bytes, err := exec.Command("git", cmd...).Output()
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func DiffStageAndHead() (string, error) {
	cmd := []string{"diff", "--cached"}
	log.Debugf("cmd: %v", cmd)
	bytes, err := exec.Command("git", cmd...).Output()
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func Commit(message string) error {
	cmd := []string{"commit", "-m", message}
	log.Debugf("cmd: %v", cmd)
	bytes, err := exec.Command("git", cmd...).Output()
	if err != nil {
		return err
	}
	log.Debugf("output: %v", string(bytes))
	return nil
}
