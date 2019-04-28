package repositoryrequester

import (
	"encoding/json"
	"log"
	"net/http"
	"regexp"
	"strings"
)

type jsonRepositories struct {
	Value []struct {
		ID      string `json:"id"`
		Name    string `json:"name"`
		URL     string `json:"url"`
		Project struct {
			ID         string `json:"id"`
			Name       string `json:"name"`
			URL        string `json:"url"`
			State      string `json:"state"`
			Revision   int    `json:"revision"`
			Visibility string `json:"visibility"`
		} `json:"project"`
		DefaultBranch string `json:"defaultBranch,omitempty"`
		RemoteURL     string `json:"remoteUrl"`
	} `json:"value"`
	Count int `json:"count"`
}

type Repository struct {
	Name           string
	RemoteURL      string
	BoundedContext string
}

type RepoProvider interface {
	ListRepos() map[string][]Repository
}

type TfsRepoProvider struct {
	BaseURL   string
	UserName  string
	AuthToken string
}

func ListRepos(r RepoProvider) map[string][]Repository {
	return r.ListRepos()
}

func (t TfsRepoProvider) ListRepos() map[string][]Repository {
	var repos map[string][]Repository
	repos = make(map[string][]Repository)

	req, err := http.NewRequest("GET", t.BaseURL+"repositories", nil)
	req.SetBasicAuth(t.UserName, t.AuthToken)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	cli := &http.Client{}
	resp, err := cli.Do(req)
	if err != nil {
		log.Fatal("request failed", err)
		return repos
	}

	if resp.StatusCode != 200 {
		log.Fatalln("Was not possible to retrive repositories")
		return repos
	}

	var jsonRepositories jsonRepositories
	json.NewDecoder(resp.Body).Decode(&jsonRepositories)
	log.Println("Number of repos", len(jsonRepositories.Value))

	for _, element := range jsonRepositories.Value {
		boundedContext, name := identifyBoundedContextAndName(element.Name)

		var repo Repository

		repo.Name = name
		repo.RemoteURL = element.RemoteURL
		repo.BoundedContext = boundedContext
		repos[boundedContext] = append(repos[boundedContext], repo)
	}

	return repos
}

func identifyBoundedContextAndName(input string) (string, string) {
	r, _ := regexp.Compile("([a-zA-Z]+)BC.")
	boundedContext := r.FindString(input)
	if boundedContext == "" {
		boundedContext = "Undefined"
		return boundedContext, input
	}
	name := strings.Replace(input, boundedContext, "", -1)

	return strings.Trim(boundedContext, "."), name
}
