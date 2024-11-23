package internal

import (
	"fmt"
	"net/http"

	"github.com/mrbelka12000/optimizer/internal/models"
)

func (s *Service) RegisterHandlers(mux *http.ServeMux) {
	mux.HandleFunc("/list", s.makeListHandler())
}

func (s *Service) makeListHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := s.next.List(r.Context(), models.Data{
			ID:          1555,
			FirstName:   "Jody",
			IsOrEnabled: true,
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			fmt.Println(err)
			return
		}

		w.Write([]byte("OK"))
	}
}
