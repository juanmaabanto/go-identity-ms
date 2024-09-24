package model

type ClientApp struct {
	ID             string   `json:"_key,omitempty"`
	Name           string   `json:"name,omitempty"`
	Description    string   `json:"description,omitempty"`
	ClientId       string   `json:"clientId,omitempty"`
	ClientSecret   string   `json:"clientSecret,omitempty"`
	RedirectUris   []string `json:"redirectUris,omitempty"`
	AllowedOrigins []string `json:"allowedOrigins,omitempty"`
	ThirdParty     bool     `json:"thirdParty,omitempty"`
	Active         bool     `json:"active,omitempty"`
	CreatedAt      int64    `json:"createdAt,omitempty"`
	CreatedBy      string   `json:"createdBy,omitempty"`
	ModifiedAt     *int64   `json:"modifiedAt,omitempty"`
	ModifiedBy     string   `json:"modifiedBy,omitempty"`
}

func (ClientApp) GetCollectionName() string {
	return "ClientApps"
}
