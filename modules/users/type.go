package users

// Implement your helping types here
type (
	// UserCreationPayload maps input to create user
	UserCreationPayload struct {
		Username string
		Email    string
		Password string
	}
)
