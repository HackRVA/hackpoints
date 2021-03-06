package user

import (
	"hackpoints/models"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/shaj13/go-guardian/v2/auth"
)

func TestGetUser(t *testing.T) {
	server := &UserServer{&stubUserStore{}, &stubAuthProvider{}}

	tests := []struct {
		TestName            string
		userName            string
		resources           []string
		expectedHTTPStastub int
		expectedResponse    string
	}{
		{
			TestName:            "should return currently logged in user",
			userName:            "test",
			expectedHTTPStastub: http.StatusOK,
			expectedResponse:    "{\"id\":\"\",\"name\":\"\",\"email\":\"test\"}",
		},
		{
			TestName:            "should return unauthorized if email doesn't exist",
			userName:            "doesn't exist",
			expectedHTTPStastub: http.StatusUnauthorized,
			expectedResponse:    "user not found\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.TestName, func(t *testing.T) {
			request := newGetUserRequest()
			response := httptest.NewRecorder()

			authInfo := auth.NewDefaultUser(tt.userName, tt.userName, tt.resources, nil)
			server.GetUser(response, auth.RequestWithUser(authInfo, request))

			assertStatus(t, response.Code, tt.expectedHTTPStastub)
			assertResponseBody(t, response.Body.String(), tt.expectedResponse)
		})
	}
}

func TestRegisterUser(t *testing.T) {
	server := &UserServer{&stubUserStore{}, &stubAuthProvider{}}

	tests := []struct {
		TestName            string
		userName            string
		resources           []string
		creds               models.Credentials
		expectedHTTPStastub int
		expectedResponse    string
	}{
		{
			TestName: "should register a user",
			userName: "doesn't exist",
			creds: models.Credentials{
				Email:    "doesn't exist",
				Password: "password",
			},
			expectedHTTPStastub: http.StatusOK,
			expectedResponse:    "{\"ack\":true}",
		},
		{
			TestName: "should fail to register a user if they already exist",
			userName: "test",
			creds: models.Credentials{
				Email:    "test",
				Password: "password",
			},
			expectedHTTPStastub: http.StatusBadRequest,
			expectedResponse:    "error registering user\n",
		},
		{
			TestName: "should fail if password isn't provided",
			userName: "test",
			creds: models.Credentials{
				Email:    "doesn't exist 1",
				Password: "",
			},
			expectedHTTPStastub: http.StatusBadRequest,
			expectedResponse:    "password must be longer\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.TestName, func(t *testing.T) {
			request := newRegisterUserRequest(tt.creds)
			response := httptest.NewRecorder()

			// authInfo := auth.NewDefaultUser(tt.userName, tt.userName, tt.resources, nil)
			server.Register(response, request)

			assertStatus(t, response.Code, tt.expectedHTTPStastub)
			assertResponseBody(t, response.Body.String(), tt.expectedResponse)
		})
	}
}
