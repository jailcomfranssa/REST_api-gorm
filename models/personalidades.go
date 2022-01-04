package models

type Personalidade struct {
	Id       int    `json:"id" gorm:"primaryKey"`
	Nome     string `json:"nome"`
	Historia string `json:"historia"`
}
