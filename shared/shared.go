package shared

import (
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sync"
)

// This is the one improved performance
// ReadBufferSize shouldn't be smaller than WriteBufferSize!
func ReadRequestBodyAndWriteToFileImproved(httpRequest *http.Request, requestBodyReadBufferSize uint32, fileWriteBufferSizeParam uint32, filePath string, wg *sync.WaitGroup) {
	defer wg.Done()
	fileToWrite1, err := os.OpenFile(filePath, os.O_RDWR|os.O_APPEND, 0660)
	if err != nil {
		log.Println("Error occurred while openin file", err)
	}
	defer fileToWrite1.Close()
	requestBodyReadBuffer := make([]byte, requestBodyReadBufferSize)
	var fileWriteBufferSize uint32 = fileWriteBufferSizeParam
	for {
		byteCount1, err := httpRequest.Body.Read(requestBodyReadBuffer)
		if err != nil {
			if err == io.EOF {
				if byteCount1 < int(fileWriteBufferSizeParam) {
					fileToWrite1.Write(requestBodyReadBuffer[:byteCount1])
					break
				}
				break
			}
			log.Println("Error reading request body:", err)
			return
		}
		for i := uint32(0); i < uint32(byteCount1); i += fileWriteBufferSize {
			if uint32(byteCount1)-i < fileWriteBufferSize {
				fileToWrite1.Write(requestBodyReadBuffer[i:])
				break
			}
			fileToWrite1.Write(requestBodyReadBuffer[i : i+fileWriteBufferSize])
		}
		//log.Println(string(requestBodyReadBuffer[:byteCount1]))
	}
}

// This method is for benchmark test for go routines
func ReadRequestBodyAndWriteToFileImprovedMultipleFilesWithRoutine(httpRequest *http.Request, requestBodyReadBufferSize uint32, fileWriteBufferSizeParam uint32, filePath string, wg *sync.WaitGroup) {
	defer wg.Done()
	x1, _ := os.Getwd()
	fileToWrite1, err := os.OpenFile(filepath.Join(x1, "..", "static_files", "text2.txt"), os.O_RDWR|os.O_APPEND, 0660)
	fileToWrite2, err := os.OpenFile(filepath.Join(x1, "..", "static_files", "text3.txt"), os.O_RDWR|os.O_APPEND, 0660)
	fileToWrite3, err := os.OpenFile(filepath.Join(x1, "..", "static_files", "text4.txt"), os.O_RDWR|os.O_APPEND, 0660)
	fileToWrite4, err := os.OpenFile(filepath.Join(x1, "..", "static_files", "text5.txt"), os.O_RDWR|os.O_APPEND, 0660)
	if err != nil {
		log.Println("Error occurred while openin file", err)
	}
	defer fileToWrite1.Close()
	requestBodyReadBuffer := make([]byte, requestBodyReadBufferSize)
	var fileWriteBufferSize uint32 = fileWriteBufferSizeParam
	for {
		byteCount1, err := httpRequest.Body.Read(requestBodyReadBuffer)
		if err != nil {
			if err == io.EOF {
				if byteCount1 < 5 {
					fileToWrite1.Write(requestBodyReadBuffer[:byteCount1])
					fileToWrite2.Write(requestBodyReadBuffer[:byteCount1])
					fileToWrite3.Write(requestBodyReadBuffer[:byteCount1])
					fileToWrite4.Write(requestBodyReadBuffer[:byteCount1])
					break
				}
				break
			}
			log.Println("Error reading request body:", err)
			return
		}
		for i := uint32(0); i < uint32(byteCount1); i += fileWriteBufferSize {
			if uint32(byteCount1)-i < fileWriteBufferSize {
				fileToWrite1.Write(requestBodyReadBuffer[i:])
				fileToWrite2.Write(requestBodyReadBuffer[i:])
				fileToWrite3.Write(requestBodyReadBuffer[i:])
				fileToWrite4.Write(requestBodyReadBuffer[i:])
				break
			}
			fileToWrite1.Write(requestBodyReadBuffer[i : i+fileWriteBufferSize])
			fileToWrite2.Write(requestBodyReadBuffer[i : i+fileWriteBufferSize])
			fileToWrite3.Write(requestBodyReadBuffer[i : i+fileWriteBufferSize])
			fileToWrite4.Write(requestBodyReadBuffer[i : i+fileWriteBufferSize])
		}
		//log.Println(string(requestBodyReadBuffer[:byteCount1]))
	}
}

