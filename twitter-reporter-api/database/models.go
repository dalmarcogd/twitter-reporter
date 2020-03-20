package database

type AccountModel struct {
	Id             string `gorm:"primary_key"`
	DocumentNumber string
}

func (AccountModel) TableName() string {
	return "accounts"
}