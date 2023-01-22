package presenter

import "github.com/sparkymat/fundock/model"

//nolint:tagliatelle
type APIToken struct {
	ID           string  `json:"id"`
	ClientName   string  `json:"client_name"`
	Token        string  `json:"token"`
	LastUsedTime *string `json:"last_used_time"`
}

//nolint:tagliatelle
type APITokensList struct {
	PageNumber uint32     `json:"page_number"`
	PageSize   uint32     `json:"page_size"`
	Items      []APIToken `json:"items"`
}

func APITokenFromModel(tk model.APIToken) APIToken {
	presentedToken := APIToken{
		ID:         tk.ID,
		ClientName: tk.ClientName,
		Token:      tk.Token,
	}

	if tk.LastUsedAt != nil {
		lastUsedTime := tk.LastUsedAt.String()
		presentedToken.LastUsedTime = &lastUsedTime
	}

	return presentedToken
}

func APITokensListFromModels(pageNumber uint32, pageSize uint32, apiTokens []model.APIToken) APITokensList {
	presentedList := APITokensList{
		PageNumber: pageNumber,
		PageSize:   pageSize,
		Items:      []APIToken{},
	}

	for _, tk := range apiTokens {
		presentedList.Items = append(presentedList.Items, APITokenFromModel(tk))
	}

	return presentedList
}
