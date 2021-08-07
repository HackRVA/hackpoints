package score

import (
	"hackpoints/datastore/in_memory"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetScore(t *testing.T) {
	server := &ScoreServer{
		Store: &in_memory.InMemoryScoreStore{
			BountyStore: &in_memory.InMemoryBountyStore{},
		},
	}

	tests := []struct {
		TestName            string
		expectedHTTPStastub int
		expectedResponse    string
	}{
		{
			TestName:            "should just respond",
			expectedHTTPStastub: http.StatusOK,
			expectedResponse:    `{"score":0}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.TestName, func(t *testing.T) {
			request := getScoreRequest()
			response := httptest.NewRecorder()

			server.Get(response, request)

			assertStatus(t, response.Code, tt.expectedHTTPStastub)
			assertResponseBody(t, response.Body.String(), tt.expectedResponse)
		})
	}
}

func getScoreRequest() *http.Request {
	req, _ := http.NewRequest(http.MethodPost, "/api/score/", nil)
	return req
}

func assertStatus(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("did not get correct status, got %d, want %d", got, want)
	}
}

func assertResponseBody(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("response body is wrong, got %q want %q", got, want)
	}
}
