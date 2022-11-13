package transactiondto

type CreateTransactionRequest struct {
	AdminID   int    `json:"admin_id"`
	BuyerID   int    `json:"buyer_id"`
	Title     string `json:"title"`
	Desc      string `json:"desc"`
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
	Price     string `json:"price"`
	Status    string `json:"status"`
}

type UpdateTransactionRequest struct {
	Status      string `json:"status"`
	ProjectDesc string `json:"projectDesc"`
	Image1      string `json:"image1"`
	Image2      string `json:"image2"`
	Image3      string `json:"image3"`
	Image4      string `json:"image4"`
	Image5      string `json:"image5"`
}

type TransactionResponse struct {
	ID          int    `json:"id"`
	AdminID     int    `json:"admin_id"`
	BuyerID     int    `json:"buyer_id"`
	Title       string `json:"title"`
	Desc        string `json:"desc"`
	StartDate   string `json:"startDate"`
	EndDate     string `json:"endDate"`
	Price       string `json:"price"`
	Status      string `json:"status"`
	ProjectDesc string `json:"projectDesc"`
	Image1      string `json:"image1"`
	Image2      string `json:"image2"`
	Image3      string `json:"image3"`
	Image4      string `json:"image4"`
	Image5      string `json:"image5"`
}
