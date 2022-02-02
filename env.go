package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"reflect"
	"strings"
)

func ReplaceEnv(config *Config, envType reflect.Type, envValue reflect.Value) {
	envFilePath := fmt.Sprintf("%s/.env", config.BootEnvConfig.General.BootDirectory)
	input, err := ioutil.ReadFile(envFilePath)
	
	if err != nil {
		fmt.Printf("Error: No .env file found in `%s`\n", config.BootEnvConfig.General.BootDirectory)
	}

	lines := strings.Split(string(input), "\n")
	for i, line := range lines {
		fields := reflect.VisibleFields(envType)
		for _, field := range fields {
			variable := field.Tag.Get("bootvar")
			if strings.Contains(line, fmt.Sprintf("%s=", variable)) {
				lines[i] = fmt.Sprintf("%s=%s", variable, envValue.FieldByName(field.Name))
			}
		}
	}

	output := strings.Join(lines, "\n")
	err = ioutil.WriteFile(".env", []byte(output), 0644)
	if err != nil {
		log.Fatalln(err)
	}
}