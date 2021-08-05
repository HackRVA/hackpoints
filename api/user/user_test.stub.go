package user

import (
	"bytes"
	"encoding/json"
	"errors"
	"hackpoints/models"
	"net/http"
	"strings"
	"testing"

	"github.com/shaj13/go-guardian/v2/auth"
)

type stubUserStore struct{}

func (stub *stubUserStore) GetMemberByEmail(email string) (models.Member, error) {
	if val, ok := testUsers[strings.ToLower(email)]; !ok {
		return val, errors.New("not a valid member email")
	}

	return testUsers[email], nil
}

func (stub *stubUserStore) RegisterUser(creds models.Credentials) error {
	if _, ok := testUsers[creds.Email]; ok {
		return errors.New("error registering user")
	}
	return nil
}

var testUsers = map[string]models.Member{
	"test": {
		Email: "test",
	},
}

type stubAuthProvider struct{}

func (stub *stubAuthProvider) User(r *http.Request) auth.Info {
	return auth.UserFromCtx(r.Context())
}

func (stub *stubAuthProvider) IssueAccessToken(info auth.Info) (string, error) {
	return "", nil
}

type fakeHandler struct{}

func (f fakeHandler) ServeHTTP(http.ResponseWriter, *http.Request) {}

func (stub *stubAuthProvider) AuthMiddleware(next http.Handler) http.Handler {
	return fakeHandler{}
}

func newGetUserRequest() *http.Request {
	req, _ := http.NewRequest(http.MethodGet, "/api/user", nil)
	return req
}

func newRegisterUserRequest(creds models.Credentials) *http.Request {
	reqBody, _ := json.Marshal(creds)
	req, _ := http.NewRequest(http.MethodGet, "/api/user", bytes.NewReader(reqBody))
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
