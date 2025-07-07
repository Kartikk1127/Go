package auth

func extractSession() string {
	return "logged in"
}

func GetSession() string {
	return extractSession()
}
