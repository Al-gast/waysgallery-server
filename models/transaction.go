package models

type Transaction struct {
	ID          int    `json:"id" gorm:"primary_key:auto_increment"`
	AdminID     int    `json:"admin_id"`
	Admin       User   `json:"admin" gorm:"foreignKey:AdminID"`
	BuyerID     int    `json:"buyer_id"`
	Buyer       User   `json:"buyer" gorm:"foreignKey:BuyerID"`
	Title       string `json:"title" form:"title" gorm:"type: varchar(255)"`
	Desc        string `json:"desc" form:"desc" gorm:"type: varchar(255)"`
	StartDate   string `json:"startDate" form:"startDate" gorm:"type: varchar(255)"`
	EndDate     string `json:"endDate" form:"endDate" gorm:"type: varchar(255)"`
	Price       string `json:"price" form:"price" gorm:"type: varchar(255)"`
	Status      string `json:"status" form:"status" gorm:"type: varchar(255)"`
	ProjectDesc string `json:"projectDesc" form:"projectDesc" gorm:"type: varchar(255)"`
	Image1      string `json:"image1" form:"image1" gorm:"type: varchar(255)"`
	Image2      string `json:"image2" form:"image2" gorm:"type: varchar(255)"`
	Image3      string `json:"image3" form:"image3" gorm:"type: varchar(255)"`
	Image4      string `json:"image4" form:"image4" gorm:"type: varchar(255)"`
	Image5      string `json:"image5" form:"image5" gorm:"type: varchar(255)"`
}
