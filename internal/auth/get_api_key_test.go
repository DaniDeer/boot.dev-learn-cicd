package auth

import (
	"errors"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestGetAPIKey(t *testing.T) {

	// Table driven tests
	tests := map[string]struct {
		testHeader map[string][]string
		want       string
		error      error
	}{

		"valid API key": {
			testHeader: map[string][]string{
				"Authorization": {"ApiKey test-api-key"},
			},
			want:  "test-api-key",
			error: nil,
		},

		"missing Authorization header": {
			testHeader: map[string][]string{
				"No-Authorization": {"ApiKey test-api-key"},
			},
			want:  "",
			error: ErrNoAuthHeaderIncluded,
		},

		"malformed Authorization header": {
			testHeader: map[string][]string{
				"Authorization": {"Bearer test-api-key"},
			},
			want:  "",
			error: errors.New("malformed authorization header"),
		},
	}

	// Iterate over the test cases
	for name, tc := range tests {

		t.Run(name, func(t *testing.T) {
			got, err := GetAPIKey(tc.testHeader)

			// Using cmp.Diff to compare the expected and actual results
			diff := cmp.Diff(tc.want, got)

			if diff != "" {
				t.Errorf("mismatch (-want +got):\n%s", diff)
			}
			if err != nil && err.Error() != tc.error.Error() {
				t.Errorf("expected error: %s, got: %s", tc.error.Error(), err.Error())
			}
			if err != nil && tc.error == nil {
				t.Errorf("unexpected error: %v", err)
			}
		})
	}
}
