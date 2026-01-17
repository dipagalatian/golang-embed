package golangembed

import (
	"embed"
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

//go:embed files/a.txt
//go:embed files/b.txt
//go:embed files/c.txt
var files embed.FS
func TestEmbedMultipleFiles(t *testing.T) {

	dataA, err := files.ReadFile("files/a.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println("dataA:", string(dataA))

	dataB, err := files.ReadFile("files/b.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println("dataB:", string(dataB))

	dataC, err := files.ReadFile("files/c.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println("dataC:", string(dataC))

}