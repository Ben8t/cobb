package main

import (
	"cobb"
	"fmt"
)

func main() {
	archive := cobb.Archive{Camera: "Canon A1", Roll: "Portra 400", Date: "202211"}
	fmt.Println(archive.MakeArchiveName())
}
