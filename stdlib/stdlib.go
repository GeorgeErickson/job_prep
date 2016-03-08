package stdlib

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func ReadFileByLine(fn string) error {
	f, err := os.Open(fn)
	if err != nil {
		return err
	}
	defer f.Close()
	return ReadByLine(f)

}

func ReadByLine(r io.Reader) error {
	s := bufio.NewScanner(r)

	for s.Scan() {
		fmt.Println(s.Text())
	}

	return s.Err()
}
