package lib

import (
	"log"
	"time"

	"github.com/bmorrisondev/go-utils"
)

type Post struct {
	Id          *int64     `json:"id"`
	Text        *string    `json:"text"`
	ParentId    *int64     `json:"parentId"`
	SendAt      *time.Time `json:"sendAt"`
	ResendAt    *time.Time `json:"resendAt"`
	ThreadCount *int       `json:"threadCount"`
	Status      *int       `json:"status"`
	Media       []Media    `json:"media"`
}

func (p *Post) GetSendAtSqlTimestamp() *string {
	if p.SendAt != nil {
		returnValue := p.SendAt.Format("2006-01-02 15:04:05")
		return &returnValue
	}
	return nil
}

func (p *Post) GetResendAtSqlTimestamp() *string {
	if p.ResendAt != nil {
		returnValue := p.ResendAt.Format("2006-01-02 15:04:05")
		return &returnValue
	}
	return nil
}

func (p *Post) MediaJson() *string {
	if p.Media != nil {
		jstr, err := utils.ConvertToJsonString(p.Media)
		if err != nil {
			log.Printf("error while converting media to json: %v", err)
		}
		return &jstr
	}
	return nil
}

type Media struct {
	Id         *string `json:"id"`
	PreviewUrl *string `json:"preview_url"`
}

type User struct {
	Id        *int64
	LastLogin time.Time
}
