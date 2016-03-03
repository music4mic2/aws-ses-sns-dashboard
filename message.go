package main

type mail struct {
  timestamp string
  sendingAccountId string
  source string
  sourceArn string
  messageId string
  destination []string
}

type message struct {
  notificationType string
  mail mail
}
