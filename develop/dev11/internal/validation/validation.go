package validation

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
)

const (
	badRequest  = "Bad request"
	timeExample = "2006-01-02"
)

func ParseParams(r *http.Request) (int, time.Time, error) {
	id, ok := r.URL.Query()["user_id"]
	if ok {
		date, ok := r.URL.Query()["date"]
		if ok {
			return validateParams(id[0], date[0])
		} else {
			idRes, err := strconv.Atoi(id[0])
			if err != nil {
				return 0, time.Time{}, err
			}
			return idRes, time.Time{}, nil
		}
	}
	return 0, time.Time{}, fmt.Errorf(badRequest)
}

func validateParams(id string, date string) (int, time.Time, error) {

	idInt, err := strconv.Atoi(id)
	if err != nil {
		return 0, time.Time{}, err
	}

	err = ValidateID(idInt)
	if err != nil {
		return 0, time.Time{}, err
	}

	data, err := ValidateTime(date)
	if err != nil {
		return 0, time.Time{}, err
	}

	return idInt, data, nil
}

func ValidateID(id int) error {
	if id <= 0 {
		return fmt.Errorf("id should be 0 < id < 2 ^ 63")
	}
	return nil
}

func ValidateTime(t string) (time.Time, error) {
	data, err := time.Parse(timeExample, t)
	if err != nil {
		return time.Time{}, err
	}

	if data.Before(time.Now()) {
		return time.Time{}, fmt.Errorf("time before now")
	}

	return data, nil
}
