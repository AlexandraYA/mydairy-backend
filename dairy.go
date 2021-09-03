package dairy

type TasksList struct {
	Id          int    `json:"id" db:"id"`
	DateStart   string `json:"datestart" db:"datestart" binding:"required"`
	DateEnd     string `json:"dateend" db:"dateend" binding:"required"`
	DateCreate  string `json:"datecreate" db:"datecreate" binding:"required"`
	Title       string `json:"title" db:"title" binding:"required"`
	Description string `json:"description" db:"description"`
	Done        bool   `json:"done" db:"done"`
}

type UsersTasks struct {
	Id     int
	UserId int
	TaskId int
}
