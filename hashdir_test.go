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

	hash3, err := hashdir.Create("test/directory1", "sha1")
	require.NoError(t, err)

	hash4, err := hashdir.Create("test/directory1", "sha1")
	require.NoError(t, err)

	require.Equal(t, hash3, hash4, "hash3 and hash 4 are equals")
}
