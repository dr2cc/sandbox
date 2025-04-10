package main

import (
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	r := strings.NewReader("данные для чтения")

	if _, err := io.Copy(os.Stdout, r); err != nil {
		log.Fatal(err)
	}
}
