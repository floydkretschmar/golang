package customlog

import "fmt"

type LogWriter struct {
}

func (LogWriter) Write(data []byte) (int, error) {
	return fmt.Println(string(data))
}
