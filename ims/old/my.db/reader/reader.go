// package main

// import (
// 	"fmt"
// 	"io/ioutil"
// )

// func main() {
// 	filePath := "/home/user/my.db/reader/default.etcd/member/snap/db"

// 	content, err := ioutil.ReadFile(filePath)
// 	if err != nil {
// 		fmt.Printf("Error reading the file: %v", err)
// 		return
// 	}

// 	fmt.Print(string(content))
// }