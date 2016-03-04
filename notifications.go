package main

import (
	"database/sql/driver"
	"encoding/json"
)

type StringArray []string

func (l StringArray) Value() (driver.Value, error) {
	return json.Marshal(l)
}

func (l *StringArray) Scan(input interface{}) error {
	return json.Unmarshal(input.([]byte), l)
}

type Mail struct {
	ID               uint        `gjson:"-" gorm:"primary_key"`
	Destination      StringArray `json:"destination" sql:"type:text"`
	MessageID        string      `json:"messageId"`
	SendingAccountID string      `json:"sendingAccountId"`
	Source           string      `json:"source"`
	SourceArn        string      `json:"sourceArn"`
	Timestamp        string      `json:"timestamp"`
}

type Delivery struct {
	ProcessingTimeMillis int         `json:"processingTimeMillis"`
	Recipients           StringArray `json:"recipients" sql:"type:text"`
	ReportingMTA         string      `json:"reportingMTA"`
	SMTPResponse         string      `json:"smtpResponse"`
	Timestamp            string      `json:"timestamp"`
}

type Bounce struct {
	BounceSubType     string `json:"bounceSubType"`
	BounceType        string `json:"bounceType"`
	BouncedRecipients []struct {
		Action         string `json:"action"`
		DiagnosticCode string `json:"diagnosticCode"`
		EmailAddress   string `json:"emailAddress"`
		Status         string `json:"status"`
	} `json:"bouncedRecipients"`
	FeedbackID   string `json:"feedbackId"`
	ReportingMTA string `json:"reportingMTA"`
	Timestamp    string `json:"timestamp"`
}

//TODO polimorfic association

type BounceType struct {
	ID               uint   `gjson:"-" gorm:"primary_key"`
	NotificationType string `json:"notificationType"`
	Mail             Mail   `json:"mail"`
	Bounce           Bounce `json:"bounce"`
}

type DeliveryType struct {
	ID               uint     `gjson:"-" gorm:"primary_key"`
	NotificationType string   `json:"notificationType"`
	Mail             Mail     `json:"mail"`
	Delivery         Delivery `json:"delivery"`
}
