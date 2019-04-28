package parser

import (
	"log"

	"gopkg.in/yaml.v2"
)

type repositoryOwner struct {
	RepositoryName string `yaml:"repositoryName"`
	Owners         []struct {
		Name         string `yaml:"name"`
		Email        string `yaml:"email"`
		SlackChannel string `yaml:"slackChannel"`
	} `yaml:"owners"`
	Description string `yaml:"description"`
}

func Parse(input []byte) {
	var repositoryOwner repositoryOwner

	err := yaml.Unmarshal(input, &repositoryOwner)
	if err != nil {
		log.Fatalln("error: ", err)
	}

	log.Println(repositoryOwner.Owners)
}
