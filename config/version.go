package config

var APP_NAME string

const APP_VERSION = "v1.0.3"

func LoadVersion() {
	APP_NAME = AppConfigure.Name + "-" + AppConfigure.Env
}
