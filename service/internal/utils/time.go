package utils

import (
	"log"
	"time"
)

var Loc *time.Location

func init() {
	var err error
	Loc, err = time.LoadLocation("Asia/Taipei")
	if err != nil {
		log.Fatalf("failed to load location: %v", err)
	}
}
