package commonv1_test

import (
	"testing"

	commonv1 "github.com/mcrgnt/proto/gen/go/common/v1"
	"github.com/mcrgnt/proto/test/go/testutils"
)

func TestLabel_Validation(t *testing.T) {
	tests := []struct {
		name   string
		val    string
		passed bool
	}{
		{"valid_simple", "env", true},
		{"valid_with_numbers", "stage_2", true},
		{"valid_max_length", "this_is_a_very_long_label_32_ch", true},
		{"invalid_starts_with_number", "1label", false},
		{"invalid_caps", "Label", false},
		{"invalid_special_chars", "my-label", false}, // тире запрещено, только подчеркивание
		{"invalid_too_short", "", false},
		{"invalid_too_long", "label_that_is_definitely_longer_than_32_characters", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testutils.ValidateCase(t, &commonv1.Label{Value: tt.val}, tt.passed)
		})
	}
}
