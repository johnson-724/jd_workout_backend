package google

import (
	"context"
	"encoding/json"
	"io"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	oauthApi "google.golang.org/api/oauth2/v2"
)

type UserInfo struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	Pricture string `json:"picture"`
}

func CreateGoogleOAuthURL() string {
	tokFile := "credential/google/credentials.json"
	path := getAbsPath(tokFile)

	data, err := os.ReadFile(path)

	if err != nil {
		panic(err)
	}

	scopes := []string{
		oauthApi.UserinfoEmailScope,
		oauthApi.UserinfoProfileScope,
	}

	config, err := google.ConfigFromJSON(data, scopes...)

	if err != nil {
		panic(err)
	}

	return config.AuthCodeURL("state")
}

func GetAccessToken(code string) (*oauth2.Token, error) {
	config, _ := getConfig()

	token, err := config.Exchange(context.TODO(), code)

	if err != nil {
		return nil, err
	}

	return token, nil
}

func GetUserInfo(token *oauth2.Token) (*UserInfo, error) {
	config, _ := getConfig()

	client := config.Client(context.Background(), token)

	res, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")

	if err != nil {
		return nil, err
	}

	content, _ := io.ReadAll(res.Body)

	defer res.Body.Close()

	userInfo := &UserInfo{}

	if err := json.Unmarshal(content, &userInfo); err != nil {
		return nil, err
	}

	return userInfo, nil
}

func getConfig() (*oauth2.Config, error) {
	tokFile := "credential/google/credentials.json"
	path := getAbsPath(tokFile)

	data, err := os.ReadFile(path)

	if err != nil {
		return nil, err
	}

	scopes := []string{
		oauthApi.UserinfoEmailScope,
		oauthApi.UserinfoProfileScope,
	}

	config, err := google.ConfigFromJSON(data, scopes...)

	if err != nil {
		return nil, err
	}

	return config, nil
}
