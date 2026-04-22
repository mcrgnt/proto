package commonv1_test

import (
	"testing"

	commonv1 "github.com/mcrgnt/proto/gen/go/common/v1"

	"github.com/mcrgnt/proto/test/go/testutils"
)

func TestHost_Validation(t *testing.T) {
	tests := []struct {
		name   string
		addr   string
		passed bool
	}{
		{"valid_ip", "127.0.0.1", true},
		{"valid_domain", "gitlab.com", true},
		{"empty_address", "", false},
		{"too_long_address", "://limits.com", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testutils.ValidateCase(t, &commonv1.Host{Address: tt.addr}, tt.passed)
		})
	}
}
