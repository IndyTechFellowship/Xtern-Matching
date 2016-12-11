package tests

import (
	"Xtern-Matching/handlers/services"
	"os"
	"testing"

	"google.golang.org/appengine/aetest"
)

func TestStudentExport(t *testing.T) { //&aetest.Options{StronglyConsistentDatastore: true}
	ctx, done, err := aetest.NewContext()
	if err != nil {
		t.Fatal(err)
	}
	defer done()
	createStudent(ctx)
	createStudent(ctx)
	export, err := services.ExportStudents(ctx)
	if err != nil {
		t.Fatal(err)
	}
	f, err := os.Create("export.csv")
	f.WriteString(export)
	f.Sync()
	//	println(export)
}
