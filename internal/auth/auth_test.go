package auth

import (
	"fmt"
	"net/http"
	"strings"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		key       string
		value     string
		expect    string
		expectErr string
	}{
		{
			expectErr: "no authorization header included",
			expect:    "", // Explicitly validate empty output on error.
		},
		{
			key:       "Authorization",
			expectErr: "no authorization header included",
			expect:    "", // Empty-value headers should also expect empty output.
		},
		{
			key:       "Authorization",
			value:     "-",
			expectErr: "malformed authorization header",
			expect:    "", // Explicitly validate output on error.
		},
		{
			key:       "Authorization",
			value:     "Bearer xxxxxx",
			expectErr: "malformed authorization header",
			expect:    "", // Explicitly validate output on error.
		},
		{
			key:       "Authorization",
			value:     "ApiKey xxxxxx",
			expect:    "xxxxxx",
			expectErr: "",
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("TestGetAPIKey Case #%v:", i), func(t *testing.T) {
			header := http.Header{}
			header.Add(test.key, test.value)

			output, err := GetAPIKey(header)
			if err != nil {
				if strings.Contains(err.Error(), test.expectErr) {
					return
				}
				t.Errorf("Unexpected: TestGetAPIKey:%v\n", err)
				return
			}

			if output != test.expect {
				t.Errorf("Unexpected: TestGetAPIKey:%s", output)
				return
			}
			if err != nil {
				// Validate if the error message contains the expected error.
				if !strings.Contains(err.Error(), test.expectErr) {
					t.Errorf("Unexpected error: got %v, want %v", err, test.expectErr)
				}
				// Also validate that the output is empty when there's an error.
				if output != test.expect {
					t.Errorf("Unexpected output: got %s, expected %s", output)
				}
			}
		})
	}
}
