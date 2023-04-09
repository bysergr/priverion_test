package utils

import (
	"errors"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type ENV struct {
	HOST_DB     string
	PORT_DB     int16
	PORT_API    int16
	USER_DB     string
	PASSWORD_DB string
	DATABASE    string
	JWT         string
}

var (
	env *ENV
)

// GetENV returns the environment variables
// Using the singleton pattern
func GetENV() ENV {

	if env == nil {
		err := godotenv.Load()
		if err != nil {
			panic(errors.New("error loading .env file"))
		}

		port_api, err := strconv.ParseInt(os.Getenv("PORT_API"), 10, 16)
		if err != nil {
			panic(errors.New("port api should be a number"))
		}

		port_db, err := strconv.ParseInt(os.Getenv("PORT_DB"), 10, 16)
		if err != nil {
			panic(errors.New("port database should be a number"))
		}

		env = &ENV{
			HOST_DB:     os.Getenv("HOST_DB"),
			PORT_DB:     int16(port_db),
			PORT_API:    int16(port_api),
			USER_DB:     os.Getenv("USER_DB"),
			PASSWORD_DB: os.Getenv("PASSWORD_DB"),
			DATABASE:    os.Getenv("DATABASE"),
			JWT:         os.Getenv("JWT"),
		}
	}

	return *env
}
