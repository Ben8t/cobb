package main

import "testing"

func TestMakeArchiveName(t *testing.T) {
	archive := Archive{"Canon A1", "Portra 400", "202212"}
	//want := "202212_Portra 400_Canon A1"
	want := "202212_portra400_canonA1"
	msg := archive.MakeArchiveName()
	if want != msg {
		t.Fatalf(`archive.MakeArchiveName() = %q, want match for %#q, nil`, msg, want)
	}
}
