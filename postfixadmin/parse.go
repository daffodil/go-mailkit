

package postfixadmin

import(
	"errors"
	"strings"
)

type Addr struct {
	Address string
	User string
	Domain string
}


func ParseAddress(email_address string) (*Addr, error) {

	stripped := strings.TrimSpace(email_address)
	if len(stripped) == 0 {
		return nil, errors.New("Invalid Email - zero length")
	}

	if strings.Contains(stripped, "@") == false {
		return nil, errors.New("Invalid Email, no @ in `" + email_address + "` ")
	}

	user_domain :=  strings.Split(stripped, "@")
	if DomainExists(user_domain[1]) == false {
		return nil, errors.New("Domain not exist in Db for email `" + email_address + "` ")
	}

	em := new(Addr)
	em.Address = stripped
	em.User = user_domain[0]
	em.Domain = user_domain[1]

	return em, nil

}

