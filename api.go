package drive

import (
	"context"
	"log"
	"math"
	"net/http"
	"time"

	"golang.org/x/net/context/ctxhttp"
)

type apiRequester struct {
	maxRetry     int
	retryBackoff int
	hc           *http.Client
}

func (ar *apiRequester) Do(ctx context.Context, req *http.Request) (res *http.Response, err error) {

	for i := 0; i < ar.maxRetry+1; i++ {
		if i > 0 {
			log.Printf("[Playcourt Drive] Retry on %d time(s) | %s %s", i, req.Method, req.URL.Path)
			backoffN := math.Pow(float64(ar.retryBackoff), float64(i))
			sleepDuration := time.Duration(backoffN) * time.Millisecond
			time.Sleep(sleepDuration)
		}

		res, err = ctxhttp.Do(ctx, ar.hc, req)
		if err != nil {
			if err == context.DeadlineExceeded {
				return
			}
			continue
		}

		break
	}

	return
}
