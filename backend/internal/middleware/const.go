package middleware

type contextKey string

type RoleName string


const (
	UserIDKey    contextKey = "userID"
	UserEmailKey contextKey = "userEmail"
)

const (
	UserRoleKey   contextKey = "userRole"
	RoleUser      RoleName   = "user"
	RoleModerator RoleName   = "moderator"
	RoleAdmin     RoleName   = "admin"
)

var roleWeights = map[RoleName]int{
	RoleUser:      1,
	RoleModerator: 50,
	RoleAdmin:     100,
}