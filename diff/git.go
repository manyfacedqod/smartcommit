package diff

import (
	"os/exec"
)

func GetStagedDiff(options []string) (string, error) {
	args := append(options, "diff", "--cached");
	out, err := exec.Command("git", args...).Output()
	if err != nil {
		return "", err
	}
	return string(out), nil
}
