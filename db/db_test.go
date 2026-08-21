package db

import (
	"fmt"
	"testing"
)

func TestInitDB(t *testing.T) {
	db, got, err := InitDB(false, false)
	want := "connection established succesfully"
	fmt.Println(db, err)
	if got[0] != want {
		t.Errorf("Add(2, 3) = %d; want %d", got, want)
	}
}
