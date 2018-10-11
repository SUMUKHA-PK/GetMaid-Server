package handlers

import (
	"GetMaid/handlers/methods"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"io"
	"net/http"
)

type MaidSignUp struct {
}

type HireSignUp struct {
	Email   string  `json:"email"`
	Phone   string  `json:"phone"`
	Name    string  `json:"name"`
	Address Address `json:"address"`
}

type Address struct {
	HouseNo  string `json:"house_no"`
	Locality string `json:"locality"`
	City     string `json:"city"`
	PinCode  string `json:"pin_code"`
}

func SignupGetHandler(res http.ResponseWriter) {
	io.WriteString(res, "Signup Route")

}

func SignupPostHandler(req *http.Request, res http.ResponseWriter) {

	var e error

	defer methods.ErrorHandler(res, &e)

	req.ParseForm()

	if len(req.Form["email"]) == 0 || len(req.Form["password"]) == 0 {
		panic("EMAIL OR PASSWORD MISSING")
	}

	hpw, e := bcrypt.GenerateFromPassword([]byte(req.Form["password"][0]), bcrypt.MinCost)

	if e != nil {
		panic(500)
	}

	fmt.Println(hpw)


}

func SignupHandler(res http.ResponseWriter, req *http.Request) error {
	var e error

	defer methods.ErrorHandler(res, &e)

	switch {
	case methods.CheckCase("GET", "/signup", req):
		SignupGetHandler(res)

	case methods.CheckCase("POST", "/signup", req):

		SignupPostHandler(req, res)
	default:
		panic(404)
	}



	return e
}
