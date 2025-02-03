package user

import (
	"encoding/json"
	"net/http"
	"pockethealth/internchallenge/pkg/router"
	"regexp"
	"strings"
)

// An UserApiController binds http requests to an api service and writes the service results to the http response
type UserApiController struct {
	service UserApiService
}

// NewUserApiController creates a new api controller
func NewUserApiController(s UserApiService) router.Router {
	return UserApiController{service: s}
}

// Routes returns all of the api route for the UserApiController
func (c UserApiController) Routes() router.Routes {
	return router.Routes{
		{
			Name:        "PostRegister",
			Method:      strings.ToUpper("Post"),
			Pattern:     "/register",
			HandlerFunc: c.PostRegister,
		},
	}
}

type PostRegisterBody struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	FavColor string `json:"favColor"`
}

type PostRegisterResponse struct {
	UserId string `json:"user_id"`
}

func ValidateName(name string) bool {
	return len(name) > 0
}

func ValidateEmail(email string) bool {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return emailRegex.MatchString(email)
}

func ValidateHexColor(hexCode string) bool {
	hexColorRegex := regexp.MustCompile(`^#(?:[0-9a-fA-F]{3}){1,2}$`)
	return hexColorRegex.MatchString(hexCode)
}

// PostRegister - Register a New User
func (c UserApiController) PostRegister(w http.ResponseWriter, r *http.Request) {
	// read request body
	data := &PostRegisterBody{}

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Validate inputs
	if !ValidateName(data.Name) {
		http.Error(w, "Please enter a name.", http.StatusBadRequest)
		return
	}
	if !ValidateEmail(data.Email) {
		http.Error(w, "Invalid email form.", http.StatusBadRequest)
		return
	}
	if !ValidateHexColor(data.FavColor) {
		http.Error(w, "Invalid hex color code form. Pleas use form of #000000", http.StatusBadRequest)
		return
	}

	// call service
	userId, err := c.service.PostRegister(r.Context(), data.Name, data.Email, data.FavColor)
	if err != nil {
		panic(err)
	}

	// create and send response
	resp := PostRegisterResponse{
		UserId: userId,
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		panic(err)
	}
}
