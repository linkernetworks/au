package client

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"bitbucket.org/linkernetworks/aurora/src/cmd/au/util/httpclient"
	pb "gopkg.in/cheggaaa/pb.v1"
)

//FileUploader - file uploader
type FileUploader struct {
	Input  chan string
	HasSSL bool
	WsID   string
}

//Run - run uploder worker
func (f *FileUploader) Run() {
	var filePath string
	for filePath = <-f.Input; filePath != ""; filePath = <-f.Input {
		ret, err := uploadFileToWorkspace(f.HasSSL, f.WsID, filePath)
		if err != nil {
			log.Println(" uploadFileToWorkspace err:", err)
		}
		printResult(ret)
	}
}

//uploadFileToWorkspace - upload simple file to server by specific ws ID
func uploadFileToWorkspace(hasSSL bool, wsID string, filePath string) ([]byte, error) {
	client := httpclient.NewHttpClient(hasSSL)

	params := make(map[string]string)
	file, err := os.Open(filePath)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	params = map[string]string{
		"name": file.Name(),
	}

	return client.ConnectWithMultpart(GetUploadImageURL(wsID), "POST", params, "upload", filePath)
}

func recursiveWalkerFileUpload(wsID string, filePath string, bar *pb.ProgressBar, fileChan chan string) error {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return err
	}

	filepath.Walk(filePath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		fileChan <- path
		bar.Increment()
		return err
	})
	return nil
}

//CountFilesInDir - counting all files number in path
func CountFilesInDir(path string) (int, error) {
	i := 0
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return 0, err
	}
	for _, file := range files {
		if !file.IsDir() {
			i++
		}
	}
	return i, nil
}
