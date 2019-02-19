package login

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"keep-server/model/user"
	"keep-server/session"
	"keep-server/utill/hash"

	"github.com/labstack/echo"
)

var (
	db = user.Connect()
)

// Login login
func Login(c echo.Context) (err error) {

	var json struct {
		Name     string `json:"name"`
		Password string `json:"pass"`
	}

	if err = c.Bind(&json); err != nil {
		fmt.Println("json error")
		return
	}

	if json.Name == "" || json.Password == "" {
		fmt.Println("null")
		return c.String(http.StatusBadRequest, "null")
	}

	var u user.User
	db.Find(&u, "Name = ?", json.Name)

	if u.Name == "" {
		fmt.Println("unexist")
		return c.String(http.StatusBadRequest, "unexist")
	}

	if hash.Sha1(json.Password) != u.Password {
		return c.String(http.StatusBadRequest, "wrong pass")
	}

	sessionKey := genSession(c, u.Code)

	var result struct {
		Name    string `json:"name"`
		Code    string `json:"code"`
		Session string `json:"session"`
	}
	result.Name = u.Name
	result.Code = u.Code
	result.Session = sessionKey

	fmt.Println("login ok:" + u.Name)
	return c.JSON(http.StatusOK, result)
}

func genSession(c echo.Context, code string) string {
	rand.Seed(time.Now().UnixNano())
	sessionKey := hash.Sha1(strconv.Itoa(rand.Int()))
	session.Set(code, sessionKey)

	return sessionKey
}

// IsLogin .
func IsLogin(c echo.Context, code string, key string) bool {
	sessionKey := session.Get(code)
	fmt.Println("islogin:", sessionKey, key)

	if sessionKey == key {
		return true
	}
	return false
}
