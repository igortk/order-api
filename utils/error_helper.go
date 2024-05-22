package utils

import log "github.com/sirupsen/logrus"

func IsError(err error, msg string) {
	if err != nil {
		log.WithError(err).Errorf("%s", msg)
	}
}
