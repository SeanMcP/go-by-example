// embed-directive.go
package main

import "embed"

//go:embed temp/single_file.txt
var fileString string

//go:embed temp/single_file.txt
var fileByte []byte

//go:embed temp/single_file.txt
//go:embed temp/*.hash
var folder embed.FS

func main() {
	print(fileString)
	print(string(fileByte))

	content1, _ := folder.ReadFile("temp/a.hash")
	print(string(content1))

	content2, _ := folder.ReadFile("temp/b.hash")
	print(string(content2))
}

/**

RUN THIS FIRST:

mkdir -p temp
echo "hello go:embed" > temp/single_file.txt
echo "123" > temp/a.hash
echo "456" > temp/b.hash

*/
