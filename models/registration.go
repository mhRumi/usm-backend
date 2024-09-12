package models

import "time"

type registration struct {
	Course_id         int64     `json:"id"`
	Student_id        int64     `json:"id"`
	IsApproved        bool      `json:"isapproved"`
	Registration_date time.Time `json:"createdat"`
}
