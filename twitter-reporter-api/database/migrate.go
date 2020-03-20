package database


func Migrate() {
	GetConnection().AutoMigrate(&AccountModel{})
}