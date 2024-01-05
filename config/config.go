package config

import (
	"os"
	"path/filepath"

	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

const (
	projectName = "go-spotify-cli"
)

type Config struct {
	ClientId        string
	ClientSecret    string
	RequestedScopes string
}

type EnvVarConfig struct {
	ClientId     string `yaml:"ClientId"`
	ClientSecret string `yaml:"ClientSecret"`
}

type CombinedTokenStructure struct {
	ModifyToken      UserModifyTokenStructure      `yaml:"ModifyToken"`
	ReadToken        UserReadTokenStructure        `yaml:"ReadToken"`
	LibraryReadToken UserLibraryReadTokenStructure `yaml:"LibraryReadToken"`
}

type UserModifyTokenStructure struct {
	UserModifyToken          string `yaml:"UserModifyToken"`
	UserModifyRefreshToken   string `yaml:"UserModifyRefreshToken"`
	UserModifyTokenExpiresIn int64  `yaml:"UserModifyTokenExpiresIn"`
}

type UserReadTokenStructure struct {
	UserReadToken          string `yaml:"UserReadToken"`
	UserReadRefreshToken   string `yaml:"UserReadRefreshToken"`
	UserReadTokenExpiresIn int64  `yaml:"UserReadTokenExpiresIn"`
}

type UserLibraryReadTokenStructure struct {
	UserLibraryReadToken          string `yaml:"UserLibraryReadToken"`
	UserLibraryReadRefreshToken   string `yaml:"UserLibraryReadRefreshToken"`
	UserLibraryReadTokenExpiresIn int64  `yaml:"UserLibraryReadTokenExpiresIn"`
}

var GlobalConfig Config

func LoadConfiguration() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		logrus.WithError(err).Error("Error getting home directory")
		return
	}

	folderPath := filepath.Join(homeDir, "."+projectName)
	filePath := filepath.Join(folderPath, projectName+".yaml")

	data, err := os.ReadFile(filePath)
	if err != nil {
		return
	}

	var config EnvVarConfig
	// Unmarshal the YAML data into a Configuration struct
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		logrus.WithError(err).Error("Error unmarshalling YAML data")
		return
	}

	GlobalConfig = Config{
		ClientId:     config.ClientId,
		ClientSecret: config.ClientSecret,
	}
}

func VerifyConfigExists() bool {
	// Get the home directory for the current user
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return false
	}

	// Define the folder and file paths
	folderPath := filepath.Join(homeDir, "."+projectName)
	filePath := filepath.Join(folderPath, projectName+".yaml")

	// Check if the folder exists
	folderExists, err := os.Stat(folderPath)
	if err != nil || os.IsNotExist(err) || !folderExists.IsDir() {
		return false
	}

	// Check if the file exists
	_, err = os.Stat(filePath)

	if err != nil {
		return false
	}

	if len(GlobalConfig.ClientSecret) == 0 || len(GlobalConfig.ClientId) == 0 {
		return false
	}

	return true
}
