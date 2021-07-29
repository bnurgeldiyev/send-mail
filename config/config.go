package config

import (
	"encoding/json"
	"io/ioutil"
	"time"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

type Config struct {
	Addr     string `json:"addr"`
	Mail     string `json:"mail"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     string `json:"port"`
}

var Conf *Config
var isFirst bool = true

func ReadConfig(source string) (err error) {
	if !isFirst {
		for {
			if Conf != nil {
				break
			}
			time.Sleep(10 * time.Millisecond)
		}
		if Conf != nil {
			return nil
		} else {
			return errors.New("error Conf not found!!!")
		}
	}
	isFirst = false

	var raw []byte
	raw, err = ioutil.ReadFile(source)
	if err != nil {
		wMsg := "error reading config from file, creating new sample"
		log.Warn(wMsg)

		err = createDefaultConfig(source)
		if err != nil {
			eMsg := "error creating config file"
			log.WithError(err).Error(eMsg)
			err = errors.Wrap(err, eMsg)
			return
		}

		raw, err = ioutil.ReadFile(source)
		if err != nil {
			eMsg := "error reading config from file"
			log.WithError(err).Error(eMsg)
			err = errors.Wrap(err, eMsg)
			return
		}
	}
	err = json.Unmarshal(raw, &Conf)
	if err != nil {
		eMsg := "error parsing config from json"
		log.WithError(err).Error(eMsg)
		err = errors.Wrap(err, eMsg)
		Conf = nil
		return
	}
	return
}

func createDefaultConfig(source string) (err error) {
	c := Config{
		Addr:     "127.0.0.1:8081",
		Mail:     "example@gmail.com",
		Password: "pass",
		Host:     "smtp.gmail.com",
		Port:     "587",
	}

	b, err := json.MarshalIndent(c, "", "\t")

	if err != nil {
		eMsg := "error marshall config file"
		log.WithError(err).Error(eMsg)
		err = errors.Wrap(err, eMsg)
		return
	}

	err = ioutil.WriteFile(source, b, 0644)
	if err != nil {
		eMsg := "error creating config file"
		log.WithError(err).Error(eMsg)
		err = errors.Wrap(err, eMsg)
		return
	}
	return
}
