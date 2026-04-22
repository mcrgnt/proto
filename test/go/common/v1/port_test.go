package commonv1_test

import (
	"testing"

	commonv1 "github.com/mcrgnt/proto/gen/go/common/v1"
	"github.com/mcrgnt/proto/test/go/testutils"
)

func TestPort_Validation(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		val    uint32
		passed bool
	}{
		{"min_port", 1, true},
		{"max_port", 65535, true},
		{"invalid_zero", 0, false},
		{"out_of_range", 70000, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testutils.ValidateCase(t, &commonv1.Port{Value: tt.val}, tt.passed)
		})
	}
}
