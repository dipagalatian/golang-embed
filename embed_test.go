package golangembed

import (
	_ "embed"
	"fmt"
	"io/fs"
	"os"
	"testing"
)

//go:embed version.txt
var version string

//go:embed version.txt
var version2 string

func TestEmbedString(t *testing.T) {

	fmt.Println("Version:", version)
	fmt.Println("Version2:", version2)
	
}

//go:embed cat_image.jpg
var cat_image []byte 
func TestEmbedByte(t *testing.T) {

	err := os.WriteFile("copy_cat_image.jpg", cat_image, fs.ModePerm)
	if err != nil {
		panic(err)
	}
	
}