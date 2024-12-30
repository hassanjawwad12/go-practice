package main

import (
	"fmt"
	"os"
)

func main() {

	//returns 2 things , file obj and error
	f, err := os.Open("example.txt")

	if err != nil {
		panic(err)
	}

	fileInfo, err := f.Stat()
	if err != nil {
		panic(err)
	}
	fmt.Println("file name is:", fileInfo.Name())
	fmt.Println("file being a directory:", fileInfo.IsDir())
	fmt.Println("file size is:", fileInfo.Size()) //shows the size in bytes
	fmt.Println("file permission is:", fileInfo.Mode())
	fmt.Println("file modified at:", fileInfo.ModTime())

	//read file
	defer f.Close()
	buf := make([]byte, fileInfo.Size())
	d, err := f.Read(buf) //the file will be read in the buffer
	if err != nil {
		panic(err)
	}
	println("number of bytes read:", d)
	for i := 0; i < len(buf); i++ {
		fmt.Println("data", string(buf[i]))
	}

	//this is not recommended because it read the whole file at once in memeory
	//it is not viable for big files such as big files in GB
	a, err := os.ReadFile("example.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(a))

	//reading the directories
	dir, err := os.Open(".") //root folder
	if err != nil {
		panic(err)
	}
	defer dir.Close()

	fileInformation, err := dir.ReadDir(-1)
	fmt.Println("The directories in the folder are")
	for _, i := range fileInformation {
		fmt.Println(i.Name(), i.IsDir())
	}

	//create a file
	b, err := os.Create("example2.txt")
	if err != nil {
		panic(err)
	}
	defer b.Close()
	b.WriteString("hi go") //appends the file

	//alternative
	bytes := []byte(" Go-lang")
	b.Write(bytes)

}
