package database

type ReporterModel struct {
	Id             string `gorm:"primary_key"`
	Tag string
}

func (ReporterModel) TableName() string {
	return "reporters"
}