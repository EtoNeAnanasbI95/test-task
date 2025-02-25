package config

import (
	"log"
	"testing"
)

func TestMustLoadConfig(t *testing.T) {
	cfg := MustLoadConfig("../../.env")
	log.Println(cfg)
}
