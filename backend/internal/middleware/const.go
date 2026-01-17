package middleware

type contextKey string

type RoleName string


const (
	UserIDKey    contextKey = "userID"
	UserEmailKey contextKey = "userEmail"
	UserRoleKey   contextKey = "userRole"
)

const (
	RoleUser      RoleName   = "user"
	RoleModerator RoleName   = "moderator"
	RoleAdmin     RoleName   = "admin"
	RolePublisher RoleName   = "publisher"
)

var roleWeights = map[RoleName]int{
	RoleUser:      1,
	RolePublisher: 30,
	RoleModerator: 50,
	RoleAdmin:     100,
}