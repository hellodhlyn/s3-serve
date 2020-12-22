package sources

import (
	"io"
)

type Source interface {
	Read(key string) (io.Reader, error)
	Write(key string, body io.Reader) error
}
