package main

type Mail struct {
  timestamp string
  sendingAccountId string
  source string
  sourceArn string
  messageId string
  destination []string
}

type Message struct {
  notificationType string
  mail Mail
}
