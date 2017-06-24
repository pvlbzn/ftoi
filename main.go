package main

import (
	"image"
	"image/png"
	"io/ioutil"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		panic("no file path provided")
	}

	path := os.Args[1]
	data := readFile(path)
	img := createImage(&data)
	writeImage(img)
}

func readFile(path string) []byte {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return data
}

func createImage(data *[]byte) *image.Gray {
	img := image.NewGray(image.Rect(0, 0, 120, len(*data)/120))
	img.Pix = *data
	return img
}

func writeImage(img *image.Gray) {
	w, err := os.Create("out.png")
	if err != nil {
		panic("write error")
	}
	defer w.Close()
	png.Encode(w, img)
}
