package handlers

// import (
// 	"net/http"
// 	"strings"
// 	"time"
// )

// /*
// PostSignIn ...
// */
// func PostSignIn(w http.ResponseWriter, r *http.Request) {
// 	// Check whether the cookie is valid, and if so, get the captcha solution from it
// 	if captchaSolution, cookieValid := checkSignInCookie(r); cookieValid {
// 		r.ParseForm()

// 		var paramsSlice []string

// 		errors := make(map[string]bool, 6)

// 		errors["usernameLength"], errors["usernameChars"] = verifyUsername(r.Form.Get("username"))
// 		errors["captcha"] = verifyCaptcha(captchaSolution, r.Form.Get("captcha"))
// 		errors["passwordLength"], errors["passwordNumber"], errors["passwordUpper"] = verifyPassword(r.Form.Get("password"))

// 		for key, value := range errors {
// 			if !value {
// 				paramsSlice = append(paramsSlice, key)
// 			}
// 		}

// 		if len(paramsSlice) > 0 {
// 			paramsString := strings.Join(paramsSlice, ",")

// 			// Some errors in the form
// 			http.Redirect(w, r, "/signin?error="+paramsString, http.StatusSeeOther)
// 		} else {
// 			// Get user from DB
// 			cookieValues := map[string]string{
// 				"user": "test",
// 			}

// 			if encodedValues, err := secCookie.Encode("website", cookieValues); err == nil {
// 				cookie := &http.Cookie{
// 					Name:     "website",
// 					Value:    encodedValues,
// 					Path:     "/",
// 					HttpOnly: true,
// 					Expires:  time.Now().Add(1 * time.Minute),
// 				}
// 				http.SetCookie(w, cookie)
// 			}

// 			if encodedValues, err := secCookie.Encode("websitecaptcha", cookieValues); err == nil {
// 				cookie := &http.Cookie{
// 					Name:     "websitecaptcha",
// 					Value:    encodedValues,
// 					Path:     "/",
// 					HttpOnly: true,
// 					Expires:  time.Now().Add(-5 * time.Minute),
// 				}
// 				http.SetCookie(w, cookie)
// 			}

// 			// Go to '/'
// 			http.Redirect(w, r, "/", http.StatusSeeOther)

// 			// Everything is going according to plan. Now sign in or up!
// 			// if action := r.Form.Get("action"); action == "signin" {
// 			// 	postSignIn(w, r)
// 			// } else {
// 			// 	postSignUp(w, r)
// 			// }
// 		}
// 	} else {
// 		// Cookie expired
// 		http.Redirect(w, r, "/signin?error=expired", http.StatusSeeOther)
// 	}
// }

// func postSignIn(w http.ResponseWriter, r *http.Request) {

// }

// func postSignUp(w http.ResponseWriter, r *http.Request) {

// }
