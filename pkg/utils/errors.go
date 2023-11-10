package utils

import (
	"database/sql"
	"errors"
	"net/http"
	"syscall"
	"time"

	"github.com/getsentry/sentry-go"
	sentryhttp "github.com/getsentry/sentry-go/http"
)

var sentryHandler = sentryhttp.New(sentryhttp.Options{})

// Verify if there is an error and, in this case, try to send it to Sentry
// and return it. If is a running test, the error isn't sent to Sentry.
func ManageError(err error) error {
	if err == sql.ErrNoRows {
		return err
	}
	if errors.Is(err, syscall.ECONNRESET) {
		return err
	}
	if err != nil && sentry.CurrentHub().Client() != nil {
		sentry.CaptureException(err)
		sentry.Flush(2 * time.Second)
	}
	return err
}

// Encapsulate the handler with Sentry
func EnhanceSentryEvent(handler http.HandlerFunc) http.HandlerFunc {
	return sentryHandler.HandleFunc(func(rw http.ResponseWriter, r *http.Request) {
		if hub := sentry.GetHubFromContext(r.Context()); hub != nil {
			hub.Scope().SetTag("someRandomTag", "maybeYouNeedIt")
		}
		handler(rw, r)
	})
}
