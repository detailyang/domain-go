package domain

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsValid(t *testing.T) {
	for _, ct := range []struct {
		domain string
		valid  bool
	}{
		{
			"www.baidu.com",
			true,
		},
	} {
		require.Equal(t, ct.valid, IsValid(ct.domain))
	}
}
