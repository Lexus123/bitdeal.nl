package handlers

import (
	"net/http"
	"strings"
	"unicode"

	"github.com/mojocn/base64Captcha"
)

func addTemplate(template string) []string {
	var templates []string
	templates = append([]string{template}, baseTemplates...)
	return templates
}

func addExtraTemplate(alreadyAdded []string, template string) []string {
	var templates []string
	templates = append([]string{template}, alreadyAdded...)
	return templates
}

func checkCookie(r *http.Request) bool {
	if cookie, err := r.Cookie("website"); err == nil {
		value := make(map[string]string)
		if err = secCookie.Decode("website", cookie.Value, &value); err == nil {
			// Make call to DB to check if value["user"] exists
			if value["user"] == "test" {
				return true
			}
		}
	}
	return false
}

func checkSignInCookie(r *http.Request) (string, bool) {
	if cookie, err := r.Cookie("websitecaptcha"); err == nil {
		value := make(map[string]string)
		if err = secCookie.Decode("websitecaptcha", cookie.Value, &value); err == nil {
			if value["captchaID"] != "" {
				return value["captchaID"], true
			}
		}
	}
	return "Some error occured (your captcha probably expired)", false
}

func verifyUsername(username string) (threeOrMore, isBase62 bool) {
	if usernameLength := len(username); usernameLength >= 3 {
		for _, char := range username {
			isBase62 = strings.Contains(base62Chars, string(char))
		}
		return true, isBase62
	}
	return false, true
}

func verifyPassword(s string) (eightOrMore, number, upper bool) {
	charCount := 0
	for _, c := range s {
		switch {
		case unicode.IsNumber(c):
			number = true
			charCount++
		case unicode.IsUpper(c):
			upper = true
			charCount++
		case unicode.IsLetter(c) || unicode.IsSpace(c):
			charCount++
		case unicode.IsSymbol(c) || unicode.IsPunct(c):
			charCount++
		default:
			return false, false, false
		}
	}
	eightOrMore = charCount >= 8
	return
}

func verifyCaptcha(idkey, verifyValue string) bool {
	return base64Captcha.VerifyCaptcha(idkey, verifyValue)
}
