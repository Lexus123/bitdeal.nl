package handlers

import (
	"fmt"
	"net/http"
)

/*
GetHomePage ...
*/
func GetHomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hoi hoi!")
}

// /*
// GetSignInPage ...
// */
// func GetSignInPage(w http.ResponseWriter, r *http.Request) {
// 	// Check whether the client already has a cookie
// 	if checkCookie(r) {
// 		// You have a cookie, then go to homepage
// 		http.Redirect(w, r, "/", http.StatusSeeOther)
// 	} else {
// 		// Functions for the template
// 		var funcMap = template.FuncMap{
// 			"trimCaptcha":     trimCaptcha,
// 			"printErrorParam": printErrorParam,
// 		}

// 		// Setting up the template
// 		templates := addTemplate("templates/pages/auth/signin.gohtml", false)
// 		tmpl := template.Must(template.New("signin.gohtml").Funcs(funcMap).ParseFiles(templates...))

// 		// Configuration struct for digits captcha
// 		var config = base64Captcha.ConfigDigit{
// 			Height:     100,
// 			Width:      300,
// 			MaxSkew:    0.7,
// 			DotCount:   80,
// 			CaptchaLen: 5,
// 		}

// 		// Creation of the captcha base64 image and the answer (idKey)
// 		idKey, cap := base64Captcha.GenerateCaptcha("", config)
// 		base64string := base64Captcha.CaptchaWriteToBase64Encoding(cap)

// 		// Values that need to be stored in the cookie
// 		cookieValues := map[string]string{
// 			"captchaID": idKey,
// 		}

// 		// Creation of the actual cookie for signing in/up
// 		if encodedValues, err := secCookie.Encode("websitecaptcha", cookieValues); err == nil {
// 			cookie := &http.Cookie{
// 				Name:     "websitecaptcha",
// 				Value:    encodedValues,
// 				Path:     "/",
// 				HttpOnly: true,
// 				Expires:  time.Now().Add(5 * time.Minute),
// 			}
// 			http.SetCookie(w, cookie)
// 		}

// 		// If the URL has parameters, put them inside this var. Example "/signin?error=username,count"
// 		var params []string

// 		// Error parameter present, then do something
// 		if errorParam := r.URL.Query()["error"]; len(errorParam) > 0 {
// 			params = strings.Split(r.URL.Query()["error"][0], ",")
// 		}

// 		// Data for the template
// 		templateData := models.IndexData{
// 			Title:   "Sign in",
// 			Captcha: base64string,
// 			Params:  params,
// 		}

// 		// Execute the template
// 		tmpl.ExecuteTemplate(w, "layout", templateData)
// 	}
// }

// /*
// GetHomePage ...
// */
// func GetHomePage(w http.ResponseWriter, r *http.Request) {
// 	if checkCookie(r) {
// 		templates := addTemplate("templates/pages/homepage.gohtml", true)
// 		tmpl := template.Must(template.ParseFiles(templates...))

// 		data := models.IndexData{
// 			Title: "Homepage",
// 		}

// 		tmpl.ExecuteTemplate(w, "layout", data)
// 	} else {
// 		http.Redirect(w, r, "/signin", http.StatusSeeOther)
// 	}
// }

// /*
// GetTest ...
// */
// func GetTest(w http.ResponseWriter, r *http.Request) {
// 	if checkCookie(r) {
// 		fmt.Fprintf(w, "Salty: %v\n", security.CreateSalt())
// 		fmt.Fprintf(w, "'Salt + Hoi' SHA-3 512: %v\n", security.HashPassword(security.CreateSalt(), "hoi"))
// 	} else {
// 		http.Redirect(w, r, "/signin", http.StatusSeeOther)
// 	}
// }
