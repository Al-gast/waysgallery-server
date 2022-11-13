package postdto

type CreatePostRequest struct {
	UserID int    `json:"userID"`
	Title  string `json:"title"`
	Desc   string `json:"desc"`
	Date   string `json:"date"`
	Image1 string `json:"image1"`
	Image2 string `json:"image2"`
	Image3 string `json:"image3"`
	Image4 string `json:"image4"`
	Image5 string `json:"image5"`
}

type PostResponse struct {
	ID     int    `json:"id"`
	Title  string `json:"title" `
	Desc   string `json:"desc"`
	Date   string `json:"date"`
	UserID int    `json:"userID"`
	Image1 string `json:"image1" `
	Image2 string `json:"image2" `
	Image3 string `json:"image3" `
	Image4 string `json:"image4" `
	Image5 string `json:"image5" `
}
