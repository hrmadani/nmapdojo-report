package models

import (
	"time"

	"github.com/hrmadani/nmapdojo-report/pkg/config"

	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

type UserReport struct {
	gorm.Model
	ReportId   int    `json:"report_id"`   //ziro is null
	ReportType string `json:"report_type"` //Police - Traffic Jam - Car Crash
	UserId     int    `json:"user_id"`
	Action     string `json:"action"` //add , like , dislike
	CreatedAt  time.Time
}

//All Report
type Report struct {
	gorm.Model
	ReportType int       `json:"report_type"`
	ExpireTime time.Time `gorm:"index" json:"expire_time"`
}

//The Log Store
//Which User
//Which Action (Add , Like, Dislike)
//Which Report
type ReportLog struct {
	gorm.Model
	ReportId  int    `json:"report_id"`
	UserId    int    `json:"user_id"`
	Action    string `json:"action"` //add , like , dislike
	CreatedAt time.Time
}

//Database connection
func init() {
	//Connect to database
	config.Connect()
	db = config.GetDB()
}

//Save to the reports table
func (r *Report) save() error {
	//Todo: Save message to the reports table
	return nil
}

//Save to the report_logs table
func (rl *ReportLog) save() error {
	//Todo: Save message to the reports table
	return nil
}

//updateExpireTime with + or -
func (r *Report) updateExpireTime() error {
	//Todo: Save message to the reports table
	return nil
}
