package main

type Configuration struct {
	Database  Database
	BasicAuth BasicAuth
}

type Database struct {
	Adapter  string
	User     string
	Database string
	Password string
	Host     string
	Port     string
}

type BasicAuth struct {
	User     string
	Password string
}
