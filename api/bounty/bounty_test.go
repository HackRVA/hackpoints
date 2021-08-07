package bounty

import (
	"bytes"
	"encoding/json"
	"hackpoints/datastore/in_memory"
	"hackpoints/models"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/shaj13/go-guardian/v2/auth"
)

func TestNewBounty(t *testing.T) {
	bs := &BountyServer{
		&in_memory.Store{},
	}

	successResponse, _ := json.Marshal(models.EndpointSuccess{
		Ack: true,
	})

	tests := []struct {
		TestName            string
		ID                  string
		Title               string
		Description         string
		Endorsements        []models.Member
		IsOpen              bool
		expectedHTTPStastub int
		expectedResponse    string
	}{
		{
			TestName:            "should create a new bounty",
			ID:                  "someID",
			Title:               "a new bounty",
			Description:         "this is a fake bounty",
			Endorsements:        []models.Member{{Email: "test"}},
			IsOpen:              true,
			expectedHTTPStastub: http.StatusOK,
			expectedResponse:    string(successResponse),
		},
		{
			TestName:            "should fail if we try to create a bounty without title",
			ID:                  "someID1",
			Title:               "",
			Description:         "this is a fake bounty",
			Endorsements:        []models.Member{{Email: "test"}},
			IsOpen:              true,
			expectedHTTPStastub: http.StatusBadRequest,
			expectedResponse:    ErrNoTitle.Error() + "\n",
		},
		{
			TestName:            "should fail if we try to create a bounty with a short description",
			ID:                  "someID2",
			Title:               "Some TITLE",
			Description:         "this is a",
			Endorsements:        []models.Member{{Email: "test"}},
			IsOpen:              true,
			expectedHTTPStastub: http.StatusBadRequest,
			expectedResponse:    ErrBadDescription.Error() + "\n",
		},
		{
			TestName:            "should fail if we try to create a bounty with no description",
			ID:                  "someID2",
			Title:               "Some TITLE",
			Description:         "",
			Endorsements:        []models.Member{{Email: "test"}},
			IsOpen:              true,
			expectedHTTPStastub: http.StatusBadRequest,
			expectedResponse:    ErrBadDescription.Error() + "\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.TestName, func(t *testing.T) {
			request := newBountyRequest(models.Bounty{
				Title:        tt.Title,
				Description:  tt.Description,
				Endorsements: tt.Endorsements,
				IsOpen:       tt.IsOpen,
			})
			response := httptest.NewRecorder()

			bs.New(response, request)

			assertStatus(t, response.Code, tt.expectedHTTPStastub)
			assertResponseBody(t, response.Body.String(), tt.expectedResponse)
		})
	}
}

func newBountyRequest(m models.Bounty) *http.Request {
	reqBody, _ := json.Marshal(m)
	req, _ := http.NewRequest(http.MethodPost, "/api/bounty/new", bytes.NewReader(reqBody))
	return req
}

func TestUpdateBounty(t *testing.T) {
	bs := &BountyServer{
		&in_memory.Store{},
	}

	successResponse, _ := json.Marshal(models.EndpointSuccess{
		Ack: true,
	})

	bs.Store.NewBounty(models.Bounty{
		Title:       "a new bounty",
		Description: "this is a fake bounty",
	})

	tests := []struct {
		TestName            string
		ID                  string
		Title               string
		Description         string
		Endorsements        []models.Member
		IsOpen              bool
		expectedHTTPStastub int
		expectedResponse    string
	}{
		{
			TestName:            "should update an existing bounty",
			ID:                  "1",
			Title:               "a new bounty",
			Description:         "this is a fake bounty",
			Endorsements:        []models.Member{{Email: "test"}},
			IsOpen:              true,
			expectedHTTPStastub: http.StatusOK,
			expectedResponse:    string(successResponse),
		},
		{
			TestName:            "should fail to update if bounty doesn't exist",
			ID:                  "shouldn't exist",
			Title:               "a new bounty",
			Description:         "this is a fake bounty",
			Endorsements:        []models.Member{{Email: "test"}},
			IsOpen:              true,
			expectedHTTPStastub: http.StatusBadRequest,
			expectedResponse:    ErrUpdatingBounty.Error() + "\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.TestName, func(t *testing.T) {
			request := updateBountyRequest(models.Bounty{
				ID:           tt.ID,
				Title:        tt.Title,
				Description:  tt.Description,
				Endorsements: tt.Endorsements,
				IsOpen:       tt.IsOpen,
			})
			response := httptest.NewRecorder()

			bs.Update(response, request)

			assertStatus(t, response.Code, tt.expectedHTTPStastub)
			assertResponseBody(t, response.Body.String(), tt.expectedResponse)
		})
	}
}

func updateBountyRequest(m models.Bounty) *http.Request {
	reqBody, _ := json.Marshal(m)
	req, _ := http.NewRequest(http.MethodPost, "/api/bounty/update", bytes.NewReader(reqBody))
	return req
}

