package client

import (
	"bytes"
	"encoding/json"
	"log"
	"strconv"

	"github.com/linkernetworks/aurora/src/cmd/au/util/httpclient"
	"github.com/linkernetworks/aurora/src/entity"
	"gopkg.in/mgo.v2/bson"
)

//WorkspaceResourceCollection - workspace resource collection.
type WorkspaceResourceCollection struct {
	hasSSL bool
}

//Workspaces -
func Workspaces(hasSSL bool) *WorkspaceResourceCollection {
	return &WorkspaceResourceCollection{hasSSL: hasSSL}
}

//Create - create workspace resource collection
func (w *WorkspaceResourceCollection) Create(wsType string, numItemPreview int, datasetIDs []string) (*WorkspaceResource, error) {
	client := httpclient.NewHttpClient(w.hasSSL)
	req := &WorkspaceWebRequest{
		Type:     wsType,
		NumItems: numItemPreview,
	}
	for _, v := range datasetIDs {
		req.DatasetIds = append(req.DatasetIds, bson.ObjectIdHex(v))
	}

	ret, err := json.Marshal(req)
	if err != nil {
		log.Println("command format err:", err)
		return nil, err
	}

	ret, err = client.Connect(CreateWorkspaceURL, "POST", nil, bytes.NewBuffer(ret))
	if err != nil {
		log.Println("error on create workspace:", err)
		return &WorkspaceResource{}, err
	}

	printResult(ret)
	wsResult := entity.Workspace{}
	var ws *WorkspaceResource
	if err := json.Unmarshal(ret, &wsResult); err != nil {
		log.Println("error during unmarshal, err:", err)
	}
	ws = &WorkspaceResource{
		hasSSL: w.hasSSL,
		WsID:   wsResult.ID.Hex(),
	}
	return ws, nil
}

//Delete - delete target
func (w *WorkspaceResourceCollection) Delete(wsID string) error {
	client := httpclient.NewHttpClient(w.hasSSL)
	ret, err := client.Connect(GetWorkspaceURL(wsID), "DELETE", nil, nil)
	printResult(ret)
	return err
}

//Find -
func (w *WorkspaceResourceCollection) Find(wsID string) *WorkspaceResource {
	client := httpclient.NewHttpClient(w.hasSSL)
	ret, err := client.Connect(GetWorkspaceURL(wsID), "GET", nil, nil)
	printResult(ret)
	if err != nil {
		return &WorkspaceResource{}
	}
	ws := &WorkspaceResource{
		hasSSL: w.hasSSL,
		WsID:   wsID,
	}
	return ws
}

//Browse -
func (w *WorkspaceResourceCollection) Browse(page int, fileString string) error {
	client := httpclient.NewHttpClient(w.hasSSL)

	//prepare query string
	params := make(map[string]string)
	params["page"] = strconv.Itoa(page)
	params["filter"] = fileString

	ret, err := client.Connect(QueryWorkspaceURL, "GET", params, nil)
	if err == nil {
		printResult(ret)
	}

	return nil
}
