package jwt

import (
	"backend/pkg/rpc/gateway"
	"context"
	"net/http"
	"time"
)

func (c *JWT) SetTokenToCookie(ctx context.Context, token string) {
	gateway.SetCookie(ctx, &http.Cookie{
		Name:    c.CookieKey,
		Value:   token,
		Domain:  c.CookieDomain,
		Expires: time.Now().Add(time.Duration(c.ExpireTime)),
		Secure:  true,
	})
}
