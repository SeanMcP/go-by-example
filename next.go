package main

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

// TODO: add dryrun flag
func main() {
	data, err := os.ReadFile("README.md")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	reg := regexp.MustCompile(`(\d+). \[ \] (.+)\n`)
	result := reg.FindStringSubmatch(string(data))

	fmt.Println(result[2]+":", result[1]+"/79")

	example := strings.ToLower(result[2])
	slug := strings.ReplaceAll(example, " ", "-")
	exec.Command("open", "https://gobyexample.com/"+slug).Run()

	filename := slug + ".go"
	filepath := "examples/" + filename

	os.WriteFile(filepath, []byte("// "+filename+"\n"), 0644)

	exec.Command("code", filepath+":2", "-g").Run()

	readme := strings.Replace(string(data), "[ ]", "[x]", 1)
	os.WriteFile("README.md", []byte(readme), 0644)
}
