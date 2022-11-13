package models

type User struct {
	ID        int    `json:"id" gorm:"primary_key:auto_increment;"`
	Name      string `json:"name" form:"name" gorm:"type: varchar(255)"`
	Email     string `json:"email" form:"email" gorm:"type: varchar(255)"`
	Password  string `json:"password" form:"password" gorm:"type: varchar(255)"`
	Image     string `json:"image" form:"image" gorm:"type: varchar(255)"`
	Greeting  string `json:"greeting" form:"greeting" gorm:"type: varchar(255)"`
	BestArt   string `json:"bestArt" form:"bestArt" gorm:"type: varchar(255)"`
	Following string `json:"following" form:"following" gorm:"type: varchar(255)"`
	// Posts     []PostsUserResponse `json:"posts"`
}
