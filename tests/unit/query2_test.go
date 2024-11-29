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
		"2020-01-01",
		"2020-01-02",
		"2020-01-03",
		"2020-01-04",
		"2020-01-05",
		"2020-01-06",
		"2020-01-07",
		"2020-01-08",
		"2020-01-09",
		"2020-01-10",
		"2020-01-11",
		"2020-01-12",
		"2020-01-13",
		"2020-01-14",
		"2020-01-15",
		"2020-01-16",
		"2020-01-17",
		"2020-01-18",
		"2020-01-19",
		"2020-01-20",
		"2020-01-21",
		"2020-01-22",
		"2020-01-23",
		"2020-01-24",
		"2020-01-25",
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
		return 35 * time.Second
	case n <= 4:
		return 30 * time.Second
	case n <= 6:
		return 25 * time.Second
	case n <= 8:
		return 23 * time.Second
	case n == 9:
		return 20 * time.Second
	case n == 10:
		return 19 * time.Second
	default:
		return 18 * time.Second
	}
}
