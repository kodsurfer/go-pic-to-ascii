package main

import (
	"fmt"
	"image"
	"os"
)

func main() {

	picPath := os.Args[1]
	file, err := os.Open(picPath)
	if err != nil {
		return
	}
	defer file.Close()

	pic, _, err := image.Decode(file)
	if err != nil {
		return
	}

	fmt.Println(pic.Bounds().Dx())
	fmt.Println(pic.Bounds().Dy())
}