func ReadRequestBodyAndWriteToFileImprovedMultipleFilesWithoutRoutines(httpRequest *http.Request, requestBodyReadBufferSize uint32, fileWriteBufferSizeParam uint32, filePath string) {
	x1, _ := os.Getwd()
	fileToWrite1, err := os.OpenFile(filepath.Join(x1, "..", "static_files", "text2.txt"), os.O_RDWR|os.O_APPEND, 0660)
	fileToWrite2, err := os.OpenFile(filepath.Join(x1, "..", "static_files", "text3.txt"), os.O_RDWR|os.O_APPEND, 0660)
	fileToWrite3, err := os.OpenFile(filepath.Join(x1, "..", "static_files", "text4.txt"), os.O_RDWR|os.O_APPEND, 0660)
	fileToWrite4, err := os.OpenFile(filepath.Join(x1, "..", "static_files", "text5.txt"), os.O_RDWR|os.O_APPEND, 0660)
	if err != nil {
		log.Println("Error occurred while openin file", err)
	}
	defer fileToWrite1.Close()
	requestBodyReadBuffer := make([]byte, requestBodyReadBufferSize)
	var fileWriteBufferSize uint32 = fileWriteBufferSizeParam
	for {
		byteCount1, err := httpRequest.Body.Read(requestBodyReadBuffer)
		if err != nil {
			if err == io.EOF {
				if byteCount1 < 5 {
					fileToWrite1.Write(requestBodyReadBuffer[:byteCount1])
					fileToWrite2.Write(requestBodyReadBuffer[:byteCount1])
					fileToWrite3.Write(requestBodyReadBuffer[:byteCount1])
					fileToWrite4.Write(requestBodyReadBuffer[:byteCount1])
					break
				}
				break
			}
			log.Println("Error reading request body:", err)
			return
		}
		for i := uint32(0); i < uint32(byteCount1); i += fileWriteBufferSize {
			if uint32(byteCount1)-i < fileWriteBufferSize {
				fileToWrite1.Write(requestBodyReadBuffer[i:])
				fileToWrite2.Write(requestBodyReadBuffer[i:])
				fileToWrite3.Write(requestBodyReadBuffer[i:])
				fileToWrite4.Write(requestBodyReadBuffer[i:])
				break
			}
			fileToWrite1.Write(requestBodyReadBuffer[i : i+fileWriteBufferSize])
			fileToWrite2.Write(requestBodyReadBuffer[i : i+fileWriteBufferSize])
			fileToWrite3.Write(requestBodyReadBuffer[i : i+fileWriteBufferSize])
			fileToWrite4.Write(requestBodyReadBuffer[i : i+fileWriteBufferSize])
		}
		//log.Println(string(requestBodyReadBuffer[:byteCount1]))
	}
}

func ReadRequestBodyAndWriteToFile(httpRequest *http.Request, requestBodyReadBufferSize uint32, fileWriteBufferSize uint32, filePath string) {
	readBufferSize := requestBodyReadBufferSize
	requestBodyReadBuffer := make([]byte, readBufferSize)
	var readStartIndex uint32 = 0
	fileToWrite1, err := os.OpenFile(filePath, os.O_RDWR|os.O_APPEND, 0660)
	if err != nil {
		log.Println("Error occurred while openin file", err)
	}
	for uint32(httpRequest.ContentLength) > readStartIndex {
		readByteCount, err := httpRequest.Body.Read(requestBodyReadBuffer)
		if err != nil {
			if err == io.EOF {
				bufferSize := readByteCount
				buffer := make([]byte, bufferSize)
				startingByteIndex := 0
				for len(requestBodyReadBuffer[:readByteCount]) > startingByteIndex {
					if len(requestBodyReadBuffer[:readByteCount])-startingByteIndex < bufferSize {
						bufferSize = len(requestBodyReadBuffer[:readByteCount]) - startingByteIndex
					}
					for i := 0; i < bufferSize; i++ {
						buffer[i] = requestBodyReadBuffer[:readByteCount][startingByteIndex]
						startingByteIndex++
					}
					_, err = fileToWrite1.Write(buffer)
					if err != nil {
						panic(err)
					}
					buffer = make([]byte, bufferSize)
				}
				break // Reached the end of the request body
			}
			panic(err)
		}
		wrifeFileBufferSize := fileWriteBufferSize
		writeFileBuffer := make([]byte, wrifeFileBufferSize)
		var writeStartIndex uint32 = 0
		for uint32(len(requestBodyReadBuffer[:readByteCount])) > writeStartIndex {
			if uint32(len(requestBodyReadBuffer[:readByteCount]))-writeStartIndex < wrifeFileBufferSize {
				wrifeFileBufferSize = uint32(len(requestBodyReadBuffer[:readByteCount])) - writeStartIndex
			}
			for i := uint32(0); i < wrifeFileBufferSize; i++ {
				writeFileBuffer[i] = requestBodyReadBuffer[:readByteCount][writeStartIndex]
				writeStartIndex++
			}
			_, err = fileToWrite1.Write(writeFileBuffer)
			if err != nil {
				panic(err)
			}
			writeFileBuffer = make([]byte, wrifeFileBufferSize)
		}
		readStartIndex += uint32(readByteCount)
	}
}
