package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

// Config represents the configuration options for MongoDB.
type Config struct {
	MongoClient MongoClient `yaml:"mongoClient"`
}

type MongoClient struct {
	MONGOURI          string `yaml:"MONGOURI"`
	Database          string `yaml:"database"`
	Collection        string `yaml:"collection"`
	Collection_Device string `yaml:"collection_device"`
	ConnectTimeout    int    `yaml:"connectTimeout"`
	SocketTimeout     int    `yaml:"socketTimeout"`
	ServerSelection   int    `yaml:"serverSelection"`
	Port              string `yaml:"port"`
}

// ReadConfigFromFile reads the configuration from a YAML file and returns a Config object.
func ReadConfigFromFile(filePath string) (Config, error) {
	var config Config

	// Read the YAML file
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return config, err
	}

	// Unmarshal the YAML data into the Config struct
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return config, err
	}

	return config, nil
}

// WriteConfigToFile writes the Config object to a YAML file.
func WriteConfigToFile(filePath string) error {
	config := Config{MongoClient: MongoClient{
		MONGOURI:        "mongodb+srv://test:testtoc@cluster0.04jgpbc.mongodb.net",
		Database:        "testdb",
		Collection:      "organization",
		ConnectTimeout:  5000,
		SocketTimeout:   10000,
		ServerSelection: 3000,
	}}
	// Marshal the Config object into YAML data
	data, err := yaml.Marshal(config)
	if err != nil {
		return err
	}

	// Write the YAML data to the file
	err = ioutil.WriteFile(filePath, data, 0644)
	if err != nil {
		return err
	}

	return nil
}
