package middleware

type AuthenticatedUser struct {
	UserID      string
	Permissions map[string]bool
}

func (ru *AuthenticatedUser) IsEmpty() bool {
	return ru.UserID == "" && len(ru.Permissions) == 0
}
