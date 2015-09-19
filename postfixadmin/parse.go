

package postfixadmin

import(

	"fmt"
	//"net/http"
	//"encoding/json"
	//"errors"
	"errors"
	"strings"

)

type EmailParts struct {
	Email string
	User string
	Domain string
}


func ParseEmail(raw_email string) (*EmailParts, error) {

	stripped := strings.TrimSpace(raw_email)
	if len(stripped) == 0 {
		return nil, errors.New("Invalid Email - zero length")
	}

	if strings.Contains(stripped, "@") == false {
		return nil, errors.New("Invalid Email, no @ in `" + raw_email + "` ")
	}
	user_domain :=  strings.Split(stripped, "@")
	fmt.Println(user_domain)

	if DomainExists(user_domain[1]) == false {
		return nil, errors.New("Domain not exist in Db for email `" + raw_email + "` ")
	}

	em := new(EmailParts)

	return em, nil

}

