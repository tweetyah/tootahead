package lib

import "time"

type Tweet struct {
	Id          *int64     `json:"id"`
	Text        *string    `json:"text"`
	ParentId    *int64     `json:"parentId"`
	SendAt      *time.Time `json:"sendAt"`
	RetweetAt   *time.Time `json:"retweetAt"`
	ThreadCount *int       `json:"threadCount"`
}

func (t *Tweet) GetSendAtSqlTimestamp() *string {
	if t.SendAt != nil {
		returnValue := t.SendAt.Format("2006-01-02 15:04:05")
		return &returnValue
	}
	return nil
}

func (t *Tweet) GetRetweetAtSqlTimestamp() *string {
	if t.RetweetAt != nil {
		returnValue := t.RetweetAt.Format("2006-01-02 15:04:05")
		return &returnValue
	}
	return nil
}
