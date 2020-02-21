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
		switch {
		case b[i] >= 'a' && b[i] <= 'z':
			b[i] = 'a' + (b[i]-'a'+13)%26
		case b[i] >= 'A' && b[i] <= 'Z':
			b[i] = 'A' + (b[i]-'A'+13)%26
		}
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
