package middleware

import (
	"net/http"

	"github.com/lanxre/kyokusulib/internal/utils/response"
)

type RoleName string

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

func hasSufficientRole(user, required RoleName) bool {
	uWeight, uOk := roleWeights[user]
	rWeight, rOk := roleWeights[required]
	return uOk && rOk && uWeight >= rWeight
}

func RoleGuard(next http.HandlerFunc, required RoleName) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctxValue, ok := r.Context().Value(UserRoleKey).(string)
		if !ok {
			response.Error(w, http.StatusForbidden, "Forbidden: role not found")
			return
		}

		userRole := RoleName(ctxValue)
		if !hasSufficientRole(userRole, required) {
			response.Error(w, http.StatusForbidden, "Forbidden: insufficient permissions")
			return
		}

		next(w, r)
	}
}