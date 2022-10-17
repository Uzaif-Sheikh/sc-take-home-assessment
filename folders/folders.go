package folders

import (
	// "fmt"
	"github.com/gofrs/uuid"
)

func GetAllFolders(req *FetchFolderRequest) (*FetchFolderResponse, error) {

	// The GetAllFolders function given a FetchFolderRequest which is a struct containing the orgID
	// returns a struct FetchFolderResponse containing the array of pointers to the folder struct.
	// The function GetAllFolders make use of the function FetchAllFolderByOrgID to get all the 
	// folder which have the orgID equal to the req's OrgID.

	// used variable
	// var (
	// 	err error
	// 	f1  Folder
	// 	fs  []*Folder
	// )

	// In the code below, the function FetchALlFolderByOrgID is being called which returns 
	// the array of pointer to the folder `r` and then the code loop through the returned array `r` 
	// stores the value of the pointers in the array of the struct folder `f` then loop through the array `f`,
	// and stores the reference to each struct value in a new array of pointer to struct folder named `fp`, Lastly,
	// make a FetchFolderResponse struct and store the `fp` array in it and returns the response.

	// The code below can be improved by just using the array `r` given by the FetchAllFoldersByOrgID function:
	//	
	//			r, _ := FetchAllFoldersByOrgID(req.OrgID)
	//			var ffr *FetchFolderResponse
	//			ffr = &FetchFolderResponse{Folders: r}
	//
	//			return ffr, nil

	/*

		--|-------------------------------------------------------------------|--
		  |	 I had to change the implemention to pass the unit test,      |
		  |	 the problem was that when copying it stores the reference&   |
		  |	 which only copy the same pointer in the array.		      |
		--|-------------------------------------------------------------------|--


	f := []Folder{}
	fmt.Println(*r[1])
	for _, v := range r {
		f = append(f, *v)
	}

	fmt.Println(f)
	var fp []*Folder
	for _, v1 := range f {
		fp = append(fp, &v1)
	}
	*/

	r, _ := FetchAllFoldersByOrgID(req.OrgID)
	var ffr *FetchFolderResponse
	ffr = &FetchFolderResponse{Folders: r}
	return ffr, nil
}

func FetchAllFoldersByOrgID(orgID uuid.UUID) ([]*Folder, error) {
	folders := GetSampleData()

	resFolder := []*Folder{}
	for _, folder := range folders {
		if folder.OrgId == orgID {
			// fmt.Println(folder.Name)
			resFolder = append(resFolder, folder)
		}
	}
	return resFolder, nil
}
