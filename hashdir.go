package hashdir

import (
	"crypto/md5"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// Create an md5 hash with local path
func Create(path string) (string, error) {

	md5Hash := md5.New()

	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil
		}
		io.WriteString(md5Hash, path)
		return nil
	})

	if err != nil {
		return "", nil
	}

	return fmt.Sprintf("%x", md5Hash.Sum(nil)), nil
}
