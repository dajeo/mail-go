package server

import (
	"github.com/gin-contrib/cors"
	"mail/config"
	"time"
)

func Init() {
	r := NewRouter()

	// See this before deploy to production
	// https://pkg.go.dev/github.com/gin-gonic/gin#readme-don-t-trust-all-proxies
	proxyErr := r.SetTrustedProxies(nil)
	if proxyErr != nil {
		return
	}

	corsConfig := cors.Config{
		AllowAllOrigins:  true,
		AllowCredentials: true,
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"Authorization"},
		MaxAge:           12 * time.Hour,
	}

	r.Use(cors.New(corsConfig)) // Configure before production

	err := r.Run(config.GetConfig().GetString("server.addr"))
	if err != nil {
		return
	}
}
