package tests

import (
	"testing"
	"google.golang.org/appengine/aetest"
)

func TestPost(t *testing.T) {
	ctx, done, err := aetest.NewContext()
	if err != nil {
		t.Fatal(err)
	}
	defer done()
}

func TestGet(t *testing.T) {
	ctx, done, err := aetest.NewContext()
	if err != nil {
		t.Fatal(err)
	}
	defer done()
}
