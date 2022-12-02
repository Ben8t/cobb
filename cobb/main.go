package main

import (
	"fmt"
	"strings"
)

type Archive struct {
	Camera string
	Roll   string
	Date   string
}

func ProcessStringField(text string) string {
	texts := strings.Fields(text)
	texts[0] = strings.ToLower(texts[0])
	processed_text := strings.Join(texts, "")
	return processed_text
}

func (archive Archive) MakeArchiveName() string {
	return archive.Date + "_" + ProcessStringField((archive.Roll)) + "_" + ProcessStringField((archive.Camera))
}

func main() {
	archive := Archive{"Canon A1", "Portra 400", "202211"}
	fmt.Println(archive.MakeArchiveName())
}
