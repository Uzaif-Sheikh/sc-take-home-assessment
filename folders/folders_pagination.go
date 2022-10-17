package folders

// Copy over the `GetFolders` and `FetchAllFoldersByOrgID` to get started
import (
	// "fmt"
	//"github.com/gofrs/uuid"
)

func GetFolders(req *FetchPagRequest) (*FetchPagResponse, error) {
	r, _ := FetchAllFoldersByOrgID(req.OrgID)
	tok := req.Token

	res := []*Folder{}
	for i := (tok * 10); i <= (tok + 9) ; i++ {
		res = append(res, r[i])	
	}
	var ffr *FetchPagResponse
	ffr = &FetchPagResponse{Folders: res, Token: (tok+1)}
	return ffr, nil
}
