package constants

type UserStatus string

const (
	Online 	UserStatus = "online"
	Offline UserStatus = "offline"
	Idle  	UserStatus = "idle"
	Ban   	UserStatus = "ban"
)


func (us UserStatus) String() string {
	return string(us)
}

func (us UserStatus) IsValid() bool {
	switch us {
	case Online, Offline, Idle, Ban:
		return true
	}
	return false
}