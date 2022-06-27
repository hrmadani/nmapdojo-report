package models

import (
	"fmt"
	"log"
	"time"

	"github.com/hrmadani/nmapdojo-report/pkg/config"

	"gorm.io/gorm"
)

var (
	db         *gorm.DB
	reportType ReportType
)

//The struct of report types includes the Police, Traffic Jam, and Car crash reports
type ReportType struct {
	gorm.Model
	Type           string `json:"report_type"` //Police - Traffic Jam - Car Crash
	LifeSpan       int    `json:"life_span"`
	FeedbackEffect int    `json:"feedback_effect"`
}

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
	ReportType string    `json:"report_type"` //Police - Traffic Jam - Car Crash
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

//Get the report type properties
func GetReportTypeProperties(userReportType string) ReportType {
	reportType = ReportType{Type: userReportType}
	db.First(&reportType)

	return reportType
}

//Save to the reports table
func (r *Report) Save(userReport UserReport) error {
	GetReportTypeProperties(userReport.ReportType)
	//Save message to the reports table
	var CreatedAt = userReport.CreatedAt
	r.ID = 0
	r.ExpireTime = CreatedAt.Add(time.Duration(reportType.LifeSpan) * time.Minute)
	r.ReportType = userReport.ReportType
	fmt.Printf("Inserted Report : %v \n", r)
	result := db.Create(r)
	if result.Error != nil {
		log.Printf("[x] Inserting report fail: %v \n", result.Error)
		return result.Error
	}

	fmt.Printf("Inserted Report ID: %v \n", r.ID)
	return nil
}

//Save to the report_logs table
func (rl *ReportLog) Save() error {
	//Todo: Save message to the reports table
	return nil
}

//updateExpireTime with + or -
func (r *Report) UpdateExpireTime() error {
	//Todo: Save message to the reports table
	return nil
}
