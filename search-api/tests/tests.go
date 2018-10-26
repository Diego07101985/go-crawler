package api-tests

import (
	"errors"
	"testing"
)

var c = ReposClient{}

func TestGet(t *testing.T) {
	assert := assert.New(t)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	tests := []struct {
		description   string
		responder     httpmock.Responder
		expectedRepos []Repo
		expectedError error
	}{
		{
			description:   "github api success",
			responder:     httpmock.NewStringResponder(200, `[{"name": "test", "description": "a test"}]`),
			expectedRepos: []Repo{Repo{Name: "test", Description: "a test"}},
			expectedError: nil,
		}, {
			description:   "github api success, no repos",
			responder:     httpmock.NewStringResponder(200, `[]`),
			expectedRepos: []Repo{},
			expectedError: nil,
		}, {
			description:   "github api failure, not found",
			responder:     httpmock.NewStringResponder(404, `{"message": "not found"}`),
			expectedRepos: []Repo(nil),
			expectedError: errors.New("github api: no results found"),
		},
		// not all cases are tested, but this is enough of a sample
	}

	for _, tc := range tests {
		httpmock.RegisterResponder("GET", "https://api.github.com/users/fake/repos", tc.responder)

		r, err := c.Get("fake")

		assert.Equal(r, tc.expectedRepos, tc.description)
		assert.Equal(err, tc.expectedError, tc.description)

		httpmock.Reset()
	}
}