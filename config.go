package main

import (
	"log"
	"github.com/BurntSushi/toml"
)

type program struct {
	Name string
	Stdout string
	Stderr string
}

type programs struct {
	Program []program
}

var hostedPrograms programs

func _readConfigFile() {
	if (isExists(".webtail.toml")) {
		if _, err := toml.DecodeFile(".webtail.toml", &hostedPrograms); err != nil {                                                                                                                                 
			panic(err)                                                                                                                                                                                                 
			return                                                                                                                                                                                                     
		} 
	} else {
		log.Fatal(".webtail.toml not existed!")
	}
}

func _validateConfig() bool {
	for _, p := range hostedPrograms.Program {
		if (!isExists(p.Stderr)) {
			log.Fatal(p.Stderr + "not existed!")
			return false
		}
		if (!isExists(p.Stdout)) {
			log.Fatal(p.Stdout + "not existed!")
			return false
		}
	}
	return true
}

func loadConfig () programs {
	_readConfigFile()
	if (!_validateConfig()) {
		log.Fatal("Start Webtail Failed!")
	} else {
		log.Println("Configuration Validation Passed!")
	}
	return hostedPrograms
}
