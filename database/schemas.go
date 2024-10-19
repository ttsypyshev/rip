package database

import "time"

type Lang struct {
	ID               int    `gorm:"primaryKey"`
	Name             string `gorm:"size:255"`
	ImgLink          string `gorm:"size:255"`
	ShortDescription string `gorm:"size:255"`
	Author           string `gorm:"size:255"`
	Year             string `gorm:"size:4"`
	Version          string `gorm:"size:50"`
	Description      string `gorm:"type:text"`
	List             string `gorm:"type:text"`
}

type Project struct {
	ID           int `gorm:"primaryKey"`
	CreationDate time.Time
	CreationTime time.Time
	Status       int
}

type File struct {
	ID        int    `gorm:"primaryKey"`
	IDLang    int    `gorm:"column:id_lang"`
	IDProject int    `gorm:"column:id_project"`
	Code      string `gorm:"type:text"`
}
