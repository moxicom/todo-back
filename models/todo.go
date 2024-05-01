package models

type User struct {
	Id       int    `json:"-" gorm:"primaryKey"`
	Name     string `json:"name" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"` // Not hash for a while

	UserLists []UserList `gorm:"foreignKey:UserId" json:"-"`
}

type UserList struct {
	Id     int `gorm:"primaryKey" json:"-"`
	UserId int
	ListId int
}

type TodoList struct {
	Id          int    `json:"id" gorm:"primaryKey"`
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`

	UserLists []UserList `gorm:"foreignKey:ListId" json:"-"`
	ItemLists []ListItem `gorm:"foreignKey:ListId" json:"-"`
}

type ListItem struct {
	Id     int `gorm:"primaryKey"`
	ListId int
	ItemId int
}

type Item struct {
	Id          int    `json:"id" gorm:"primaryKey"`
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	Done        bool   `json:"done"`

	ListItems []ListItem `gorm:"foreignKey:ItemId" json:"-"`
}
