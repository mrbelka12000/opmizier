package internal

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/mrbelka12000/optimizer/internal/models"
)

func (s *service) RegisterHandlers(mux *http.ServeMux) {
	mux.HandleFunc("/list", s.makeFilterHandler())
	mux.Handle("/metrics", promhttp.Handler())
}

func (s *service) makeFilterHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var args []any

		args = append(args, r.URL.Query().Get("country"))
		args = append(args, r.URL.Query().Get("customers_count"))

		if err := s.next.List(r.Context(), models.Query1, args); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			s.log.Error("cannot handle request", "error", err)
			return
		}

		w.Write([]byte("OK"))
	}
}
