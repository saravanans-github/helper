package fs

import (
	"io/ioutil"
	"testing"
)

func TestWriteToFile(t *testing.T) {
	dataToWrite := []string{"test", "this", "now"}
	dataChannel := make(chan []byte)

	var err error
	var isDone = false
	WriteToFile("test.txt", dataChannel, &err, &isDone)
	for _, data := range dataToWrite {
		dataChannel <- []byte(data)
	}
	close(dataChannel)
	for !isDone {
		// wait
	}

	dat, err := ioutil.ReadFile("test.txt")

	if string(dat) != "testthisnow" || err != nil {
		t.FailNow()
	}
}
