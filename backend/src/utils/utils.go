package utils

import (
	"io"
	"os"
)

func CopyFIle(_source string, _destination string) (string,error) {
	file1, err := os.Open(_source)
	if err != nil {
		return "Error in opening source file",err
	}
	defer file1.Close()
	fileinfo, err := file1.Stat()
	if err != nil {
		return "Error in getting source file properties",err
	}

	file2, err := os.OpenFile(_destination,os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return  "Error in opening/creating destination file",err
	}
	defer file2.Close()

	BUFFERSIZE := fileinfo.Size()

	buf := make([]byte, BUFFERSIZE);
	for {
		n, err := file1.Read(buf)
		if err != nil && err != io.EOF {
			return "Error in reading sourcefile buffer",err
		}
		if n == 0 {
			break
		}

		if _, err := file2.Write(buf[:n]); err != nil {
			return "Error in writing destination buffer",err
		}
	}
	return "",nil
}
