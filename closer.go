package encoding

import "io"

var _ io.ReadCloser = (*readCloser)(nil)

type readCloser []io.Reader

func (r readCloser) Read(p []byte) (n int, err error) {
	return r[0].Read(p)
}

func (r readCloser) Close() error {
	for i := len(r) - 1; i >= 0; i-- {
		if closer, ok := r[i].(io.Closer); ok {
			if err := closer.Close(); err != nil {
				return err
			}
		}
	}
	return nil
}

var _ io.WriteCloser = (*writeCloser)(nil)

type writeCloser []io.Writer

func (w writeCloser) Write(p []byte) (n int, err error) {
	return w[0].Write(p)
}

func (w writeCloser) Close() error {
	for _, wr := range w {
		if closer, ok := wr.(io.Closer); ok {
			if err := closer.Close(); err != nil {
				return err
			}
		}
	}
	return nil
}
