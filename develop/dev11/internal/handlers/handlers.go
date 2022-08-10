package server

import (
	"encoding/json"
	"net/http"

	event "dev11/internal/model"
	"dev11/internal/usecase"
	"dev11/internal/validation"

	"go.uber.org/zap"
)

type Handler struct {
	repo   usecase.Repository
	logger *zap.SugaredLogger
}

func New(repo usecase.Repository, logger *zap.SugaredLogger) Handler {
	return Handler{repo: repo, logger: logger}
}

func Register(repo usecase.Repository, logger *zap.SugaredLogger) *http.ServeMux {
	h := New(repo, logger)

	mux := http.NewServeMux()
	mux.HandleFunc("/create_event", h.Create)
	mux.HandleFunc("/delete_event", h.Delete)
	mux.HandleFunc("/update_event", h.Update)
	mux.HandleFunc("/events_for_day", h.Today)
	mux.HandleFunc("/events_for_week", h.Week)
	mux.HandleFunc("/events_for_month", h.Month)

	return mux
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	id, date, err := validation.ParseParams(r)
	if err != nil {
		h.logger.Error(
			zap.String("ERROR", err.Error()),
			zap.Int("ANSWER STATUS", http.StatusBadRequest),
		)

		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	m, err := event.NewEvent(int64(id), date)
	if err != nil {
		h.logger.Error(
			zap.String("ERROR", err.Error()),
			zap.Int("ANSWER STATUS", http.StatusBadRequest),
		)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	model, err := h.repo.Create(m.ID, m.Date)
	if err != nil {
		h.logger.Error(
			zap.String("ERROR", err.Error()),
			zap.Int("ANSWER STATUS", http.StatusInternalServerError),
		)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	res, err := json.Marshal(model)
	if err != nil {
		h.logger.Error(
			zap.String("ERROR", err.Error()),
			zap.Int("ANSWER STATUS", http.StatusInternalServerError),
		)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	status, err := w.Write(res)
	if err != nil {
		h.logger.Error(
			zap.String("ERROR", err.Error()),
			zap.Int("ANSWER STATUS", status),
		)
		http.Error(w, err.Error(), status)
		return
	}

}

func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	id, t, err := validation.ParseParams(r)
	if err != nil {
		h.logger.Error(
			zap.String("ERROR", err.Error()),
			zap.Int("ANSWER STATUS", http.StatusBadRequest),
		)
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	if r.URL.Query()["new_date"][0] == "" {
		h.logger.Error(
			zap.String("ERROR", err.Error()),
			zap.Int("ANSWER STATUS", http.StatusBadRequest),
		)
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	newTime, err := validation.ValidateTime(r.URL.Query()["new_date"][0])
	if err != nil {
		h.logger.Error(
			zap.String("ERROR", err.Error()),
			zap.Int("ANSWER STATUS", http.StatusBadRequest),
		)
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	err = h.repo.Update(int64(id), t, newTime)
	if err != nil {
		h.logger.Error(
			zap.String("ERROR", err.Error()),
			zap.Int("ANSWER STATUS", http.StatusInternalServerError),
		)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.WriteHeader(200)
	_, err = w.Write([]byte("All events with updated"))
	if err != nil {
		h.logger.Error(
			zap.String("ERROR", err.Error()),
			zap.Int("ANSWER STATUS", http.StatusInternalServerError),
		)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	id, time, err := validation.ParseParams(r)
	if err != nil {
		h.logger.Error(
			zap.String("ERROR", err.Error()),
			zap.Int("ANSWER STATUS", http.StatusBadRequest),
		)
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	err = h.repo.Delete(int64(id), time)
	if err != nil {
		h.logger.Error(
			zap.String("ERROR", err.Error()),
			zap.Int("ANSWER STATUS", http.StatusInternalServerError),
		)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	_, err = w.Write([]byte("Event deleted"))
	if err != nil {
		h.logger.Error(
			zap.String("ERROR", err.Error()),
			zap.Int("ANSWER STATUS", http.StatusInternalServerError),
		)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h *Handler) Today(w http.ResponseWriter, r *http.Request) {

	id, _, err := validation.ParseParams(r)
	if err != nil {
		h.logger.Error(
			zap.String("ERROR", err.Error()),
			zap.Int("ANSWER STATUS", http.StatusBadRequest),
		)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	res, err := h.repo.Day(int64(id))
	if err != nil {
		h.logger.Error(
			zap.String("ERROR", err.Error()),
			zap.Int("ANSWER STATUS", http.StatusBadRequest),
		)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = w.Write(res)
	if err != nil {
		h.logger.Error(
			zap.String("ERROR", err.Error()),
			zap.Int("ANSWER STATUS", http.StatusInternalServerError),
		)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h *Handler) Week(w http.ResponseWriter, r *http.Request) {

	id, _, err := validation.ParseParams(r)
	if err != nil {
		h.logger.Error(
			zap.String("ERROR", err.Error()),
			zap.Int("ANSWER STATUS", http.StatusBadRequest),
		)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	res, err := h.repo.Week(int64(id))
	if err != nil {
		h.logger.Error(
			zap.String("ERROR", err.Error()),
			zap.Int("ANSWER STATUS", http.StatusBadRequest),
		)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = w.Write(res)
	if err != nil {
		h.logger.Error(
			zap.String("ERROR", err.Error()),
			zap.Int("ANSWER STATUS", http.StatusInternalServerError),
		)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h *Handler) Month(w http.ResponseWriter, r *http.Request) {
	id, _, err := validation.ParseParams(r)
	if err != nil {
		h.logger.Error(
			zap.String("ERROR", err.Error()),
			zap.Int("ANSWER STATUS", http.StatusBadRequest),
		)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	res, err := h.repo.Month(int64(id))
	if err != nil {
		h.logger.Error(
			zap.String("ERROR", err.Error()),
			zap.Int("ANSWER STATUS", http.StatusBadRequest),
		)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = w.Write(res)
	if err != nil {
		h.logger.Error(
			zap.String("ERROR", err.Error()),
			zap.Int("ANSWER STATUS", http.StatusInternalServerError),
		)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
