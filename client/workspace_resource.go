package client

import (
	"fmt"

	pb "gopkg.in/cheggaaa/pb.v1"
	"gopkg.in/mgo.v2/bson"
)

type WorkspaceWebRequest struct {
	DatasetIds []bson.ObjectId `bson:"datasetIds" json:"datasetIds"`
	Type       string          `bson:"type" json:"type"`
	NumItems   int             `bson:"numItems" json:"numItems"`
}

type WorkspaceResource struct {
	hasSSL       bool
	WsID         string
	fileUploader chan string
}

//UploadFile -
func (ws *WorkspaceResource) UploadFile(filepathes []string) {
	//ws is not creat/find correctly, just return.
	if ws.WsID == "" {
		return
	}
	uploader := ws.NewFileUploader()
	go uploader.Run()

	var count int
	for _, v := range filepathes {
		if num, err := CountFilesInDir(v); err == nil {
			count = count + num
		}
	}
	fmt.Println("Find total files:", count)
	// Remember total files
	bar := pb.StartNew(count)

	for _, v := range filepathes {
		recursiveWalkerFileUpload(ws.WsID, v, bar, uploader.Input)
	}

	bar.FinishPrint("Upload Completed")
	close(uploader.Input)
}

//NewFileUploader -
func (ws *WorkspaceResource) NewFileUploader() *FileUploader {
	return &FileUploader{HasSSL: ws.hasSSL, WsID: ws.WsID}
}

//DownloadZip -
func (ws *WorkspaceResource) DownloadZip(filepathes ...string) error {
	//TODO: Need wait until process completed.
	return nil
}

func (ws *WorkspaceResource) DownloadTar(filepathes ...string) error {
	//TODO: Need wait until process completed.
	return nil
}
