package utils

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/getsentry/sentry-go"
	sentryhttp "github.com/getsentry/sentry-go/http"
)

var sentryHandler = sentryhttp.New(sentryhttp.Options{})

func ManageError(err error) error {
	if err == sql.ErrNoRows {
		return err
	}
	if err != nil && sentry.CurrentHub().Client() != nil {
		sentry.CaptureException(err)
		sentry.Flush(2 * time.Second)
	}
	return err
}

func EnhanceSentryEvent(handler http.HandlerFunc) http.HandlerFunc {
	return sentryHandler.HandleFunc(func(rw http.ResponseWriter, r *http.Request) {
		if hub := sentry.GetHubFromContext(r.Context()); hub != nil {
			hub.Scope().SetTag("someRandomTag", "maybeYouNeedIt")
		}
		handler(rw, r)
	})
}
