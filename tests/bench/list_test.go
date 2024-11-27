package bench

import (
	"fmt"
	"math/rand"
	"net/http"
	"testing"
)

func BenchmarkList(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		var countries = []string{
			"Afganistan",
			"France",
			"Italy",
			"Japan",
		}
		for pb.Next() {
			url := fmt.Sprintf("http://localhost:8080/list?county=%s&customers_count=%v", countries[rand.Intn(len(countries))], rand.Intn(30))

			resp, err := http.Get(url)
			if err != nil {
				panic(err)
			}

			if resp.StatusCode != http.StatusOK {
				b.Fail()
			}
		}
	})
}
