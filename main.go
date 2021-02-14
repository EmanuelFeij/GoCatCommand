package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}

// ReadFile reads a file and prints it to stdout.
func ReadFile(filename string) (string, error) {

	data, err := ioutil.ReadFile(filename)
	checkError(err)

	str := fmt.Sprintf("Filename --> %v \n %v", filename, string(data))
	return str, nil
}

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		return
	}
	for _, file := range args {
		str, err := ReadFile(file)
		checkError(err)
		fmt.Println(str)
	}

}
