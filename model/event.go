package model

import "time"

type Event struct {
	Id   int64     `json:"id,omitempty"`
	Name string    `json:"name,omitempty"`
	Date time.Time `json:"date,omitempty"`
}
