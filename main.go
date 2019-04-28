package main

import (
	"encoding/csv"
	"io/ioutil"
	"log"
	"os"

	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
	"github.com/mmedum/repository-overview/pkg/parser"
	"github.com/mmedum/repository-overview/pkg/repositoryrequester"
)

type config struct {
	Username  string `env:"USERNAME" envDefault:"abe"`
	AuthToken string `env:"AUTHTOKEN"`
	BaseURL   string `env:"BASEURL"`
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalln("File .env not found, reading configuration from ENV")
		return
	}

	var cfg config
	if err := env.Parse(&cfg); err != nil {
		log.Fatalln("Failed to parse ENV")
	}

	yamlFile, err := ioutil.ReadFile("test/repo.yaml")

	if err != nil {
		log.Println("yamlFile.Get err ", err)
	}

	parser.Parse(yamlFile)

	tfs := repositoryrequester.TfsRepoProvider{BaseURL: cfg.BaseURL, UserName: cfg.Username, AuthToken: cfg.AuthToken}
	repos := repositoryrequester.ListRepos(tfs)
	log.Println("Number of repos:", len(repos))

	outputToCsv(repos)

}

func outputToCsv(boundedContexts map[string][]repositoryrequester.Repository) {
	file, err := os.Create("export.csv")
	if err != nil {
		log.Fatalln("Not possible to create export.csv, with err", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	headers := []string{
		"Bounded Context",
		"Name",
	}

	// write column headers
	writer.Write(headers)

	for key := range boundedContexts {
		for _, repository := range boundedContexts[key] {

			r := make([]string, 0, 1+len(headers))
			r = append(
				r,
				repository.BoundedContext,
				repository.Name,
			)

			writer.Write(r)
		}
	}
}
