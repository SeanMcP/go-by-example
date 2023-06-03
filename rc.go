package main

import (
	"fmt"
	"os/exec"
	"regexp"
)

// Run changed
func main() {
	cmd := exec.Command("git", "status", "-s")
	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	out := string(stdout)

	re := regexp.MustCompile(`([\w-]+).go`)

	matches := re.FindAllStringSubmatch(out, 1)

	if len(matches) > 0 {
		file := matches[0][0]

		if file == "rc.go" {
			fmt.Println("No recursive runs")
			return
		}

		fmt.Println("Running", file)

		cmd := exec.Command("go", "run", file)
		stdout, err := cmd.Output()

		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println(string(stdout))
	} else {
		fmt.Println("No files to run")
	}
}
