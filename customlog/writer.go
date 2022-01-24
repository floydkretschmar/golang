package customlog

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type LogWriter struct {
}

func (LogWriter) Write(data []byte) (int, error) {
	return fmt.Println(string(data))
}

func HttpInterface() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		log.Fatal(err)
	}

	//data := make([]byte, 10000)
	//data, err = ioutil.ReadAll(resp.Body)
	//if err != nil {
	//	customlog.Fatal(err)
	//}
	//
	//body := string(data)
	//fmt.Println(body)
	logger := LogWriter{}
	io.Copy(logger, resp.Body)
}
