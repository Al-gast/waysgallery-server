package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	authdto "waysgallery/dto/auth"
	dto "waysgallery/dto/result"
	userdto "waysgallery/dto/user"
	"waysgallery/models"
	"waysgallery/pkg/bcrypt"
	jwtToken "waysgallery/pkg/jwt"
	"waysgallery/repositories"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
)

type handlerAuth struct {
	AuthRepository repositories.AuthRepository
}

func HandlerAuth(AuthRepository repositories.AuthRepository) *handlerAuth {
	return &handlerAuth{AuthRepository}
}

func (h *handlerAuth) Register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	request := new(userdto.CreateUserRequest)
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	password, err := bcrypt.HashingPassword(request.Password)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}

	dataEmail, _ := h.AuthRepository.LoginUser(request.Email)
	if dataEmail.Email != "" {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: "Email has been used"}
		json.NewEncoder(w).Encode(response)
		return
	}

	user := models.User{
		Name:     request.Name,
		Email:    request.Email,
		Password: password,
	}

	data, err := h.AuthRepository.RegisterUser(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: data}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerAuth) Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	request := new(authdto.AuthRequest)
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	user := models.User{
		Email:    request.Email,
		Password: request.Password,
	}

	user, err := h.AuthRepository.LoginUser(user.Email)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: "Email not registered!"}
		json.NewEncoder(w).Encode(response)
		return
	}

	isValid := bcrypt.CheckPasswordHash(request.Password, user.Password)
	if !isValid {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: "Wrong password!"}
		json.NewEncoder(w).Encode(response)
		return
	}

	gnrtToken := jwt.MapClaims{}
	gnrtToken["id"] = user.ID
	gnrtToken["exp"] = time.Now().Add(time.Hour * 3).Unix()
	gnrtToken["name"] = user.Name
	gnrtToken["email"] = user.Email
	gnrtToken["password"] = user.Password
	gnrtToken["photo"] = user.Image
	gnrtToken["greeting"] = user.Greeting
	gnrtToken["bestArt"] = user.BestArt
	gnrtToken["following"] = user.Following

	token, err := jwtToken.GenerateToken(&gnrtToken)
	if err != nil {
		fmt.Println("Unauthorize")
		return
	}

	AuthResponse := authdto.AuthResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Password:  user.Password,
		Token:     token,
		Greeting:  user.Greeting,
		BestArt:   user.BestArt,
		Image:     user.Image,
		Following: user.Following,
	}

	w.Header().Set("Content-Type", "application/json")
	response := dto.SuccessResult{Code: http.StatusOK, Data: AuthResponse}
	json.NewEncoder(w).Encode(response)
}

// func (h *handlerAuth) Login(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")

// 	request := new(authdto.RequestLogin)
// 	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
// 		json.NewEncoder(w).Encode(response)
// 		return
// 	}

// 	user := models.User{
// 		Email:    request.Email,
// 		Password: request.Password,
// 	}

// 	// Check email
// 	user, err := h.AuthRepository.Login(user.Email)
// 	if err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
// 		json.NewEncoder(w).Encode(response)
// 		return
// 	}

// 	// Check password
// 	isValid := bcrypt.CheckPasswordHash(request.Password, user.Password)
// 	if !isValid {
// 		w.WriteHeader(http.StatusBadRequest)
// 		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: "wrong email or password"}
// 		json.NewEncoder(w).Encode(response)
// 		return
// 	}

// 	//generate token
// 	claims := jwt.MapClaims{}
// 	claims["id"] = user.ID
// 	claims["name"] = user.Name
// 	claims["email"] = user.Email
// 	claims["exp"] = time.Now().Add(time.Hour * 2).Unix() // 2 hours expired

// 	token, errGenerateToken := jwtToken.GenerateToken(&claims)
// 	if errGenerateToken != nil {
// 		log.Println(errGenerateToken)
// 		fmt.Println("Unauthorize")
// 		return
// 	}

// 	loginResponse := authdto.ResponseLogin{
// 		ID:    user.ID,
// 		Name:  user.Name,
// 		Email: user.Email,
// 		Token: token,
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	response := dto.SuccessResult{Code: http.StatusOK, Data: loginResponse}
// 	json.NewEncoder(w).Encode(response)

// }

func (h *handlerAuth) CheckAuth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userInfo := r.Context().Value("authInfo").(jwt.MapClaims)
	userId := int(userInfo["id"].(float64))

	user, err := h.AuthRepository.Getuser(userId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	CheckAuthResponse := authdto.CheckAuthResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Password:  user.Password,
		Greeting:  user.Greeting,
		BestArt:   user.BestArt,
		Image:     user.Image,
		Following: user.Following,
	}

	w.Header().Set("Content-Type", "application/json")
	response := dto.SuccessResult{Code: http.StatusOK, Data: CheckAuthResponse}
	json.NewEncoder(w).Encode(response)
}
