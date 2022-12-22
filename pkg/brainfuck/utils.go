package brainfuck

import (
	"io"
	"os"
)

func OpenFile(filename string) (io.Reader, error) {
	return os.Open(filename)
}
