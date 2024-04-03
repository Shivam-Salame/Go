package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("working with files")

	writeFile()
	readFile("./mytxtfile.txt")
}

func writeFile() {
	content := "This needs to go in a file"

	file, erro := os.Create("./mytxtfile.txt")
	checkNilError(erro)

	length, errr := io.WriteString(file, content)
	checkNilError(errr)
	fmt.Println("length is:", length)

	defer file.Close()
}

func readFile(filename string) {
	databyte, err := os.ReadFile(filename)
	checkNilError(err)

	fmt.Println("text data inside the file is\n", string(databyte))
}

func checkNilError(err error)  {
	if err != nil {
		panic(err)
	}
}