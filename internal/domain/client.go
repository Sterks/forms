package domain

import "time"

type Client struct {
	Firstname  string
	Lastname   string
	Patronomic string
	Position   string
	Company    string
	Phone      string
	Email      string
	CreateAt   time.Time
}
