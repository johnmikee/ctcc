package tcc

import (
	"os"
	"os/exec"
	"strings"
)

type Users struct {
	Name string `json:"name"`
	DB   string `json:"home"`
}

// ListUsers returns a list of users on the system
func ListUsers() []Users {
	users := listUsers()
	var userHomes []Users

	for _, user := range users {
		tccDBPath := user + "/Library/Application Support/com.apple.tcc/tcc.db"

		_, err := os.Stat(tccDBPath)

		if err == nil {
			userHomes = append(userHomes, Users{Name: removePath(user), DB: tccDBPath})
		}
	}

	return userHomes
}

func listUsers() []string {
	cmd := exec.Command("dscl", "/Local/Default", "-list", "/Users", "NFSHomeDirectory")
	output, err := cmd.Output()
	if err != nil {
		return nil
	}

	lines := strings.Split(string(output), "\n")

	var userHomes []string
	for _, line := range lines {
		if strings.Contains(line, "/var/empty") {
			continue
		}
		fields := strings.Fields(line)
		if len(fields) == 2 && strings.HasPrefix(fields[1], "/") {
			userHomes = append(userHomes, fields[1])
		}
	}

	return userHomes
}

func removePath(user string) string {
	prefix := "/Users/"

	if !strings.HasPrefix(user, prefix) {
		// The prefix doesn't exist in the input string
		return user
	}

	result := strings.Replace(user, prefix, "", 1)
	return result
}
