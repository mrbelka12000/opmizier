package unit

import (
	"context"
	"fmt"
	"math/rand"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// 128s
// 86s
// 45s create index idx_test on customers (city, "Subscription Date");
func TestQuery2(t *testing.T) {
	var (
		deadline     = 10 * time.Second
		requestCount = 100
		start        = time.Now()
	)

	ctx, cancel := context.WithTimeout(context.Background(), deadline)
	defer cancel()

	go func() {
		<-ctx.Done()
		t.Fail()
	}()

	for i := 0; i < requestCount; i++ {
		assert.Nil(t, sendRequestForQuery2())
	}

	if t.Failed() {
		fmt.Println("time exceeded")
	}

	fmt.Printf("elapsed: %v\n", time.Since(start).Seconds())
}

func sendRequestForQuery2() error {
	var dates = []string{
		"2020-01-01",
		"2020-01-15",
		"2020-03-21",
		"2020-02-27",
		"2020-11-04",
		"2021-01-09",
		"2021-03-04",
		"2021-04-01",
		"2021-05-11",
		"2021-03-07",
		"2021-01-23",
		"2020-12-06",
		"2020-11-23",
	}

	url := fmt.Sprintf("http://localhost:8080/second?subscription_date=%v&customers_count=%v", dates[rand.Intn(len(dates))], rand.Intn(65))
	fmt.Println(url)

	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	if resp.StatusCode != http.StatusOK {
		return assert.AnError
	}

	return nil
}
