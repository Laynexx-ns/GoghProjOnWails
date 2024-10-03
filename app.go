package main

import (
	"context"
	"encoding/json"
	"fmt"
)

// App struct
type App struct {
	ctx context.Context
}

type APIResponse []interface{}
type Gist struct {
	Description string      `json:"description"`
	Public      bool        `json:"public"`
	Files       interface{} `json:"files"`
}

const BaseUrl = "htpps://api.github.com"

var gitHubResponse APIResponse

// GetPublicRepositories извлекает список публичных репозиториев из API GitHub через запрос GET
func (a *App) GetPublicRepositories() (APIResponse, error) {
	url := fmt.Sprintf("%s/repositories", BaseUrl)
	response, err := MakeGetRequest(url, "")
	if err != nil {
		return nil, err
	}

	errUnpack := json.Unmarshal(response, &gitHubResponse)
	if errUnpack != nil {
		return nil, errUnpack
	}
	return gitHubResponse, nil

}

// GetPublicGists извлекает список публичных репозиториев ищ API GitHub через запрос GET
func (a *App) GetPublicGists() (APIResponse, error) {
	url := fmt.Sprintf("%s/gists/public", BaseUrl)

	response, err := MakeGetRequest(url, "")
	if err != nil {
		return nil, err
	}

	errUnpack := json.Unmarshal(response, &gitHubResponse)
	if errUnpack != nil {
		return nil, errUnpack
	}
	return gitHubResponse, nil
}

// GetRepositoriesForAuthenticatedUser используется для получения списка частных репозиториев через запрос GET
func (a *App) GetRepositoriesForAuthenticatedUser() (APIResponse, error) {
	url := fmt.Sprintf("%s/gists", BaseUrl)
	response, err := MakeGetRequest(url, "")
	if err != nil {
		return nil, err
	}
	errUnpack := json.Unmarshal(response, &gitHubResponse)
	if errUnpack != nil {
		return nil, errUnpack
	}
	return gitHubResponse, nil
}

/*
GetMoreInformationFromURL используется для получения дополнительной информации о репозитории. Эта информация может быть историей коммитов, список участников
или списком ползователей, которые отметили репозиторий звездой. Она принимает два параметра: URL-адрес для вызова и токен аутентификации. Для публичных репозиториев токен
будет пустой строкой
*/
func (a *App) GetMoreInformationFromURL(url, token string) (APIResponse, error) {
	response, err := MakeGetRequest(url, token)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	errUnpack := json.Unmarshal(response, &gitHubResponse)
	if errUnpack != nil {
		return nil, errUnpack
	}
	return gitHubResponse, nil

}

/*
	используется для получения содержимого Gist. Эта функция принимает URL для необработанного содержимого GIST и токен аутентификации (пустая строка для публичных Gist)

она возвращает строку, соответствует содежимому Gist
*/
func (a *App) GetGistContent(url, token string) (APIResponse, error) {
	response, err := MakeGetRequest(url, token)
	//обработка ошибок
	if err != nil {
		return nil, err
	}
	errUnpack := json.Unmarshal(response, &gitHubResponse)
	if errUnpack != nil {
		return nil, errUnpack
	}
	return gitHubResponse, nil

}

// CreateNewGist  используется для создания нового Gist для аутентифицированного пользователя. Эта функция принимает два параметра Gist, который нужно создать, а так же токен
func (a *App) CreateNewGist(gist Gist, token string) (interface{}, error) {
	var githubResponse interface{}

	requestBody, _ := json.Marshal(gist)
	url := fmt.Sprintf("%s/gitsts", BaseUrl)
	response, err := MakePostRequest(url, token, requestBody)
	if err != nil {
		return nil, err
	}
	errUnpack := json.Unmarshal(response, &githubResponse)
	if errUnpack != nil {
		return nil, errUnpack
	}
	return githubResponse, nil

}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}
