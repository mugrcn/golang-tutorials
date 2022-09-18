package main

import (
	"fmt"
	"path"
)

func pathSplit() {
	_, file := path.Split("css/main.css")
	fmt.Println("file", file)
}
