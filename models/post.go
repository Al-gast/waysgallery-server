package models

type Post struct {
	ID     int    `json:"id" gorm:"primary_key:auto_increment;"`
	Title  string `json:"title" form:"title" gorm:"type: varchar(255)"`
	Desc   string `json:"desc" form:"desc" gorm:"type: varchar(255)"`
	Date   string `json:"date" form:"date" gorm:"type: varchar(255)"`
	UserID int    `json:"userID"`
	User   User   `json:"userId" gorm:"foreignKey:UserID"`
	Image1 string `json:"image1" form:"image1" gorm:"type: varchar(255)"`
	Image2 string `json:"image2" form:"image2" gorm:"type: varchar(255)"`
	Image3 string `json:"image3" form:"image3" gorm:"type: varchar(255)"`
	Image4 string `json:"image4" form:"image4" gorm:"type: varchar(255)"`
	Image5 string `json:"image5" form:"image5" gorm:"type: varchar(255)"`
}

// type PostsUserResponse struct {
// 	ID     int    `json:"id"`
// 	Title  string `json:"title"`
// 	Desc   string `json:"desc"`
// 	Date   string `json:"date"`
// 	Image1 string `json:"image1"`
// 	Image2 string `json:"image2"`
// 	Image3 string `json:"image3"`
// 	Image4 string `json:"image4"`
// 	Image5 string `json:"image5"`
// }

// func (PostsUserResponse) TableName() string {
// 	return "posts"
// }
