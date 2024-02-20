package service

import (
	"net"
	"net/http"
	"strings"
)

func ObtenerIP(r *http.Request) string {
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return ""
	}

	// Comprobar si la IP es una dirección IPv6
	if strings.Count(ip, ":") >= 2 {
		if ip[0] == '[' {
			ip = strings.Trim(ip, "[]")
		}
	}

	return ip
}
