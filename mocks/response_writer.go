package mocks

import "net/http"

func NewMockWriter() *MockWriter {
	mockWriter := &MockWriter{}
	mockWriter.Mock.Header = make(http.Header)
	return mockWriter
}

type MockWriter struct {
	Assert struct {
		Bytes  []byte
		Status int
	}
	Mock struct {
		Write struct {
			BytesLength int
			Error       error
		}
		Header http.Header
	}
}

func (w *MockWriter) Header() http.Header {
	return w.Mock.Header
}

func (w *MockWriter) Write(bytes []byte) (int, error) {
	w.Assert.Bytes = bytes

	length := len(bytes)
	if w.Mock.Write.BytesLength > 0 {
		length = w.Mock.Write.BytesLength
	}

	return length, w.Mock.Write.Error
}

func (w *MockWriter) WriteHeader(statusCode int) {
	w.Assert.Status = statusCode
}
