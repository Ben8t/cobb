package cobb

type Archive struct {
	Camera string
	Roll   string
	Date   string
}

func (archive Archive) MakeArchiveName() string {
	return archive.Date + "_" + archive.Roll + "_" + archive.Camera
}

func main() {

}
