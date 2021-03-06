package mockinterface

import (
	"testing"
	"io"
	"errors"
)

type ReaderMock struct {
	ReadMock func([]byte) (int, error)
}

func (m ReaderMock) Read(p []byte) (int, error) {
	return m.ReadMock(p)
}

var _ io.Reader = (*ReaderMock)(nil)

func TestReadN_bufSize(t *testing.T) {
	total := 0
	mr := &ReaderMock{func(b []byte) (int, error) {
		total = len(b)
		return 0, nil
	}}
	readN(mr, 5)
	if total != 5 {
		t.Fatalf("expected 5, got %d", total)
	}
}

func TestReadN_error(t *testing.T) {
	expect := errors.New("some non-nil error")
	mr := &ReaderMock{func(b []byte) (int, error) {
		return 0, expect
	}}
	_, err := readN(mr, 5)
	if err != expect {
		t.Fatal("expected error")
	}
}
