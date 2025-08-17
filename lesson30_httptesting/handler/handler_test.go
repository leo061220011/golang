package handler

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/ozontech/cute"
	"github.com/ozontech/cute/asserts/json"
)

func TestStatusHandler(t *testing.T) {
	// 	type want struct {
	// 		code        int
	// 		response    string
	// 		contentType string
	// 	}
	// 	tests := []struct {
	// 		name string
	// 		want want
	// 	}{
	// 		// TODO: Add test cases.
	// 		{
	// 			name: "Simple test #1",
	// 			want: want{
	// 				code:        200,
	// 				response:    `{"status": "ok"}`,
	// 				contentType: "application/json",
	// 			},
	// 		},
	// 	}

	// 	for _, tt := range tests {
	// 		t.Run(tt.name, func(t *testing.T) {
	// 			request := httptest.NewRequest(http.MethodGet, "/status", nil)
	// 			w := httptest.NewRecorder()
	// 			StatusHandler(w, request)
	// 			res := w.Result()
	// 			assert.Equal(t, tt.want.code, res.StatusCode)
	// 			defer res.Body.Close()
	// 			resBody, err := io.ReadAll(res.Body)
	// 			require.NoError(t, err)
	// 			assert.JSONEq(t, tt.want.response, string(resBody))
	// 			assert.Equal(t, tt.want.contentType, res.Header.Get("Content-Type"))

	// 		})

	// 	}

	testCases := []struct {
		name     string
		url      string
		expected map[string]interface{}
	}{
		{
			name: "Simple Status test",
			url:  "/status",
			expected: map[string]interface{}{
				"status": "ok",
			},
		},
		{
			name: "Simple Healthy Test",
			url:  "/healthy",
			expected: map[string]interface{}{
				"healthy": true,
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			cute.NewTestBuilder().
				Title(tc.name).
				Create().
				RequestBuilder(
					cute.WithURI("http://127.0.0.1:8080"+tc.url),
					cute.WithMethod(http.MethodGet),
				).ExpectExecuteTimeout(10*time.Second).
				ExpectStatus(http.StatusOK).
				AssertBody(
					json.Equal("$", tc.expected),
				).ExecuteTest(context.Background(), t)
		})
	}

}
