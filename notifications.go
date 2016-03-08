package main

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"time"
)

type StringArray []string

func (l StringArray) Value() (driver.Value, error) {
	return json.Marshal(l)
}

func (l *StringArray) Scan(input interface{}) error {
	return json.Unmarshal(input.([]byte), l)
}

type Mail struct {
	ID               uint        `json:"-"`
	Destination      StringArray `json:"destination" sql:"type:text"`
	MessageID        string      `json:"messageId"`
	SendingAccountID string      `json:"sendingAccountId"`
	Source           string      `json:"source"`
	SourceArn        string      `json:"sourceArn"`
	Timestamp        string      `json:"timestamp"`
}

type Delivery struct {
	ID                   uint        `json:"-"`
	ProcessingTimeMillis int         `json:"processingTimeMillis"`
	Recipients           StringArray `json:"recipients" sql:"type:text"`
	ReportingMTA         string      `json:"reportingMTA"`
	SMTPResponse         string      `json:"smtpResponse"`
	Timestamp            string      `json:"timestamp"`
}

type Bounce struct {
	ID            uint   `json:"-"`
	BounceSubType string `json:"bounceSubType"`
	BounceType    string `json:"bounceType"`
	FeedbackID    string `json:"feedbackId"`
	ReportingMTA  string `json:"reportingMTA"`
	Timestamp     string `json:"timestamp"`
}

type Notification struct {
	ID               uint   `json:"-"`
	NotificationType string `json:"notificationType"`
	Mail             Mail   `json:"mail"`
	MailID           sql.NullInt64
	Bounce           Bounce `json:"bounce"`
	BounceID         sql.NullInt64
	Delivery         Delivery `json:"delivery"`
	DeliveryID       sql.NullInt64
	CreatedAt        time.Time
}
