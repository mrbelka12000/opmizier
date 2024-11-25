package internal

import (
	"net/http"

	"github.com/mrbelka12000/optimizer/internal/models"
)

func (s *Service) RegisterHandlers(mux *http.ServeMux) {
	mux.HandleFunc("/list", s.makeFilterHandler())
}

func (s *Service) makeFilterHandler() http.HandlerFunc {
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
