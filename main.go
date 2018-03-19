package main

import (
	"bufio"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	t, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	payload := strings.Split(t, ".")[1]

	data, err := base64.RawStdEncoding.DecodeString(payload)
	if err != nil {
		log.Fatal(err)
	}

	var claims jwt.StandardClaims
	err = json.Unmarshal(data, &claims)
	if err != nil {
		log.Fatal(err)
	}

	expiration := time.Unix(claims.ExpiresAt, 0)
	if expiration.Before(time.Now()) {
		fmt.Println("Token is expired!")
	} else {
		fmt.Printf("Expires in %s (%s)\n", expiration.Sub(time.Now()), expiration)
	}
}
