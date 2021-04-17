package test

import (
	"testing"

	"github.com/itscharlieliu/golang-find/pkg"
)

func TestFindNoArgs(t *testing.T) {
	commandOutputString := pkg.RunCmd("go", []string{"run", "../cmd/find.go", "."})

	realOutputString := pkg.RunCmd("find", []string{"."})

	if commandOutputString != realOutputString {
		t.Errorf("Expected %s, got %s", realOutputString, commandOutputString)
	}
}
