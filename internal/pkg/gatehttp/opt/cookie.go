package opt

import (
	"fmt"
	"net/http"
)

type CookieString string

var (
	CookieNameJWT = "JWT"
)

func JWTCookieString(token string) CookieString {
	return CookieString(fmt.Sprintf("Name: %s Value: %s", CookieNameJWT, token))
}

func (cs CookieString) Take_httpCookie_FromCookieString() (*http.Cookie, error) {
	var r http.Cookie
	if _, err := fmt.Sscanf(string(cs), "Name: %s Value: %s", &r.Name, &r.Value); err != nil {
		return nil, err
	}
	return &r, nil
}
