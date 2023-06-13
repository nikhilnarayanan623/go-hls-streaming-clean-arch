package utils

import (
	"io"
)

func CopyFilesFormSrcToDest(src io.Reader, dest io.Writer) error {

	_, err := io.Copy(dest, src)

	return err
}
