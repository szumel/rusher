package macro

import "io"

const expectedSchema Schema = `<?xml version="1.0"?>
<macro v="1.0.0">
    <step code="printPwd"/>
    <step code="changeCwd" dir="/xx"/>
</macro>`

type filesystemMock struct {
	file file
}

func (f *filesystemMock) open(name string) (file, error) {
	return f.file, nil
}

type fileMock struct {
	data []byte
	done bool
}

func (f *fileMock) Read(p []byte) (n int, err error) {
	if f.done {
		return 0, io.EOF
	}
	for i, b := range []byte(f.data) {
		p[i] = b
	}

	f.done = true

	return len(p), nil
}

func (*fileMock) Close() error {
	return nil
}
