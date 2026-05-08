package loggerv1_test

import (
	"testing"

	loggerv1 "github.com/omcrgnt/proto/gen/go/logger/v1"
	"github.com/omcrgnt/proto/test/go/testutils"
)

func TestFormat_Validation(t *testing.T) {
	tests := []struct {
		name   string
		format string
		passed bool
	}{
		{"valid_json", "json", true},
		{"valid_text", "text", true},
		{"invalid_format", "yaml", false},
		{"empty", "", false},
		{"uppercase_fail", "JSON", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testutils.ValidateCase(t, &loggerv1.Format{Value: tt.format}, tt.passed)
		})
	}
}
