package main

import (
	"os"

	_ "github.com/RevoKey/beats-httpout"

	"github.com/elastic/beats/filebeat/beater"
	"github.com/elastic/beats/libbeat/beat"
)

func main() {
	err := beat.Run("filebeat", "", beater.New)
	if err != nil {
		os.Exit(1)
	}
}
