package model

import (
	"errors"
	"time"
)

type Event struct {
	ID   int64     `json:"user_id"`
	Date time.Time `json:"date"`
}

type Error struct {
	Err string `json:"error"`
}

func NewEvent(id int64, time time.Time) (Event, error) {
	e := Event{
		ID:   id,
		Date: time,
	}

	return e, e.validate()
}

func (e *Event) validate() error {
	if e.Date.Before(time.Now()) || e.ID <= 0 {
		return errors.New("Bad request")
	}
	return nil
}
