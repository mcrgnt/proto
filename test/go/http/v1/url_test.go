package httpv1_test

import (
	"testing"

	httpv1 "github.com/omcrgnt/proto/gen/go/http/v1"
	"github.com/omcrgnt/proto/test/go/testutils"
)

func TestURL_Validation(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		val    string
		passed bool
	}{
		{"empty", "", false},
		{"http", "http://127.0.0.1:9000", true},
		{"https", "https://s3.amazonaws.com", true},
		{"ws", "ws://127.0.0.1/ws", false},
		{"file", "file:///tmp/x", false},
		{"no_scheme", "127.0.0.1:9000", false},
		{"relative", "/path", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			testutils.ValidateCase(t, &httpv1.URL{Value: tt.val}, tt.passed)
		})
	}
}
