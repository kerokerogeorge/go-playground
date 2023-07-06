package handler

import (
	"os"
)

func CreateFile() {
	os.Create("test.pdf")
}
