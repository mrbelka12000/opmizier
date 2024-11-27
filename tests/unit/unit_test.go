package unit

import (
	"context"
	"fmt"
	"math/rand"
	"net/http"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestListWithDeadline(t *testing.T) {

	var (
		// Please do not change
		deadline      = 10 * time.Second
		requestsCount = 10000
		workersPool   = make(chan struct{}, 50)
		start         = time.Now()
		wg            sync.WaitGroup
	)

	ctx, cancel := context.WithTimeout(context.Background(), deadline)
	defer cancel()

	go func() {
		<-ctx.Done()
		t.Fail()
	}()

	for i := 0; i < requestsCount; i++ {
		workersPool <- struct{}{}
		wg.Add(1)
		go func() {
			assert.Nil(t, sendRequest())
			<-workersPool
			wg.Done()
		}()
	}

	if t.Failed() {
		fmt.Println("time exceeded")
	}

	wg.Wait()
	close(workersPool)
	fmt.Printf("elapsed: %v\n", time.Since(start).Seconds())
}

func sendRequest() error {
	var countries = []string{
		"Afganistan",
		"France",
		"Italy",
		"Japan",
	}

	url := fmt.Sprintf("http://localhost:8080/list?country=%s&customers_count=%v", countries[rand.Intn(len(countries))], rand.Intn(30))
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	if resp.StatusCode != http.StatusOK {
		return err
	}

	return nil
}
