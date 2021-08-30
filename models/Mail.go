package models

type Mail struct {
	Id      string `gorm:"not null;primary_key;" json:"id"`
	To      string `gorm:"not null;" json:"to"`
	Cc      string `json:"cc"`
	Bcc     string `json:"bcc"`
	Subject string `gorm:"not null;" json:"subject"`
	Content string `gorm:"not null;" json:"content"`
}
