package models

type User struct {
	Id       int    `json:"-" gorm:"primaryKey"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"` // Not hash for a while
}

type UserList struct {
	Id     int `gorm:"primaryKey"`
	UserId int
	ListId int

	Users     []User     `gorm:"foreignKey:Id"`
	TodoLists []TodoList `gorm:"foreignKey:Id"`
}

type TodoList struct {
	Id          int    `json:"id" gorm:"primaryKey"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type ListItem struct {
	Id     int `gorm:"primaryKey"`
	ListId int
	ItemId int

	TodoLists []TodoList `gorm:"foreignKey:Id"`
	Items     []Item     `gorm:"foreignKey:Id"`
}

type Item struct {
	Id          int    `json:"id" gorm:"primaryKey"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}
