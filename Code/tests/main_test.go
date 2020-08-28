package tests

import (
	"os"
	"testing"

	"gitlab.iut-clermont.uca.fr/esriat/gestion-tps-projet/Code/globals"
)

func TestMain(m *testing.M) {
	globals.Init()
	os.Exit(m.Run())
}
