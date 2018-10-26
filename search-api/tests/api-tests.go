package api-tests

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAuth(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		description  string
		url          string
		expectedBody string
		expectedCode int
	}{
		{
			description:  "missing arg user",
			url:          "/ping",
			expectedBody: "MISSING_ARG_USER\n",
			expectedCode: 400,
		}, {
			description:  "bad user",
			url:          "/ping?user=baduser",
			expectedBody: "NOT_AUTHORIZED\n",
			expectedCode: 403,
		}, {
			description:  "missing arg password",
			url:          "/ping?user=test",
			expectedBody: "MISSING_ARG_PASSWORD\n",
			expectedCode: 400,
		}, {
			description:  "bad password",
			url:          "/ping?user=test&password=badpassword",
			expectedBody: "NOT_AUTHORIZED\n",
			expectedCode: 403,
		},
	}

	ts := httptest.NewServer(Auth(GetTestHandler()))
	defer ts.Close()

	for _, tc := range tests {
		var u bytes.Buffer
		u.WriteString(string(ts.URL))
		u.WriteString(tc.url)

		res, err := http.Get(u.String())
		assert.NoError(err)
		if res != nil {
			defer res.Body.Close()
		}

		b, err := ioutil.ReadAll(res.Body)
		assert.NoError(err)

		assert.Equal(tc.expectedCode, res.StatusCode, tc.description)
		assert.Equal(tc.expectedBody, string(b), tc.description)
	}
}