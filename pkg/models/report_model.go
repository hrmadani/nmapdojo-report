package models

import (
	"log"
	"time"

	"github.com/hrmadani/nmapdojo-report/pkg/config"

	"gorm.io/gorm"
)

var (
	db         *gorm.DB
	reportType ReportType
	report     Report
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
	ReportId  uint   `json:"report_id"`
	UserId    uint   `json:"user_id"`
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

//Find a report
func (r *Report) FindById(id uint) Report {
	return report
}

//Save to the reports table
func (r *Report) Save(userReport UserReport) (uint, error) {
	GetReportTypeProperties(userReport.ReportType)
	//Save message to the reports table
	var CreatedAt = userReport.CreatedAt
	r.ID = 0
	r.ExpireTime = CreatedAt.Add(time.Duration(reportType.LifeSpan) * time.Minute)
	r.ReportType = userReport.ReportType

	result := db.Create(r)
	if result.Error != nil {
		log.Printf("[x] Inserting report fail: %v \n", result.Error)
		return 0, result.Error
	}
	return r.ID, nil
}

//Save to the report_logs table
func (rl *ReportLog) Save(userReport UserReport, reportID uint) error {
	//Todo: Save message to the reports table
	rl.ID = 0
	rl.ReportId = reportID
	rl.UserId = uint(userReport.UserId)
	rl.Action = userReport.Action
	rl.CreatedAt = userReport.CreatedAt

	result := db.Create(rl)
	if result.Error != nil {
		log.Printf("[x] Inserting report log fail: %v \n", result.Error)
		return result.Error
	}
	return nil
}

//updateExpireTime with + or -
func (r *Report) UpdateExpireTime(userReport UserReport) {
	//Todo: Change ExpireTime message in the reports table
	GetReportTypeProperties(userReport.ReportType)

	//Find the report by id and fill the report variable
	r.FindById(uint(userReport.ReportId))

	switch userReport.Action {
	case "like":
		report.ExpireTime = report.ExpireTime.Add(time.Duration(reportType.FeedbackEffect) * time.Second)
	default:
		report.ExpireTime = report.ExpireTime.Add(time.Duration(-(reportType.FeedbackEffect)) * time.Second)
	}
	db.Model(&report).Where("id = ?", userReport.ReportId).Update("expire_time", report.ExpireTime)
}
