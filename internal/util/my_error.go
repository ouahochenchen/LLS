package util

type MyError struct {
	MyErrorMessage string
}

func (m *MyError) Error() string {
	return m.MyErrorMessage
}
