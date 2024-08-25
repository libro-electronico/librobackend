package at

import (
	"errors"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

func URLParam(reqpath string, url string) bool {
	urls := strings.Split(url, ":")
	prefix := reqpath[:strings.LastIndex(reqpath, "/")+1]
	return prefix == urls[0]
}

func GetParam(r *http.Request) string {
	return r.URL.Path[strings.LastIndex(r.URL.Path, "/")+1:]
}

func GetAddress() (ipport string, network string) {
	port := os.Getenv("PORT")
	network = "tcp4"
	if port == "" {
		port = ":8080"
	} else if port[0:1] != ":" {
		ip := os.Getenv("IP")
		if ip == "" {
			ipport = ":" + port
		} else {
			if strings.Contains(ip, ".") {
				ipport = ip + ":" + port
			} else {
				ipport = "[" + ip + "]" + ":" + port
				network = "tcp6"
			}
		}
	}
	return
}

func GetClientIP(r *http.Request) (string, error) {
	ips := r.Header.Get("X-Forwarded-For")
	splitIps := strings.Split(ips, ",")

	if len(splitIps) > 0 {
		// get last IP in list since ELB prepends other user defined IPs, meaning the last one is the actual client IP.
		netIP := net.ParseIP(splitIps[len(splitIps)-1])
		if netIP != nil {
			return netIP.String(), nil
		}
	}

	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return "", err
	}

	netIP := net.ParseIP(ip)
	if netIP != nil {
		ip := netIP.String()
		if ip == "::1" {
			return "127.0.0.1", nil
		}
		return ip, nil
	}

	return "", errors.New("IP not found")
}

func GetIPaddress() string {

	resp, err := http.Get("https://icanhazip.com/")

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}
	return string(body)
}

// HashPassword hashes the password using bcrypt.
func HashPassword(password string) (string, error) {
	// Generate a bcrypt hash of the password
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

// CheckPasswordHash compares the hashed password with the provided password.
func CheckPasswordHash(password, hashedPassword string) bool {
	// Compare the password with the hash
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		log.Println("Password comparison failed:", err)
		return false
	}
	return true
}

// ValidatePassword compares a plain text password with a hashed password.
func ValidatePassword(plainPassword, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
	return err == nil
}