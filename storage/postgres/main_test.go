package postgres

import (
	"fmt"
	"os"
	"testing"
)

func TestMain(t *testing.M) {
	extCode := t.Run()
	fmt.Println("ext Code ", extCode)
	os.Exit(extCode)
}