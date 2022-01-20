package config

import "gorm.io/gorm"


var Database *gorm.DB

func Connect(){
	var err error
	Database, err = gorm.Open(sqllite.Open('fibergorm.db', &gorm.Config{}))
}