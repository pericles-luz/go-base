package utils

import (
	"time"

	"github.com/getsentry/sentry-go"
)

func ManageError(err error) error {
	if err != nil && sentry.CurrentHub().Client() != nil {
		sentry.CaptureException(err)
		sentry.Flush(2 * time.Second)
	}
	return err
}
