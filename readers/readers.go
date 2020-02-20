package readers

import "golang.org/x/tour/reader"

type MyReader struct{}

func (r MyReader) Read(b []byte) (int, error) {
	b[0] = 'A'
	return 1, nil
}

func run() {
	print("7. READERS: ")
	reader.Validate(MyReader{})
}

func Run() {
	run()
}
