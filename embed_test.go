package golangembed

import (
	_ "embed"
	"fmt"
	"testing"
)

//go:embed version.txt
var version string

//go:embed version.txt
var version2 string

func TestEmbed(t *testing.T) {

	fmt.Println("Version:", version)
	fmt.Println("Version2:", version2)
	
}