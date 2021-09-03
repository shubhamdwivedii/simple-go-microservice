package homepage

import (
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func Test(t *testing.T) {
	// And Example of Table Test
	tests := []struct {
		name           string
		in             *http.Request
		out            *httptest.ResponseRecorder
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "good",
			in:             httptest.NewRequest("GET", "/", nil),
			out:            httptest.NewRecorder(),
			expectedStatus: http.StatusOK,
			expectedBody:   message, // message is declared already in same package.
		}, // Try changing message to force fail.
	}

	for _, test := range tests {
		test := test // creating shadow to avoid test fail when run in parallel
		t.Run(test.name, func(t *testing.T) {
			h := NewHandlers(log.New(os.Stdout, "TESTING: ", log.LstdFlags|log.Lshortfile))
			h.Home(test.out, test.in)
			if test.out.Code != test.expectedStatus {
				t.Logf("expected: %d\ngot: %d\n", test.expectedStatus, test.out.Code)
				t.Fail()
			}

			body := test.out.Body.String()
			if body != test.expectedBody {
				t.Logf("expected: %s\ngot: %s\n", test.expectedBody, body)
				t.Fail()
			}
		})
	}
}
