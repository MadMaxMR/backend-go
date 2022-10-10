package modelos

type UserTipe struct {
	ID          string `json:"id" gorm:"primary_key;type:varchar(10);not null"`
	Description string `json:"description" gorm:"type:varchar(250)"`
}
