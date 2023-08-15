package utils_test

import (
	"testing"

	"github.com/pericles-luz/go-base/pkg/utils"
	"github.com/stretchr/testify/require"
)

func TestDatesInIntervalMustHaveOneDateIfStartEqualsEnd(t *testing.T) {
	start := "2020-01-01"
	end := "2020-01-01"
	dates := utils.DatesInInterval(start, end)
	require.Equal(t, 1, len(dates))
}

func TestDatesInIntervalMustHaveTwoDatesIfStartIsBeforeEnd(t *testing.T) {
	start := "2020-01-01"
	end := "2020-01-02"
	dates := utils.DatesInInterval(start, end)
	require.Equal(t, 2, len(dates))
}

func TestDatesInIntervalMustReturnNilIfStartIsAfterEnd(t *testing.T) {
	start := "2020-01-02"
	end := "2020-01-01"
	dates := utils.DatesInInterval(start, end)
	require.Nil(t, dates)
}

func TestDatesInIntervalMustReturnNilIfStartIsInvalid(t *testing.T) {
	start := "2020-01-0"
	end := "2020-01-01"
	dates := utils.DatesInInterval(start, end)
	require.Nil(t, dates)
}

func TestDatesInIntervalMustReturnNilIfEndIsInvalid(t *testing.T) {
	start := "2020-01-01"
	end := "2020-01-0"
	dates := utils.DatesInInterval(start, end)
	require.Nil(t, dates)
}
