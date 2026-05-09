package loggerv1_test

import (
	"testing"

	loggerv1 "github.com/omcrgnt/proto/gen/go/logger/v1"
	"github.com/omcrgnt/proto/test/go/testutils"
)

func TestLoggerLevel_Validation(t *testing.T) {
	tests := []struct {
		name   string
		level  string
		passed bool
	}{
		{"valid_info", "info", true},
		{"valid_error", "error", true},
		{"invalid_caps", "DEBUG", false}, // не пройдет, так как в enum только строчные
		{"invalid_value", "fatal", false},
		{"empty", "", false},
		{"uppercase_fail", "INFO", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Используем новое имя LoggerLevel
			testutils.ValidateCase(t, &loggerv1.LoggerLevel{Value: tt.level}, tt.passed)
		})
	}
}
