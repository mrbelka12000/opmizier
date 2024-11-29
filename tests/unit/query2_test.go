package unit

import (
	"context"
	"fmt"
	"math/rand"
	"net/http"
	"runtime"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func TestQuery2(t *testing.T) {
	var (
		deadline     = getDeadlineInSecondsForQuery2(runtime.NumCPU())
		requestCount = 100
		ticker       = time.NewTicker(95 * time.Millisecond)
		start        = time.Now()
		wg           sync.WaitGroup
		isFailed     uint32
	)

	ctx, _ := context.WithTimeout(context.Background(), deadline)

	go func() {
		<-ctx.Done()
		t.Fail()
		fmt.Printf("exceeded deadline: %v\n", deadline)
	}()

	for i := 0; i < requestCount; i++ {
		<-ticker.C
		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			if err := sendRequestForQuery2(); err != nil {
				atomic.AddUint32(&isFailed, 1)
			}
		}(&wg)
	}

	wg.Wait()
	ticker.Stop()

	if isFailed != 0 {
		t.Fail()
		fmt.Printf("server can not handle load, try to investigate problem. Lost %d connections.\n", isFailed)
	}

	fmt.Printf("elapsed: %v\n", time.Since(start).Seconds())
}

func sendRequestForQuery2() error {
	dates := []string{
		"2021-08-13",
		"2020-01-27",
		"2021-09-28",
		"2020-12-31",
		"2020-12-04",
		"2020-04-29",
		"2020-03-09",
		"2021-05-04",
		"2022-02-16",
		"2021-06-12",
		"2021-03-12",
		"2020-03-08",
		"2020-11-20",
		"2021-09-06",
		"2022-02-04",
		"2020-01-04",
		"2020-07-02",
		"2020-08-16",
		"2021-06-16",
		"2020-09-17",
		"2021-02-03",
		"2020-12-15",
		"2020-08-08",
		"2020-06-08",
		"2021-09-29",
	}

	url := fmt.Sprintf("http://localhost:8080/second?subscription_date=%s&customers_count=%v", dates[rand.Intn(len(dates))], rand.Intn(5))

	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("could not send request: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("expected status code %v, got %v", http.StatusOK, resp.StatusCode)
	}

	return nil
}

func getDeadlineInSecondsForQuery2(n int) time.Duration {
	switch {
	case n == 1:
		return 25 * time.Second
	case n <= 4:
		return 20 * time.Second
	case n <= 6:
		return 15 * time.Second
	case n <= 8:
		return 13 * time.Second
	case n == 9:
		return 12 * time.Second
	case n == 10:
		return 10 * time.Second
	default:
		return 8 * time.Second
	}
}
