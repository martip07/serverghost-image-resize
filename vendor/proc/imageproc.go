package main

import (
	"fmt"
	"os"

	"gopkg.in/h2non/bimg.v1"
)

func main() {
	fmt.Println("Howdy demo")
	buffer, err := bimg.Read("image.png")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	var ImageOptions = bimg.Options{
		Width:   400,
		Height:  300,
		Force:   true,
		Type:    bimg.ImageType(3),
		Quality: 20,
	}
	newImage, err := bimg.NewImage(buffer).Process(ImageOptions)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	bimg.Write("newimage1.png", newImage)
}
