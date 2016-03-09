Dashboard SES and SNS notifications
===================

Simple app to receive notifications from Amazon SES through Amazon SNS and save into database(PostgreSQL). Includes dashboard to see list of all notifications.

### Prerequisites ###
 - Go 1.6
 - PostgreSQL 9.5

### Execute App ###
    go run *.go

### Test Api Notification ###

**Delivery**

    curl -H "Content-Type: application/json" -X POST -d '{"notificationType":"Delivery","mail":{"timestamp":"2016-03-08T13:33:38.968Z","source":"cre <noreply@beetrack.com>","sourceArn":"arn:aws:ses:us-west-2:821745794127:identity/noreply@beetrack.com","sendingAccountId":"821745794127","messageId":"01010153566edb18-8013fc96-45de-4100-9250-9d83cfbfb8e0-000000","destination":["dev@beetrack.com"]},"delivery":{"timestamp":"2016-03-08T13:33:43.119Z","processingTimeMillis":4151,"recipients":["claudio.ruiz@beetrack.com"],"smtpResponse":"250 2.0.0 OK 1457444023 uw2si4842253pac.4 - gsmtp","reportingMTA":"a27-22.smtp-out.us-west-2.amazonses.com"}}' http://localhost:8000/notifications

**Bounce**

    curl -H "Content-Type: application/json" -X POST -d '{"notificationType":"Bounce","bounce":{"bounceSubType":"Suppressed","bounceType":"Permanent","reportingMTA":"dns; amazonses.com","bouncedRecipients":[{"emailAddress":"test.test@test.cl","status":"5.1.1","diagnosticCode":"Amazon SES has suppressed sending to this address because it has a recent history of bouncing as an invalid address. For more information about how to remove an address from the suppression list, see the Amazon SES Developer Guide: http://docs.aws.amazon.com/ses/latest/DeveloperGuide/remove-from-suppressionlist.html ","action":"failed"}],"timestamp":"2016-03-08T12:58:42.692Z","feedbackId":"01010153564ede33-7c47febe-e52d-11e5-9cad-59beabbaf6a1-000000"},"mail":{"timestamp":"2016-03-08T12:58:42.000Z","messageId":"01010153564edcd3-8480191c-9f5f-4011-b31b-abbf645b8057-000000","destination":["teresa.valdez@abcdin.cl"],"sendingAccountId":"821745794127","sourceArn":"arn:aws:ses:us-west-2:821745794127:identity/noreply@beetrack.com","source":"ABCDin <noreply@beetrack.com>"}}' http://localhost:8000/notifications
