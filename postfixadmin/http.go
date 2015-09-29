
package postfixadmin

import(

	//"fmt"
	"net/http"
	"encoding/json"


)

//= Ajax struct for `domain` all
type ErrPayload struct {
	Success bool `json:"success"` // keep extjs happy
	Error string `json:"error"`
}

func CreatePermissionErrPayload() string {

	payload := ErrPayload{}
	payload.Success = true
	payload.Error = "Permission denied"
	json_str, _ := json.MarshalIndent(payload, "" , "  ")
	return string(json_str)
}

func AjaxAuth(resp http.ResponseWriter, req *http.Request) bool {

	// Set Ajax Headers
	resp.Header().Set("Content-Type", "application/json")



	switch req.Method {

	case "POST":
		req.ParseForm()
		if req.Form.Get("auth") != conf.AuthSecret {
			http.Error(resp, CreatePermissionErrPayload(), 500)
			return false
		}

	default:

		if req.URL.Query().Get("auth") != conf.AuthSecret {
			http.Error(resp, CreatePermissionErrPayload(), 500)
			return false
		}

	}

	return true
}
