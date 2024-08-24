package model

import "time"

type Song struct {
	UUID     string
	Title    string
	Duration time.Duration
}
