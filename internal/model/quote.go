package model

import (
	"time"
)

type Quote struct {
	ID int `json:"id"`
	Text string `json:"text"`
	CreatedAt time.Time `json:"created_at"`
}