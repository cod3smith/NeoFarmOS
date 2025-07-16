package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type NeocortexConfig struct {
	DBUrl       string `yaml:"db_url"`
	SecretKey   string `yaml:"secret_key"`
	APIKey      string `yaml:"api_key"`
	NATSUrl     string `yaml:"nats_url"`
	MQTTUrl     string `yaml:"mqtt_url"`
	RedisUrl    string `yaml:"redis_url"`
	ChromaUrl   string `yaml:"chroma_url"`
	PostgresUrl string `yaml:"postgres_url"`
}

type NeohiveConfig struct {
	DBUrl     string `yaml:"db_url"`
	SecretKey string `yaml:"secret_key"`
	APIKey    string `yaml:"api_key"`
}

type NeoplantConfig struct {
	DBUrl     string `yaml:"db_url"`
	SecretKey string `yaml:"secret_key"`
	APIKey    string `yaml:"api_key"`
}

type NeofungiConfig struct {
	DBUrl     string `yaml:"db_url"`
	SecretKey string `yaml:"secret_key"`
	APIKey    string `yaml:"api_key"`
}

type NeoedgeConfig struct {
	DBUrl        string `yaml:"db_url"`
	SecretKey    string `yaml:"secret_key"`
	MQTTPassword string `yaml:"mqtt_password"`
}

type NeobiolabConfig struct {
	DBUrl     string `yaml:"db_url"`
	SecretKey string `yaml:"secret_key"`
	APIKey    string `yaml:"api_key"`
}

type NeosilkConfig struct {
	DBUrl     string `yaml:"db_url"`
	SecretKey string `yaml:"secret_key"`
	APIKey    string `yaml:"api_key"`
}

func LoadAllConfig() map[string]interface{} {
	f, err := os.Open("../../config/environments/dev/config.yaml")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var all map[string]interface{}
	decoder := yaml.NewDecoder(f)
	if err := decoder.Decode(&all); err != nil {
		log.Fatal(err)
	}
	return all
}

func GetNeocortexConfig() NeocortexConfig {
	f, err := os.Open("../../config/environments/dev/config.yaml")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	var all map[string]NeocortexConfig
	decoder := yaml.NewDecoder(f)
	if err := decoder.Decode(&all); err != nil {
		log.Fatal(err)
	}
	cfg, ok := all["neocortex"]
	if !ok {
		log.Fatal("neocortex config not found")
	}
	return cfg
}

func GetNeohiveConfig() NeohiveConfig {
	f, err := os.Open("../../config/environments/dev/config.yaml")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	var all map[string]NeohiveConfig
	decoder := yaml.NewDecoder(f)
	if err := decoder.Decode(&all); err != nil {
		log.Fatal(err)
	}
	cfg, ok := all["neohive"]
	if !ok {
		log.Fatal("neohive config not found")
	}
	return cfg
}

func GetNeoplantConfig() NeoplantConfig {
	f, err := os.Open("../../config/environments/dev/config.yaml")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	var all map[string]NeoplantConfig
	decoder := yaml.NewDecoder(f)
	if err := decoder.Decode(&all); err != nil {
		log.Fatal(err)
	}
	cfg, ok := all["neoplant"]
	if !ok {
		log.Fatal("neoplant config not found")
	}
	return cfg
}

func GetNeofungiConfig() NeofungiConfig {
	f, err := os.Open("../../config/environments/dev/config.yaml")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	var all map[string]NeofungiConfig
	decoder := yaml.NewDecoder(f)
	if err := decoder.Decode(&all); err != nil {
		log.Fatal(err)
	}
	cfg, ok := all["neofungi"]
	if !ok {
		log.Fatal("neofungi config not found")
	}
	return cfg
}

func GetNeoedgeConfig() NeoedgeConfig {
	f, err := os.Open("../../config/environments/dev/config.yaml")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	var all map[string]NeoedgeConfig
	decoder := yaml.NewDecoder(f)
	if err := decoder.Decode(&all); err != nil {
		log.Fatal(err)
	}
	cfg, ok := all["neoedge"]
	if !ok {
		log.Fatal("neoedge config not found")
	}
	return cfg
}

func GetNeobiolabConfig() NeobiolabConfig {
	f, err := os.Open("../../config/environments/dev/config.yaml")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	var all map[string]NeobiolabConfig
	decoder := yaml.NewDecoder(f)
	if err := decoder.Decode(&all); err != nil {
		log.Fatal(err)
	}
	cfg, ok := all["neobiolab"]
	if !ok {
		log.Fatal("neobiolab config not found")
	}
	return cfg
}

func GetNeosilkConfig() NeosilkConfig {
	f, err := os.Open("../../config/environments/dev/config.yaml")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	var all map[string]NeosilkConfig
	decoder := yaml.NewDecoder(f)
	if err := decoder.Decode(&all); err != nil {
		log.Fatal(err)
	}
	cfg, ok := all["neosilk"]
	if !ok {
		log.Fatal("neosilk config not found")
	}
	return cfg
} 