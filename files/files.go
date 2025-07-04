package main

import (
	"fmt"
	"os"
)

func main() {
	// result, err := os.Open("example.txt")

	// if err != nil {
	// 	// log the error
	// 	// panic mode
	// 	panic(err)
	// }

	// fileInfo, err := result.Stat()

	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println("file name:", fileInfo.Name())
	// fmt.Println("file or folder", fileInfo.IsDir())
	// fmt.Println("file size: ", fileInfo.Size())
	// fmt.Println("file permission", fileInfo.Mode())
	// fmt.Println("file modified at: ", fileInfo.ModTime())

	// //read the contents of the file
	// defer result.Close()

	// buf := make([]byte, fileInfo.Size())

	// d, error := result.Read(buf)

	// if error != nil {
	// 	panic(err)
	// }

	// for i := 0; i < len(buf); i++ {
	// 	println("data", d, string(buf[i]))
	// }

	// data, err := os.ReadFile("example.txt")

	// if err != nil {
	// 	panic(err)
	// }

	// println(string(data))

	// read folders
	// dir, err := os.Open("../")
	// if err != nil {
	// 	panic(err)
	// }

	// defer dir.Close()

	// fileInfo, err := dir.ReadDir(-1)

	// for _, fi := range fileInfo {
	// 	fmt.Println(fi.Name(), fi.IsDir())
	// }

	//creating file
	// file, err := os.Create("example2.txt")

	// if err != nil {
	// 	panic(err)
	// }

	// defer file.Close()

	// // file.WriteString("hi go")
	// // how can we replace?

	// bytes := []byte("Hello Golang")
	// file.Write(bytes)

	// read and write to another file (streaming fashion)

	// sourceFile, err := os.Open("example.txt")
	// if err != nil {
	// 	panic(err)
	// }

	// defer sourceFile.Close()

	// destFile, err := os.Create("example2.txt")
	// if err != nil {
	// 	panic(err)
	// }

	// defer destFile.Close()

	// reader := bufio.NewReader(sourceFile)
	// writer := bufio.NewWriter(destFile)

	// for {
	// 	b, err := reader.ReadByte()
	// 	if err != nil {
	// 		if err.Error() != "EOF" {
	// 			panic(err)
	// 		}
	// 		break
	// 	}
	// 	e := writer.WriteByte(b)
	// 	if err != nil {
	// 		panic(e)
	// 	}
	// }
	// writer.Flush()
	// fmt.Println("Written to new file successfully")

	//delete a file
	sourceFile, err := os.Open("example2.txt")

	if err != nil {
		panic(err)
	}

	defer sourceFile.Close()

	e := os.Remove(sourceFile.Name())
	if e != nil {
		panic(err)
	}
	fmt.Println("File deleted successfully")
}
