package handlers

func trimCaptcha(s string) string {
	m := 0
	for i := range s {
		if m >= 22 {
			return s[i:]
		}
		m++
	}
	return s[:0]
}

func printErrorParam(s string) string {
	switch s {
	case "passwordNumber":
		return "Password didn't include a number"
	case "passwordUpper":
		return "Password didn't include an uppercase letter"
	case "passwordCount":
		return "Password wasn't 8 characters or longer"
	case "captcha":
		return "Wrong captcha"
	case "usernameChars":
		return "The username included an illegal character"
	case "usernameLength":
		return "The username wasn't 3 characters or longer"
	case "expired":
		return "The captcha expired"
	}

	return "Error occured, try again"
}
