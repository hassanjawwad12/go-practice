package auth

// we made it abstract by using lowercase
func extractSession() string {
	return "logged in"
}

func Session() string {
	return extractSession()
}
