package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/jinzhu/gorm"
)

var (
	server           *httptest.Server
	reader           io.Reader
	notificationsUrl string
	DB               *gorm.DB
)

func init() {
	server = httptest.NewServer(NewRouter())
	notificationsUrl = fmt.Sprintf("%s/", server.URL)
	DB = dbInstance()
}

func TestProcessSubscription(t *testing.T) {
	payloadJson := `{
	  "Type" : "SubscriptionConfirmation",
	  "MessageId" : "165545c9-2a5c-472c-8df2-7ff2be2b3b1b",
	  "Token" : "2336412f37fb687f5d51e6e241d09c805a5a57b30d712f794cc5f6a988666d92768dd60a747ba6f3beb71854e285d6ad02428b09ceece29417f1f02d609c582afbacc99c583a916b9981dd2728f4ae6fdb82efd087cc3b7849e05798d2d2785c03b0879594eeac82c01f235d0e717736",
	  "TopicArn" : "arn:aws:sns:us-west-2:123456789012:MyTopic",
	  "Message" : "You have chosen to subscribe to the topic arn:aws:sns:us-west-2:123456789012:MyTopic.\nTo confirm the subscription, visit the SubscribeURL included in this message.",
	  "SubscribeURL" : "https://sns.us-west-2.amazonaws.com/?Action=ConfirmSubscription&TopicArn=arn:aws:sns:us-west-2:123456789012:MyTopic&Token=2336412f37fb687f5d51e6e241d09c805a5a57b30d712f794cc5f6a988666d92768dd60a747ba6f3beb71854e285d6ad02428b09ceece29417f1f02d609c582afbacc99c583a916b9981dd2728f4ae6fdb82efd087cc3b7849e05798d2d2785c03b0879594eeac82c01f235d0e717736",
	  "Timestamp" : "2012-04-26T20:45:04.751Z",
	  "SignatureVersion" : "1",
	  "Signature" : "EXAMPLEpH+DcEwjAPg8O9mY8dReBSwksfg2S7WKQcikcNKWLQjwu6A4VbeS0QHVCkhRS7fUQvi2egU3N858fiTDN6bkkOxYDVrY0Ad8L10Hs3zH81mtnPk5uvvolIC1CXGu43obcgFxeL3khZl8IKvO61GWB6jI9b5+gLPoBc1Q=",
	  "SigningCertURL" : "https://sns.us-west-2.amazonaws.com/SimpleNotificationService-f3ecfb7224c7233fe7bb5f59f96de52f.pem"
    }`

	reader = strings.NewReader(payloadJson)

	request, err := http.NewRequest("POST", notificationsUrl, reader)

	configuration := ReadConfiguration()
	auth := configuration.BasicAuth
	request.SetBasicAuth(auth.User, auth.Password)
	request.Header.Add("x-amz-sns-message-type", "SubscriptionConfirmation")

	res, err := http.DefaultClient.Do(request)

	if err != nil {
		t.Error(err)
	}

	if res.StatusCode != http.StatusOK {
		t.Errorf("Success expected: %d", res.StatusCode)
	}
}

func TestProcessDelivery(t *testing.T) {
	payloadJson := `{
	  "Type" : "Notification",
	  "MessageId" : "401349d2-854c-5217-b33c-0603cfd0333f",
	  "TopicArn" : "arn:aws:sns:us-west-2:821745794127:bee-deliveries",
	  "Message" : "{\"notificationType\":\"Delivery\",\"mail\":{\"timestamp\":\"2016-06-07T22:19:06.766Z\",\"source\":\"Ripley <noreply@beetrack.com>\",\"sourceArn\":\"arn:aws:ses:us-west-2:821745794127:identity/noreply@beetrack.com\",\"sendingAccountId\":\"821745794127\",\"messageId\":\"010101552cf2a28e-756b2cba-c111-43f6-b6cc-2c072ba9fb74-000000\",\"destination\":[\"beetrackripley@gmail.com\"]},\"delivery\":{\"timestamp\":\"2016-06-07T22:19:08.819Z\",\"processingTimeMillis\":2053,\"recipients\":[\"beetrackripley@gmail.com\"],\"smtpResponse\":\"250 2.0.0 OK 1465337948 fn4si2459578pac.157 - gsmtp\",\"reportingMTA\":\"a27-42.smtp-out.us-west-2.amazonses.com\"}}",
	  "Timestamp" : "2016-06-07T22:19:08.907Z",
	  "SignatureVersion" : "1",
	  "Signature" : "dq+oi7Agqxb2VEFo/sMUoyFA8yTZ32RrgMn+satcluHqh61hABF2JW5T9vzxzQT8mGHD6rBNXBCeFbeeXvyvVEqxYPmT5GqxZr+M4geZmtJ3+6+yt8S12S6dWHi88kIjlB18xU3I0jKLR6vi4ZgjNcJQZKSdGYzsx5h08DxVqRcW6tnRlETOlcdjH6td0rIjRvYt9VVnAfQQcTiOqFajADueF/jo274FQRakY+nc8rzbrQ/MtChUgi9csPLwPgkZoDDjVNmC2DDdNcvS60c2h6TQnyT4m32fyXKNAYH7d8P1DPkJfMIzOIj0W7DSPTzF+3/G7oTIYEvIzpepjiSNBA==",
	  "SigningCertURL" : "https://sns.us-west-2.amazonaws.com/SimpleNotificationService-bb750dd426d95ee9390147a5624348ee.pem",
	  "UnsubscribeURL" : "https://sns.us-west-2.amazonaws.com/?Action=Unsubscribe&SubscriptionArn=arn:aws:sns:us-west-2:821745794127:bee-deliveries:061dc405-c38f-4605-94b4-68829cf6a9cb"
	}`

	reader = strings.NewReader(payloadJson)

	request, err := http.NewRequest("POST", notificationsUrl, reader)

	configuration := ReadConfiguration()
	auth := configuration.BasicAuth
	request.SetBasicAuth(auth.User, auth.Password)

	var oldCount, newCount int
	DB.Table("mails").Count(&oldCount)
	res, err := http.DefaultClient.Do(request)
	DB.Table("mails").Count(&newCount)

	if err != nil {
		t.Error(err)
	}

	if res.StatusCode != http.StatusOK {
		t.Errorf("Success expected: %d", res.StatusCode)
	}

	if oldCount == newCount {
		t.Errorf("Mail count increment expected: %d", newCount+1)
	}
}
