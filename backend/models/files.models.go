package models

import "backend/db"

type File struct {
	ID   string `gorm:"primary_key" json:"id"`
	Name string `gorm:"column:name" json:"name"`
	Uri  string `gorm:"column:uri" json:"uri"`
}

func (f *File) Save() error {
	tx := db.DB.Save(f)
	return tx.Error
}

func (f *File) Delete() error {
	tx := db.DB.Delete(f)
	return tx.Error
}

func GetAllFile() ([]*File, error) {
	var files []*File
	err := db.DB.
		Find(&files).Error
	return files, err
}
