package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	dto "waysgallery/dto/result"
	userdto "waysgallery/dto/user"
	"waysgallery/models"
	"waysgallery/repositories"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/gorilla/mux"
)

type handlerUser struct {
	UserRepository repositories.UserRepository
}

func HandlerUser(UserRepository repositories.UserRepository) *handlerUser {
	return &handlerUser{UserRepository}
}

func (h *handlerUser) FindUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "aplication/json")

	dataUser := r.URL.Query().Get("search")
	date := r.URL.Query().Get("date")

	fmt.Println(dataUser)
	fmt.Println(date)

	users, err := h.UserRepository.FindUsers()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	for i, p := range users {
		users[i].Image = os.Getenv("PATH_FILE") + p.Image
	}

	for i, p := range users {
		users[i].BestArt = os.Getenv("PATH_FILE") + p.BestArt
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: users}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerUser) GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "aplication-json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	user, err := h.UserRepository.GetUser(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	user.Image = os.Getenv("PATH_FILE") + user.Image
	user.BestArt = os.Getenv("PATH_FILE") + user.BestArt

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: user}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerUser) UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	dataImage := r.Context().Value("dataFile")
	dataBestArt := r.Context().Value("dataBestArt")
	filepath := ""
	fileArt := ""

	var ctx = context.Background()
	var CLOUD_NAME = os.Getenv("CLOUD_NAME")
	var API_KEY = os.Getenv("API_KEY")
	var API_SECRET = os.Getenv("API_SECRET")

	if dataImage != nil {
		filepath = dataImage.(string)
	}

	if dataBestArt != nil {
		fileArt = dataBestArt.(string)
	}

	cld, _ := cloudinary.NewFromParams(CLOUD_NAME, API_KEY, API_SECRET)

	resp, err2 := cld.Upload.Upload(ctx, filepath, uploader.UploadParams{Folder: "waysgallery"})
	resp2, _ := cld.Upload.Upload(ctx, fileArt, uploader.UploadParams{Folder: "waysgallery"})

	if err2 != nil {
		fmt.Println(err2.Error())
	}

	request := userdto.UpdateUserRequest{
		Name:     r.FormValue("name"),
		Greeting: r.FormValue("greeting"),
	}

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	user := models.User{}

	user.ID = id

	if request.Name != "" {
		user.Name = request.Name
	}

	if request.Greeting != "" {
		user.Greeting = request.Greeting
	}

	if filepath != "" {
		user.Image = resp.SecureURL
	}

	if fileArt != "" {
		user.BestArt = resp2.SecureURL
	}

	data, err := h.UserRepository.UpdateUser(user, id)
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
