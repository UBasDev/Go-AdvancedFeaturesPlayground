package controllers

import (
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"gorm.io/gorm"
)

func AddTestControllers(router *http.ServeMux, dbContext *gorm.DB) {
	router.HandleFunc("/api/test1", func(rw http.ResponseWriter, rq *http.Request) {
		x1, _ := os.Getwd()
		staticFilePath := filepath.Join(x1, "..", "static_files", "text2.txt")
		// x2, err := os.OpenFile(staticFilePath, os.O_RDWR|os.O_APPEND, 0660)
		// if err != nil {
		// 	log.Panicln("Some error happend about opening file", err)
		// }
		// defer x2.Close()
		// directoryNames, _ := x2.Readdirnames(-1)
		// for _, currentName := range directoryNames {
		// 	log.Println("Current directory name is: ", currentName)
		// }
		// currentOffset1, _ := x2.Seek(4, 0)
		//log.Println("Current offset is: ", currentOffset1)
		//byteToRead := make([]byte, 1024)
		//err = x2.SetDeadline(time.Now().Add(5 * time.Second))
		//err = x2.SetReadDeadline(time.Now().Add(5 * time.Second))
		//err = x2.SetWriteDeadline(time.Now().Add(5 * time.Second))
		// if err != nil {
		// 	log.Panicln("Some error happend about deadline", err)
		// }
		// for {
		// 	byteReadCount, err := x2.Read(byteToRead)
		// 	if err != nil {
		// 		if err == io.EOF {
		// 			break
		// 		}
		// 		log.Panicln("Some error happend about reading file", err)
		// 	}
		// 	if byteReadCount < 1024 {
		// 		log.Println("Read byte count is:", byteReadCount)
		// 		log.Println("End of the read cycle.", string(byteToRead[:byteReadCount]))
		// 	}
		// }
		// defer rq.Body.Close()
		// wg := sync.WaitGroup{}
		// wg.Add(1)
		// go shared.ReadRequestBodyAndWriteToFileImproved(rq, 10, 6, staticFilePath, &wg)
		// wg.Wait()
		// requestBody, err := io.ReadAll(rq.Body)
		// if err != nil {
		// 	log.Println("Error occurred while reading request body", err)
		// }
		someContent := []byte{'a', 'h', 'm', 'e', 't', '1', ' ', 'v', 'e', ' ', 'm', 'e', 'h', 'm', 'e', 't', '1'}
		fileToWrite1, err := os.OpenFile(staticFilePath, os.O_RDWR|os.O_APPEND, 0660)
		if err != nil {
			log.Println("Error occurred while openin file", err)
		}
		var readStartIndex uint32 = 0
		var bufferSize uint32 = 5
		for {
			if uint32(len(someContent)) < readStartIndex+bufferSize {
				byteCount, err := fileToWrite1.Write(someContent[readStartIndex:])
				if err != nil {
					log.Println("Some error occurred while writing the file: ", err)
				}
				log.Println("Write byte count: ", byteCount)
				break
			}
			byteCount, err := fileToWrite1.Write(someContent[readStartIndex : readStartIndex+bufferSize])
			if err != nil {
				log.Println("Some error occurred while writing the file: ", err)
			}
			readStartIndex += uint32(byteCount)
			log.Println("Write byte count: ", byteCount)
		}
		err = fileToWrite1.Sync()
		if err != nil {
			log.Println("Some error occurred while syncing the file: ", err)
		}
		defer fileToWrite1.Close()
	})
	router.HandleFunc("/api/test2", func(rw http.ResponseWriter, rq *http.Request) {
		x1, _ := os.Getwd()
		staticFilePath := filepath.Join(x1, "..", "static_files", "text2.txt")
		fileToWrite1, err := os.OpenFile(staticFilePath, os.O_RDWR, 0660)
		if err != nil {
			log.Println("Error occurred while openin file", err)
		}
		defer fileToWrite1.Close()
		conn1, err := fileToWrite1.SyscallConn()
		if err != nil {
			log.Println("Error occurred while sending syscall to os: ", err)
			return
		}
		conn1.Control(func(fd uintptr) {
			log.Println("File descriptor is:", fd)

		})
		err = fileToWrite1.Truncate(5)
		if err != nil {
			log.Println("Error occurred while truncating the file: ", err)
			return
		}
	})
	router.HandleFunc("/api/test3", func(rw http.ResponseWriter, rq *http.Request) {
		x1, _ := os.Getwd()
		staticFilePath := filepath.Join(x1, "..", "static_files", "text2.txt")
		fileToRead1, err := os.OpenFile(staticFilePath, os.O_RDWR, 0660)
		if err != nil {
			log.Println("Error occurred while openin file", err)
		}
		bufferSize := 1024
		buffer := make([]byte, bufferSize)
		for {
			byteCount, err := fileToRead1.Read(buffer)
			if err != nil {
				if err == io.EOF {
					log.Println("We reached the end of the file")
				}
				log.Println("Error occurred while reading the file")
				break
			}
			if byteCount < bufferSize {
				_, err = rw.Write(buffer[:byteCount])
				if err != nil {
					log.Println("Error occured while writing to response: ", err)
				}
				break
			}
			_, err = rw.Write(buffer)
			if err != nil {
				log.Println("Error occured while writing to response: ", err)
			}
		}
	})
	router.HandleFunc("/api/test4", func(rw http.ResponseWriter, rq *http.Request) {
		x1, _ := os.Getwd()
		staticFilePath := filepath.Join(x1, "..", "static_files", "text2.txt")
		staticFilePath2 := filepath.Join(x1, "..", "static_files", "text3.txt")
		fileToRead1, err := os.OpenFile(staticFilePath, os.O_RDWR, 0660)
		if err != nil {
			log.Println("Error occurred while openin file", err)
		}
		fileToWrite1, err := os.OpenFile(staticFilePath2, os.O_RDWR, 0660)
		if err != nil {
			log.Println("Error occurred while openin file", err)
		}
		bufferSize := 1024
		buffer := make([]byte, bufferSize)
		for {
			byteCount, err := fileToRead1.Read(buffer)
			if err != nil {
				if err == io.EOF {
					log.Println("We reached the end of the file")
				}
				log.Println("Error occurred while reading the file")
				break
			}
			if byteCount < bufferSize {
				_, err = fileToWrite1.Write(buffer[:byteCount])
				if err != nil {
					log.Println("Error occured while writing to response: ", err)
				}
				break
			}
			_, err = fileToWrite1.Write(buffer)
			if err != nil {
				log.Println("Error occured while writing to response: ", err)
			}
		}
	})
	router.HandleFunc("/api/test5", func(rw http.ResponseWriter, rq *http.Request) {
		x1, _ := os.Getwd()
		staticFilePath := filepath.Join(x1, "..", "static_files", "text2.txt")
		someContent := []byte{'a', 'h', 'm', 'e', 't', '1', ' ', 'v', 'e', ' ', 'm', 'e', 'h', 'm', 'e', 't', '1'}
		fileToWrite1, err := os.OpenFile(staticFilePath, os.O_RDWR, 0660)
		if err != nil {
			log.Println("Error occurred while openin file", err)
		}
		byteCount, err := fileToWrite1.WriteAt(someContent, 8)
		if err != nil {
			log.Println("Error occurred while writing the file: ", err)
		}
		log.Println("Byte write count: ", byteCount)
		fileToWrite1.WriteString("sadsa")
	})

}
