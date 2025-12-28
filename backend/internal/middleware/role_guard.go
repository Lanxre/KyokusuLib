package middleware

import (
	"net/http"

	"github.com/lanxre/kyokusulib/internal/utils/response"
)

const (
	UserRoleKey  contextKey = "userRole" 
)

var roleWeights = map[string]int{
	"user":      1,
	"moderator": 50,
	"admin":     100,
}


func RoleGuard(next http.HandlerFunc, requiredRole string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userRole, ok := r.Context().Value(UserRoleKey).(string)
		if !ok {
			response.Error(w, http.StatusForbidden, "Forbidden: role not found")
			return
		}

		userWeight, userExists := roleWeights[userRole]
		requiredWeight, reqExists := roleWeights[requiredRole]

		if !userExists || !reqExists {
			response.Error(w, http.StatusForbidden, "Forbidden: unknown role")
			return
		}

		if userWeight < requiredWeight {
			response.Error(w, http.StatusForbidden, "Forbidden: insufficient permissions")
			return
		}

		next(w, r)
	}
}

