package handlers

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"
	"time"
	postdto "waysgallery/dto/post"
	dto "waysgallery/dto/result"
	"waysgallery/models"
	"waysgallery/repositories"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/mux"
)

type handlerPost struct {
	PostRepository repositories.PostRepository
}

func HandlerPost(PostRepository repositories.PostRepository) *handlerPost {
	return &handlerPost{PostRepository}
}

func (h *handlerPost) FindPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	post, err := h.PostRepository.FindPosts()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	for i, p := range post {
		imagePath := os.Getenv("PATH_FILE") + p.Image1
		post[i].Image1 = imagePath
	}

	for i, p := range post {
		imagePath := os.Getenv("PATH_FILE") + p.Image2
		post[i].Image2 = imagePath
	}

	for i, p := range post {
		imagePath := os.Getenv("PATH_FILE") + p.Image3
		post[i].Image3 = imagePath
	}

	for i, p := range post {
		imagePath := os.Getenv("PATH_FILE") + p.Image4
		post[i].Image4 = imagePath
	}

	for i, p := range post {
		imagePath := os.Getenv("PATH_FILE") + p.Image5
		post[i].Image5 = imagePath
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: post}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerPost) FindPostsByUserId(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	post, err := h.PostRepository.FindPostsByUserId(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	for i, p := range post {
		imagePath := os.Getenv("PATH_FILE") + p.Image1
		post[i].Image1 = imagePath
	}

	for i, p := range post {
		imagePath := os.Getenv("PATH_FILE") + p.Image2
		post[i].Image2 = imagePath
	}

	for i, p := range post {
		imagePath := os.Getenv("PATH_FILE") + p.Image3
		post[i].Image3 = imagePath
	}

	for i, p := range post {
		imagePath := os.Getenv("PATH_FILE") + p.Image4
		post[i].Image4 = imagePath
	}

	for i, p := range post {
		imagePath := os.Getenv("PATH_FILE") + p.Image5
		post[i].Image5 = imagePath
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: post}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerPost) GetPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	var post models.Post
	post, err := h.PostRepository.GetPost(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	post.Image1 = os.Getenv("PATH_FILE") + post.Image1
	post.Image2 = os.Getenv("PATH_FILE") + post.Image2
	post.Image3 = os.Getenv("PATH_FILE") + post.Image3
	post.Image4 = os.Getenv("PATH_FILE") + post.Image4
	post.Image5 = os.Getenv("PATH_FILE") + post.Image5

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: post}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerPost) CreatePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	adminInfo := r.Context().Value("authInfo").(jwt.MapClaims)
	userid := int(adminInfo["id"].(float64))
	dataUpload := r.Context().Value("dataPost")
	dataUpload2 := r.Context().Value("dataPost2")
	dataUpload3 := r.Context().Value("dataPost3")
	dataUpload4 := r.Context().Value("dataPost4")
	dataUpload5 := r.Context().Value("dataPost5")
	filepath := ""
	filepath2 := ""
	filepath3 := ""
	filepath4 := ""
	filepath5 := ""

	if dataUpload != nil {
		filepath = dataUpload.(string)
	}
	if dataUpload2 != nil {
		filepath2 = dataUpload2.(string)
	}
	if dataUpload3 != nil {
		filepath3 = dataUpload3.(string)
	}
	if dataUpload4 != nil {
		filepath4 = dataUpload4.(string)
	}
	if dataUpload5 != nil {
		filepath5 = dataUpload5.(string)
	}

	input := time.Now()

	dateParse := input.Format("2 Jan 2006 15:04")

	request := postdto.CreatePostRequest{
		Title: r.FormValue("title"),
		Desc:  r.FormValue("desc"),
		Date:  dateParse,
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	post := models.Post{
		UserID: userid,
		Title:  request.Title,
		Desc:   request.Desc,
		Date:   request.Date,
		Image1: filepath,
		Image2: filepath2,
		Image3: filepath3,
		Image4: filepath4,
		Image5: filepath5,
	}

	post, err = h.PostRepository.CreatePost(post)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	post, err = h.PostRepository.GetPost(post.ID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: post}
	json.NewEncoder(w).Encode(response)
}