func TestGetBounty(t *testing.T) {
	bs := &BountyServer{
		&in_memory.Store{},
	}

	bs.Store.NewBounty(models.Bounty{
		Title:       "a new bounty",
		Description: "this is a fake bounty",
		IsOpen:      true,
	})

	tests := []struct {
		TestName            string
		ID                  string
		Title               string
		Description         string
		Endorsements        []models.Member
		IsOpen              bool
		expectedHTTPStastub int
		expectedResponse    string
	}{
		{
			TestName:            "should update an existing bounty",
			ID:                  "1",
			expectedHTTPStastub: http.StatusOK,
			expectedResponse:    "[{\"id\":\"1\",\"title\":\"a new bounty\",\"description\":\"this is a fake bounty\",\"endorsements\":[{\"id\":\"\",\"name\":\"\",\"email\":\"test\"}],\"isOpen\":true}]",
		},
		{
			TestName:            "should throw error if we try to get a bounty that doesn't exist",
			ID:                  "doesn't exist",
			Title:               "a new bounty",
			Description:         "this is a fake bounty",
			IsOpen:              true,
			expectedHTTPStastub: http.StatusNotFound,
			expectedResponse:    ErrBountyNotFound.Error() + "\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.TestName, func(t *testing.T) {
			request := getBountyRequest(models.Bounty{
				ID: tt.ID,
			})
			response := httptest.NewRecorder()

			bs.Get(response, request)

			assertStatus(t, response.Code, tt.expectedHTTPStastub)
			assertResponseBody(t, response.Body.String(), tt.expectedResponse)
		})
	}
}

func getBountyRequest(m models.Bounty) *http.Request {
	reqBody, _ := json.Marshal(m)
	req, _ := http.NewRequest(http.MethodPost, "/api/bounty/", bytes.NewReader(reqBody))
	return req
}

func TestGetAllBounties(t *testing.T) {
	bs := &BountyServer{
		&in_memory.Store{},
	}

	bs.Store.NewBounty(models.Bounty{
		Title:       "a new bounty",
		Description: "this is a fake bounty",
	})

	tests := []struct {
		TestName            string
		ID                  string
		Title               string
		Description         string
		Endorsements        []models.Member
		IsOpen              bool
		expectedHTTPStastub int
		expectedResponse    string
	}{
		{
			TestName:            "should update an existing bounty",
			Title:               "a new bounty",
			Description:         "this is a fake bounty",
			IsOpen:              true,
			expectedHTTPStastub: http.StatusOK,
			expectedResponse:    "[{\"id\":\"1\",\"title\":\"a new bounty\",\"description\":\"this is a fake bounty\",\"endorsements\":[{\"id\":\"\",\"name\":\"\",\"email\":\"test\"}],\"isOpen\":true},{\"id\":\"2\",\"title\":\"a new bounty\",\"description\":\"this is a fake bounty\",\"endorsements\":null,\"isOpen\":true},{\"id\":\"3\",\"title\":\"a new bounty\",\"description\":\"this is a fake bounty\",\"endorsements\":null,\"isOpen\":true},{\"id\":\"4\",\"title\":\"a new bounty\",\"description\":\"this is a fake bounty\",\"endorsements\":null,\"isOpen\":true}]",
		},
		{
			TestName:            "should throw error if we try to get a bounty that doesn't exist",
			ID:                  "doesn't exist",
			Title:               "a new bounty",
			Description:         "this is a fake bounty",
			Endorsements:        []models.Member{{Email: "test"}},
			IsOpen:              true,
			expectedHTTPStastub: http.StatusNotFound,
			expectedResponse:    ErrBountyNotFound.Error() + "\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.TestName, func(t *testing.T) {
			request := getAllBountyRequest(models.Bounty{
				ID: tt.ID,
			})
			response := httptest.NewRecorder()

			bs.Get(response, request)

			assertStatus(t, response.Code, tt.expectedHTTPStastub)
			assertResponseBody(t, response.Body.String(), tt.expectedResponse)
		})
	}
}

func getAllBountyRequest(m models.Bounty) *http.Request {
	reqBody, _ := json.Marshal(m)
	req, _ := http.NewRequest(http.MethodPost, "/api/bounty/", bytes.NewReader(reqBody))
	return req
}

// currently not working until we can get the user properly
func testEndorseBounty(t *testing.T) {
	bs := &BountyServer{
		&in_memory.Store{},
	}

	bs.Store.NewBounty(models.Bounty{
		ID:           "someID1",
		Title:        "a new bounty",
		Description:  "this is a fake bounty",
		Endorsements: []models.Member{{Email: "test"}},
		IsOpen:       true,
	})

	tests := []struct {
		TestName            string
		ID                  string
		Title               string
		Description         string
		Endorsements        []models.Member
		User                models.Member
		Bounty              models.Bounty
		IsOpen              bool
		expectedHTTPStastub int
		expectedResponse    string
	}{
		{
			TestName:     "should endorse an existing bounty",
			ID:           "someID1",
			Title:        "a new bounty",
			Description:  "this is a fake bounty",
			Endorsements: []models.Member{},
			User:         models.Member{Email: "test"},
			Bounty: models.Bounty{
				ID: "someID1",
			},
			IsOpen:              true,
			expectedHTTPStastub: http.StatusOK,
			expectedResponse:    "{\"ID\":\"someID1\",\"Title\":\"a new bounty\",\"Description\":\"this is a fake bounty\",\"Endorsements\":[{\"id\":\"\",\"name\":\"\",\"email\":\"test\"}],\"IsOpen\":true}",
		},
	}

	for _, tt := range tests {
		t.Run(tt.TestName, func(t *testing.T) {
			request := postEndorsementRequest(models.Bounty{
				ID: tt.ID,
			})
			response := httptest.NewRecorder()

			bs.Endorse(response, request)

			assertStatus(t, response.Code, tt.expectedHTTPStastub)
			assertResponseBody(t, response.Body.String(), tt.expectedResponse)
		})
	}
}

func postEndorsementRequest(m models.Bounty) *http.Request {
	reqBody, _ := json.Marshal(m)
	req, _ := http.NewRequest(http.MethodPost, "/api/bounty/endorse", bytes.NewReader(reqBody))
	auth.NewDefaultUser("test", "test", []string{}, nil)
	return auth.RequestWithUser(auth.User(req), req)
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
