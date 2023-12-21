package models

type User struct {
	Id       int    `json:"-" gorm:"primaryKey"`
	Name     string `json:"name" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"` // Not hash for a while

	UserLists []UserList `gorm:"foreignKey:UserId"`
}

type UserList struct {
	Id     int `gorm:"primaryKey"`
	UserId int
	ListId int
}

type TodoList struct {
	Id          int    `json:"id" gorm:"primaryKey"`
	Title       string `json:"title"`
	Description string `json:"description"`

	UserLists []UserList `gorm:"foreignKey:ListId"`
	ItemLists []ListItem `gorm:"foreignKey:ListId"`
}

type ListItem struct {
	Id     int `gorm:"primaryKey"`
	ListId int
	ItemId int
}

type Item struct {
	Id          int    `json:"id" gorm:"primaryKey"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Done        bool   `json:"done"`

	ListItems []ListItem `gorm:"foreignKey:ItemId"`
}
