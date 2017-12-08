package client

//Store all API URI address.
const (
	ServiceURL = "https://localhost:9096/v1"

	WorkspaceURL       = ServiceURL + "/batches"
	QueryWorkspaceURL  = WorkspaceURL
	CreateWorkspaceURL = WorkspaceURL + "/archive/tar"
	UploadImageURL     = WorkspaceURL
)

//GetUploadImageURL - get server upload image url
func GetUploadImageURL(queryID string) string {
	return UploadImageURL + "/" + queryID + "/upload"
}

//GetWorkspaceURL -
func GetWorkspaceURL(queryID string) string {
	return UploadImageURL + "/" + queryID
}
