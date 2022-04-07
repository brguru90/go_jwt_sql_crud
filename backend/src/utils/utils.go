package utils

import (
	log "github.com/sirupsen/logrus"
	"io"
	"os"
)

func CopyFIle(_source string, _destination string) error {
	file1, err := os.Open(_source)
	if err != nil {
		log.WithFields(log.Fields{
			"err":err,
		}).Errorln("Error in opening file");
		return err
	}
	defer file1.Close()
	fileinfo, err := file1.Stat()
	if err != nil {
		log.WithFields(log.Fields{
			"err":err,
		}).Errorln("Error in getting file properties");
		return err
	}

	file2, err := os.Open(_destination)
	if err != nil {
		log.Errorln(err)
		return err
	}
	defer file2.Close()

	BUFFERSIZE := fileinfo.Size()

	buf := make([]byte, BUFFERSIZE);
	for {
		n, err := file1.Read(buf)
		if err != nil && err != io.EOF {
			return err
		}
		if n == 0 {
			break
		}

		if _, err := file2.Write(buf[:n]); err != nil {
			return err
		}
	}
	return nil
}
