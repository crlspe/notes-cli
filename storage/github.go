package storage

import (
	"context"
	"encoding/base64"
	"log"
	"strings"

	"github.com/google/go-github/v62/github"
)

type GitHub struct {
	Username   string
	Repository string
	Token      string
	client     *github.Client
	context    context.Context
}

func NewGitHub(username string, repo string, token string) *GitHub {
	return &GitHub{
		Username:   username,
		Repository: repo,
		Token:      token,
	}
}

func (g GitHub) Initialize() {
	if g.Username == strings.TrimSpace("") {
		log.Fatal("Unauthorized: No Username provided")
		// ask for one or read from file
	}
	if g.Repository == strings.TrimSpace("") {
		log.Fatal("Unauthorized: No Repository provided")
		// ask for one or read from file
	}
	if g.Token == strings.TrimSpace("") {
		log.Fatal("UnAuthorized: No AuthToken provided")
		// ask for one or read from file
	}
}

func (g *GitHub) Authenticate() {
	g.context = context.Background()
	g.client = github.NewClient(nil).WithAuthToken(g.Token)
}

func (g *GitHub) createFile(filename string, fileOptions github.RepositoryContentFileOptions) error {
	var _, _, err = g.client.Repositories.CreateFile(
		g.context,
		g.Username,
		g.Repository,
		filename,
		&fileOptions)
	return err
}

func (g *GitHub) updateFile(filename string, fileOptions github.RepositoryContentFileOptions) error {
	var _, _, err = g.client.Repositories.UpdateFile(
		g.context,
		g.Username,
		g.Repository,
		filename,
		&fileOptions)
	return err
}

func (g *GitHub) getContent(path string) (*github.RepositoryContent, error) {
	var file, _, _, err = g.client.Repositories.GetContents(
		g.context,
		g.Username,
		g.Repository,
		path,
		&github.RepositoryContentGetOptions{},
	)
	return file, err
}

func (g *GitHub) GetFileContent(filename string) (string, error) {
	var content, err = g.getContent(filename)
	var stringContent, _ = base64.StdEncoding.DecodeString(*content.Content)
	return string(stringContent), err
}

func (g *GitHub) UpdateFile(filename string, content string, commitComments ...string) error {
	var file, err = g.getContent(filename)

	var sha *string
	if err == nil && file != nil {
		sha = file.SHA
	}

	var fileOptions = github.RepositoryContentFileOptions{
		Message: github.String(strings.Join(commitComments, "\n")),
		Content: []byte(content),
		SHA:     sha,
	}

	if sha == nil {
		g.createFile(filename, fileOptions)
	} else {
		g.updateFile(filename, fileOptions)
	}

	return err
}
