package lib

import (
	"bufio"
	"log"
	"os"
	"strings"
)

const (
	// ServiceUID AppSettings
	ServiceUID = "serviceUID"
	// DatabaseURL AppSettings
	DatabaseURL = "databaseURL"
	// CredentialsFile AppSettings
	CredentialsFile = "credentialsFile"
)

// About enums using iota: https://golang.org/ref/spec#Iota

// AppSettings represents application context
type AppSettings struct {
	// DatabaseURL represents database url
	DatabaseURL string
	// ServiceUID represents service unique identifier
	ServiceUID string
	// CredentialsFile represents the path for the database credentials file
	CredentialsFile string
}

// ReadFromFile reads settings from file
// ex. "firebase.properties"
func ReadFromFile(path string) *AppSettings {
	// Debug working dir: os.Getwd()
	file, err := os.Open(path)

	if err == nil {
		defer file.Close()

		properties := make(map[string]string)
		stat, _ := file.Stat()
		buffer := make([]byte, stat.Size())
		file.Read(buffer)

		scanner := bufio.NewScanner(strings.NewReader(string(buffer)))

		for scanner.Scan() {
			entry := strings.Split(scanner.Text(), "=")

			key := strings.Trim(entry[0], " \t")
			value := strings.Trim(entry[1], " \t")

			properties[key] = value
		}

		// TODO Remove the map and use reflection?
		databaseURL := properties[DatabaseURL]
		serviceUID := properties[ServiceUID]
		credentialsFile := properties[ServiceUID]
		log.Println(databaseURL)

		return &AppSettings{
			DatabaseURL:     databaseURL,
			ServiceUID:      serviceUID,
			CredentialsFile: credentialsFile,
		}
	}

	return new(AppSettings)
}