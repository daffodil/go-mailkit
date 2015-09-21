

package mailbox

import(

	"fmt"

	"bytes"

	"net/http"
	"net/mail"
	"encoding/json"

	"github.com/gorilla/mux"
	"code.google.com/p/go-imap/go1/imap"
	"github.com/jhillyerd/go.enmime"
)

//===================================================================
type Message struct {
	Uid uint32 `json:"uid"`
	Folder string `json:"folder"`

	FromName string `json:"from_name"`
	FromEmail string `json:"from_email"`

	Date string `json:"date"`

	Seen bool `json:"seen"`
	Flagged bool `json:"flagged"`
	Answered bool `json:"answered"`

	Subject string `json:"subject"`
	BodyHtml string  `json:"body_html"`
	BodyText string  `json:"body_text"`

	ContentType string `json:"content_type"`
}


type MessagePayload struct {

	Success bool `json:"success"`
	Message *Message `json:"message"`

	Error string `json:"error"`
}

func GetMessage(folder, uid string, client *imap.Client) (messag *Message, e error ){

	cmd, err := client.Select(folder, true)
	if err != nil {
		return nil, err
	}

	uidlist, _ := imap.NewSeqSet(uid)
	//uidlist.Add(uid)

	fmt.Println("get_mess", folder, uid)
	mess := new(Message)
	mess.Folder = folder

	cmd, err = imap.Wait( client.UIDFetch(uidlist, "FLAGS", "INTERNALDATE", "RFC822.SIZE", "RFC822")) //  "RFC822.HEADER",  "BODY.PEEK[TEXT]") )
	if err != nil {
		return mess, err
	}
	fmt.Println( len(cmd.Data), cmd.Data)
	rsp := cmd.Data[0]
	minfo := rsp.MessageInfo()

	mess.Uid = minfo.UID

	msg, _ := mail.ReadMessage(bytes.NewReader(imap.AsBytes(minfo.Attrs["RFC822"])))
	mime, mime_err := enmime.ParseMIMEBody(msg)



	for flag, boo := range  minfo.Flags {
		if flag == "\\Seen" && boo {
			mess.Seen = true
		}
		if flag == "\\Flagged" && boo {
			mess.Flagged = true
		}
	}
	/*
	bites := imap.AsBytes(minfo.Attrs["RFC822"])
	msg, msg_err := mail.ReadMessage(bytes.NewReader(bites))
	if msg_err != nil {
		return mess, msg_err
	}
	*/

	//fmt.Println(msg.Header.Get("Content-type"))
	//fmt.Println(msg.Header.Get("To"))
	//fmt.Println(msg.Header.Get("Delivered-To"))

	// From
	from, fro_err := mail.ParseAddress(msg.Header.Get("From"))
	if fro_err != nil {
		fmt.Println("address ettot")
	} else {
		mess.FromName = from.Name
		mess.FromEmail = from.Address
	}

	//for i, m := range minfo.Attrs {

	//fmt.Println(i,m)
	//}
	// Date
	dat := imap.AsDateTime(minfo.Attrs["INTERNALDATE"])
	mess.Date = dat.Format("2006-01-02 15:04:05")

	mess.Subject = msg.Header.Get("Subject")
	mess.ContentType = msg.Header.Get("Content-Type")
	//fmt.Println("body=", cmd.Data[0].String)

	//bodyb := imap.AsBytes(minfo.Attrs["BODY[TEXT]"])
	//bb := 	bytes.NewReader(bytes.NewReader(header + bodyb))

	//fmt.Println("bodyb=", string(msg.Body))




	//fmt.Printf("----\n%v\n", mime.Html == nil)
	//mess.Body = mime.Text
	if mime_err != nil {
		fmt.Println("err=", mime_err, mime)
	}
	//*/
	//fmt.Println("body=", body)


	mess.BodyText =  mime.Text
	mess.BodyHtml =  mime.Html //imap.AsString(minfo.Attrs["RFC822"])

	return mess, nil

}



func MessageHandler(resp http.ResponseWriter, req *http.Request) {


	//= Setup AJAX Payload
	SetAjaxHeaders(resp)
	client := GetImapClient(resp, req)
	if client == nil {
		return
	}
	defer func() { client.Logout(0) }()

	payload := new(MessagePayload)
	payload.Success = true
	payload.Message = new(Message)

	var err error
	vars := mux.Vars(req)
	payload.Message, err = GetMessage(vars["folder"], vars["uid"], client)
	if err != nil {
		fmt.Println("err", err)

	}
	//payload.Message.Folder = vars["folder"]
	//payload.Message.Uid = vars["uid"]

	json_str, _ := json.MarshalIndent(payload, "" , "  ")
	fmt.Fprint(resp, string(json_str))

}



