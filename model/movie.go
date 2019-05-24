package model

import "time"

type MoveDetail struct {
	Title string

	Publish time.Time

	Type []string

	Actors []string

	Director []string

	Description string
}
