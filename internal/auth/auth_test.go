package auth

import (
	"fmt"
	"net/http"
	"strings"
	"testing"
)

func TestGetApiKey(t *testing.T) {
	tests := []struct {
		key         string
		val         string
		expected    string
		expectedErr string
	}{
		{
			expectedErr: "no authorization header included",
		},
		{
			key:         "Authorization",
			expectedErr: "no authorization header included",
		},
		{
			key:         "Authorization",
			val:         "1234567890",
			expectedErr: "malformed authorization header",
		},
		{
			key:         "Authorization",
			val:         "Bearer 1234567890",
			expectedErr: "malformed authorization header",
		},
		{
			key:         "Authorization",
			val:         "ApiKey 1234567890",
			expected:    "1234567890",
			expectedErr: "",
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("TestGetApiKey_%d", i), func(t *testing.T) {
			headers := http.Header{}
			headers.Add(test.key, test.val)

			apiKey, err := GetAPIKey(headers)
			if err != nil {
				if strings.Contains(err.Error(), test.expectedErr) {
					return
				}

				if test.expectedErr == "" {
					t.Fatalf("case %d: expected error %s, got %s", i, test.expectedErr, err)
					return
				}
				t.Fatalf("case %d: expected error %s, got %s", i, test.expectedErr, err)
				return
			}

			if apiKey != test.expected {
				t.Fatalf("case %d: expected %s, got %s", i, test.expected, apiKey)
				return
			}
		})
	}
}
