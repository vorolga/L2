package hash

import (
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"dev11/internal/model"
)

type Hash struct {
	sync.RWMutex
	hash map[int64][]*model.Event
}

func NewHash() (*Hash, error) {
	return &Hash{
		RWMutex: sync.RWMutex{},
		hash:    make(map[int64][]*model.Event, 0),
	}, nil
}

func (h *Hash) Check(id int64) (message string) {
	h.RLock()
	if _, ok := h.hash[id]; ok {
		h.RUnlock()
		return fmt.Sprintf("Id %d exists\n", id)
	} else {
		h.RUnlock()
		h.Lock()
		h.hash[id] = make([]*model.Event, 0)
		h.Unlock()
		return fmt.Sprintf("Your new ID is %d\n", id)
	}
}

func (h *Hash) Create(id int64, date time.Time) (model.Event, error) {
	h.RLock()
	if _, ok := h.hash[id]; !ok {
		h.RUnlock()
		h.Lock()
		h.hash[id] = make([]*model.Event, 0)
		h.Unlock()
	}
	nModel, err := model.NewEvent(id, date)
	if err != nil {
		return model.Event{}, err
	}
	h.Lock()
	h.hash[id] = append(h.hash[id], &nModel)
	h.Unlock()
	return nModel, nil
}

func (h *Hash) Update(id int64, date time.Time, newTime time.Time) error {
	indexesForUpdate := make([]int, 0)
	h.RLock()
	insecure := h.hash[id]
	h.RUnlock()
	for i, v := range insecure {
		if v.Date == date {
			indexesForUpdate = append(indexesForUpdate, i)
		}
	}
	if len(indexesForUpdate) == 0 {
		return fmt.Errorf("didnt'")
	}

	h.update(indexesForUpdate, id, newTime)

	return nil
}

func (h *Hash) Delete(id int64, date time.Time) error {
	indexesForDel := make([]int, 0)
	h.RLock()
	insecure := h.hash[id]
	h.RUnlock()
	for i, v := range insecure {
		if v.Date == date {
			indexesForDel = append(indexesForDel, i)
		}
	}
	if len(indexesForDel) == 0 {
		return fmt.Errorf("no events with %v date", date)
	}

	h.delete(indexesForDel, id)

	return nil
}

func (h *Hash) Day(id int64) ([]byte, error) {
	return checkTimeDay(h.hash[id])
}

func (h *Hash) Week(id int64) ([]byte, error) {
	return checkTimeWeek(h.hash[id])
}

func (h *Hash) Month(id int64) ([]byte, error) {
	return checkTimeMonth(h.hash[id])
}

var (
	dayNow, monthNow, yearNow = time.Now().Day(), time.Now().Month(), time.Now().Year()
	_, weekNow                = time.Now().ISOWeek()
)

func (h *Hash) delete(indx []int, id int64) {
	h.Lock()
	for _, i := range indx {
		h.hash[id][len(h.hash[id])-1] = h.hash[id][i]
		h.hash[id] = h.hash[id][:len(h.hash)-1]
	}
	h.Unlock()
}

func (h *Hash) update(indx []int, id int64, newTime time.Time) {
	h.Lock()
	for _, i := range indx {
		h.hash[id][i].Date = newTime
	}
	h.Unlock()
}

func checkTimeDay(userEvents []*model.Event) ([]byte, error) {
	result := make([]*model.Event, 0)
	for _, v := range userEvents {
		if v.Date.Day() == dayNow && v.Date.Month() == monthNow && v.Date.Year() == yearNow {
			result = append(result, v)
		}
	}
	return newJson(result)
}

func checkTimeWeek(userEvents []*model.Event) ([]byte, error) {
	result := make([]*model.Event, 0)
	for _, v := range userEvents {
		eventYear, eventWeek := v.Date.ISOWeek()
		if eventYear == yearNow && eventWeek == weekNow {
			result = append(result, v)
		}
	}
	return newJson(result)
}
func checkTimeMonth(userEvents []*model.Event) ([]byte, error) {
	result := make([]*model.Event, 0)
	for _, v := range userEvents {
		if v.Date.Month() == monthNow && v.Date.Year() == yearNow {
			result = append(result, v)
		}
	}
	return newJson(result)
}

type jsonEvents struct {
	Models []*model.Event `json:"events"`
}

func newJson(events []*model.Event) ([]byte, error) {
	var jsonevents jsonEvents

	for _, v := range events {
		jsonevents.Models = append(jsonevents.Models, v)
	}

	js, err := json.Marshal(jsonevents)
	if err != nil {
		return nil, err
	}

	return js, nil
}
