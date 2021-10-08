package grpcproject

type JwtSession struct {
	Session []string
}

func (receiver JwtSession) AddUser(token string) {
	receiver.Session = append(receiver.Session, token)
}
func (receiver JwtSession) RemoveUser(token string) bool {
	var indexnum int = -1
	for index, s := range receiver.Session {
		if s == token {
			indexnum = index
		}
	}
	if indexnum != -1 {

		receiver.Session = append(receiver.Session[:indexnum], receiver.Session[indexnum+1:]...)
		return true
	}
	return false
}

func (receiver JwtSession) ValidateUser(token string) bool {
	for _, s := range receiver.Session {
		if s == token {
			return true
		}
	}
	return false
}

var Localsession JwtSession
