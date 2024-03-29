package handlers

import (
	"github.com/gorilla/securecookie"
)

var (
	/*
		base62Chars ...
	*/
	base62Chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	/*
		baseTemplates ...
	*/
	baseTemplates = []string{
		"templates/defaultComponents/head.html",
		"templates/defaultComponents/footer.html",
		"templates/layouts/base.html",
	}

	/*
		hashKey is used for verification
	*/
	hashKey = securecookie.GenerateRandomKey(32)

	/*
		blockKey is used for encryption (AES-256)
	*/
	blockKey = securecookie.GenerateRandomKey(32)

	/*
		secCookie is an encrypted cookie
	*/
	secCookie = securecookie.New(hashKey, blockKey)
)
