package config

import "os"

type Config struct {
	Port                     string
	BookingDBHost            string
	BookingDBPort            string
	BookingDBName            string
	BookingDBUser            string
	BookingDBPass            string
	NatsHost                 string
	NatsPort                 string
	NatsUser                 string
	NatsPass                 string
	CreateUserCommandSubject string
	CreateUserReplySubject   string
}

func NewConfig() *Config {
	return &Config{
		Port:                     os.Getenv("BOOKING_SERVICE_PORT"),
		BookingDBHost:            os.Getenv("BOOKING_DB_HOST"),
		BookingDBPort:            os.Getenv("BOOKING_DB_PORT"),
		BookingDBName:            os.Getenv("BOOKING_DB_NAME"),
		BookingDBUser:            os.Getenv("BOOKING_DB_USER"),
		BookingDBPass:            os.Getenv("BOOKING_DB_PASS"),
		NatsHost:                 os.Getenv("NATS_HOST"),
		NatsPort:                 os.Getenv("NATS_PORT"),
		NatsUser:                 os.Getenv("NATS_USER"),
		NatsPass:                 os.Getenv("NATS_PASS"),
		CreateUserCommandSubject: os.Getenv("CREATE_USER_COMMAND_SUBJECT"),
		CreateUserReplySubject:   os.Getenv("CREATE_USER_REPLY_SUBJECT"),
	}
}
