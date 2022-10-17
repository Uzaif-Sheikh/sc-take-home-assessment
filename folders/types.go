package folders

import "github.com/gofrs/uuid"

type FetchFolderRequest struct {
	OrgID uuid.UUID
}

type FetchPagRequest struct {
	OrgID uuid.UUID
	Token int
}

type FetchPagResponse struct {
	Folders []*Folder
	Token int
}

type FetchFolderResponse struct {
	Folders []*Folder
}