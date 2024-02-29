package document

import "os"

type Document struct {
	payload string
}

func DocumentInit(payload string) *Document {
	return &Document{
		payload,
	}
}
//  https://stackoverflow.com/questions/33851692/golang-bad-file-descriptor
func (d *Document) CreateFile() error {
	filePointer, err := os.OpenFile("dataFromServer.txt", os.O_APPEND|os.O_CREATE | os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	_, err = filePointer.Write([]byte(d.payload))

	if err != nil {
		return err
	}

	return nil
}