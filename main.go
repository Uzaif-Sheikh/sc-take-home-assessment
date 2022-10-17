package main

import (
	"fmt"

	"github.com/georgechieng-sc/interns-2022/folders"
	"github.com/gofrs/uuid"
)

func main() {
	req := &folders.FetchFolderRequest{
		OrgID: uuid.FromStringOrNil(folders.DefaultOrgID),
	}

	res, err := folders.GetAllFolders(req)
	if err != nil {
		fmt.Printf("%v", err)
		return
	}

	// req1 := &folders.FetchPagRequest{
	// 	OrgID: uuid.FromStringOrNil(folders.DefaultOrgID), Token: 0,
	// }
	// r, _ := folders.GetFolders(req1)
	// fmt.Println(r)

	folders.PrettyPrint(res)
}
