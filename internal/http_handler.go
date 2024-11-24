package internal

import (
	"net/http"

	"github.com/mrbelka12000/optimizer/internal/models"
)

func (s *Service) RegisterHandlers(mux *http.ServeMux) {
	mux.HandleFunc("/list", s.makeListHandler())
}

func (s *Service) makeListHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req models.Request

		if err := s.decoder.Decode(&req, r.URL.Query()); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			s.log.Error("cannot decode url params", err)
			return
		}

		if err := s.next.List(r.Context(), req); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			s.log.Error("cannot handle request", err)
			return
		}

		w.Write([]byte("OK"))
	}
}
