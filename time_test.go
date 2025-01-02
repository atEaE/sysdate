package systime

import (
	"testing"
	"time"
)

func TestCalcElapsedDay(t *testing.T) {
	// arrange
	testCases := []struct {
		name            string
		initUTC         time.Time
		nowUTC          time.Time
		expectedElapsed int
	}{
		{
			name:            "elapsed 0 day",
			initUTC:         time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
			nowUTC:          time.Date(2021, 1, 1, 23, 58, 7, 1, time.UTC),
			expectedElapsed: 0,
		},
		{
			name:            "elapsed 1 day",
			initUTC:         time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
			nowUTC:          time.Date(2021, 1, 2, 0, 10, 57, 1, time.UTC),
			expectedElapsed: 1,
		},
		{
			name:            "elapsed 2 day",
			initUTC:         time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
			nowUTC:          time.Date(2021, 1, 3, 0, 0, 0, 1, time.UTC),
			expectedElapsed: 2,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// act
			got := calcElapsedDay(tc.initUTC, tc.nowUTC)

			// assert
			if got != tc.expectedElapsed {
				t.Errorf("got: %v, expected: %v", got, tc.expectedElapsed)
			}
		})
	}
}
