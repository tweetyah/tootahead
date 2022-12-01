package lib

import "time"

type Post struct {
	Id          *int64     `json:"id"`
	Text        *string    `json:"text"`
	ParentId    *int64     `json:"parentId"`
	SendAt      *time.Time `json:"sendAt"`
	ResendAt    *time.Time `json:"resendAt"`
	ThreadCount *int       `json:"threadCount"`
	Status      *int       `json:"status"`
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

type User struct {
	Id        *int64
	LastLogin time.Time
}
