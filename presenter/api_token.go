package presenter

import "github.com/sparkymat/fundock/model"

type APIToken struct {
	ClientName string
	Token      string
	LastUsed   *string
}

func APITokenFromModel(apiToken model.APIToken) APIToken {
	presentedToken := APIToken{
		ClientName: apiToken.ClientName,
		Token:      apiToken.Token,
	}

	if apiToken.LastUsedAt != nil {
		lastUsedAt := apiToken.LastUsedAt.String()
		presentedToken.LastUsed = &lastUsedAt
	}

	return presentedToken
}
