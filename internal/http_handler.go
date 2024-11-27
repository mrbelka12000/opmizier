package internal

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/mrbelka12000/optimizer/internal/models"
)

func (s *service) RegisterHandlers(mux *http.ServeMux) {
	mux.HandleFunc("/first", s.makeGetDataByCountryAndCustomersCountHandler())
	mux.HandleFunc("/second", s.makeGetDataByCountryAndCityAndCompanyAndCustomersCountHandler())
	mux.HandleFunc("/third", s.makeGetDataByCompaniesRankAndPastYearsHandler())
	mux.Handle("/metrics", promhttp.Handler())
}

// makeGetDataByCountryAndCustomersCountHandler to optimize a huge amount of small queries
func (s *service) makeGetDataByCountryAndCustomersCountHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		args := []any{
			r.URL.Query().Get("country"),
			r.URL.Query().Get("customers_count"),
		}

		if err := s.next.List(r.Context(), models.Query1, args); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			s.log.Error("cannot handle Query1 request", "error", err)
			return
		}

		w.Write([]byte("OK"))
	}
}

// makeGetDataByCountryAndCityAndCompanyAndCustomersCountHandler to optimize a database workload, performance
func (s *service) makeGetDataByCountryAndCityAndCompanyAndCustomersCountHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		args := []any{
			r.URL.Query().Get("country"),
			r.URL.Query().Get("city"),
			r.URL.Query().Get("company"),
			r.URL.Query().Get("customers_count"),
		}

		if err := s.next.List(r.Context(), models.Query2, args); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			s.log.Error("cannot handle Query2 request", "error", err)
			return
		}

		w.Write([]byte("OK"))
	}
}

// makeGetDataByCompaniesRankAndPastYearsHandler to optimize a database workload, performance but harder than previous
func (s *service) makeGetDataByCompaniesRankAndPastYearsHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		args := []any{
			r.URL.Query().Get("country"),
			r.URL.Query().Get("past_years"),
			r.URL.Query().Get("rank"),
		}

		if err := s.next.List(r.Context(), models.Query3, args); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			s.log.Error("cannot handle Query3 request", "error", err)
			return
		}

		w.Write([]byte("OK"))
	}
}
