package tests

import (
	"Xtern-Matching/handlers/services"
	"os"
	"testing"
	"time"

	"google.golang.org/appengine/aetest"
)

/*
	Test needs to be verified manually in a csv viewer
	Will convert to assert framework once backend refactoring is pushed.
*/
func TestStudentExport(t *testing.T) {
	t.Parallel()
	ctx, done, err := aetest.NewContext()
	if err != nil {
		t.Fatal(err)
	}
	defer done()
	for i := 0; i < 50; i++ {
		_, _, err = createStudent(ctx)
		if err != nil {
			t.Fatal(err)
		}
	}
	time.Sleep(time.Millisecond * 500)
	export, err := services.ExportStudents(ctx)
	if err != nil {
		t.Fatal(err)
	}
	f, err := os.Create("manual/export.csv")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	f.Write(export)
	f.Sync()
}
