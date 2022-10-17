package folders

import (
	// "fmt"
	"testing"
	"github.com/gofrs/uuid"
	// "github.com/georgechieng-sc/interns-2022/folders"
	// "github.com/stretchr/testify/assert"
)

func Test_GetAllFolders(t *testing.T) {
	t.Run("test", func(t *testing.T) {
		// your test/s here
		u, _ := uuid.FromString("c1556e17-b7c0-45a3-a6ae-9546248fb17a")
		ffr := FetchFolderRequest{OrgID: u}
		req := &ffr
		getFolder, _ := GetAllFolders(req)

		if len(getFolder.Folders) != 999 {
			t.Errorf("error-1")
		}

		u1, _ := uuid.FromString("72dc6a27-14d5-4643-a882-087990903069")
		ff := Folder{u1,"balanced-bella",u,true}

		if *getFolder.Folders[1] != ff {
			t.Errorf("error-2")
		}
	})
}
