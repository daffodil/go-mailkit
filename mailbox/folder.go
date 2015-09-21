
package mailbox

import(

	"fmt"
	"net/http"
	"encoding/json"

	"code.google.com/p/go-imap/go1/imap"
)

//===================================================================

type Folder struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type FoldersPayload struct {
	Success bool `json:"success"`
	Folders[] *Folder `json:"folders"`
	Error string `json:"error"`
}

//= Return list of IMAP folders
func GetFolders(client *imap.Client )([]*Folder, error){

	folders := make([]*Folder, 0)

	cmd, err := imap.Wait( client.List("", "*") )
	if err != nil {
		return folders, err
	}

	for idx := range cmd.Data {
		info := cmd.Data[idx].MailboxInfo()
		fol := new(Folder)
		fol.Name = info.Name
		for flag, boo := range  info.Attrs {
			fmt.Println( info.Name, boo, flag)
			if info.Name == "INBOX"  && boo {
				fol.Type = "inbox"

			} else if flag == "\\Junk" && boo {
				fol.Type = "junk"

			} else if flag == "\\Trash" && boo {
				fol.Type = "trash"

			} else if flag == "\\Sent" && boo {
				fol.Type = "sent"

			} else if flag == "\\Drafts" && boo {
				fol.Type = "drafts"
			}

		}
		folders = append(folders, fol)
	}

	return folders, nil
}


// Handle /mailbox/<email>/folders
func FoldersHandler(resp http.ResponseWriter, req *http.Request) {

	SetAjaxHeaders(resp)

	client := GetImapClient(resp, req)
	if client == nil {
		return
	}
	defer func() { client.Logout(0) }()

	payload := new(FoldersPayload)
	payload.Success = true

	var err error
	payload.Folders, err = GetFolders(client)
	if err != nil {
		payload.Error = err.Error()
	}

	json_str, _ := json.MarshalIndent(payload, "" , "  ")
	fmt.Fprint(resp, string(json_str))
}
