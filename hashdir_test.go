package hashdir_test

import (
	"testing"

	"github.com/sger/go-hashdir"
	"github.com/stretchr/testify/require"
)

func TestHashDir(t *testing.T) {
	hash1, err := hashdir.Create("test/directory", "md5")
	require.NoError(t, err)

	hash2, err := hashdir.Create("test/directory", "md5")
	require.NoError(t, err)

	require.Equal(t, hash1, hash2, "equals")
}
