package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"
)

func getPath(childPath string) string {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("No caller information")
	}

	path := filepath.Join(filename)
	dir := filepath.Dir(path)

	return filepath.Join(dir, childPath)
}

func main() {
	dryRunPtr := flag.Bool("dryrun", false, "dry run")
	verbosePtr := flag.Bool("verbose", false, "verbose")

	flag.Parse()

	runCheck := func(log string, fn func()) {
		if *dryRunPtr {
			fmt.Println(log)
		} else {
			fn()
		}
	}

	verboseCheck := func(log string) {
		if *verbosePtr {
			fmt.Println(log)
		}
	}

	readmePath := getPath("README.md")

	data, err := os.ReadFile(readmePath)
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	reg := regexp.MustCompile(`(\d+). \[ \] (.+)\n`)
	result := reg.FindStringSubmatch(string(data))

	fmt.Println(result[2]+":", result[1]+"/79")

	example := strings.ToLower(result[2])
	slug := strings.ReplaceAll(example, "/ ", "")
	slug = strings.ReplaceAll(slug, " ", "-")
	url := "https://gobyexample.com/" + slug

	runCheck("opening "+url, func() {
		exec.Command("open", "https://gobyexample.com/"+slug).Run()
	})

	filename := slug + ".go"
	filepath := getPath("examples/" + filename)
	contents := "// " + filename + "\n"

	runCheck("creating "+filepath, func() {
		os.WriteFile(filepath, []byte(contents), 0644)
	})

	verboseCheck(contents)

	runCheck("opening "+filepath, func() {
		exec.Command("code", filepath+":2", "-g").Run()
	})

	readme := strings.Replace(string(data), "[ ]", "[x]", 1)

	runCheck("updating "+readmePath, func() {
		os.WriteFile(readmePath, []byte(readme), 0644)
	})

	verboseCheck(readme)
}
