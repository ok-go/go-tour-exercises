package rot13reader

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (r rot13Reader) Read(b []byte) (n int, err error) {
	if n, err = r.r.Read(b); err != nil {
		return
	}

	for i := 0; i < n; i++ {
		var s, e byte
		if b[i] >= 'a' && b[i] <= 'z' {
			s = 'a'
			e = 'z' - s
		} else if b[i] >= 'A' && b[i] <= 'Z' {
			s = 'A'
			e = 'Z' - s
		} else {
			continue
		}

		b[i] = ((b[i] - s + 13) % (e + 1)) + s
	}

	return
}

func run() {
	print("8. ROT13READER:")
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
	println()
}

func Run() {
	run()
}
