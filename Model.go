package main

import "time"

type User struct {
	id int
	firstName string
	lastName string
	birthDate time.Time
	gender Gender
	city string
	interests []string
	password string

}

type Gender int

const (
	Male Gender = iota
	Female
)
