package testutils

import (
	"testing"

	"buf.build/go/protovalidate"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func ValidateCase(t *testing.T, msg protoreflect.ProtoMessage, shouldPass bool) {
	v, err := protovalidate.New()
	require.NoError(t, err)

	err = v.Validate(msg)
	if shouldPass {
		require.NoError(t, err, "Message should be valid: %v", msg)
	} else {
		require.Error(t, err, "Message should be invalid: %v", msg)
	}
}
