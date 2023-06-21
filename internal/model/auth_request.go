package model

// Параметры запроса POST /oauth2/access_token
type AuthRequest struct {
	ClientID     string `json:"client_id,omitempty"`     // ID Интеграции
	ClientSecret string `json:"client_secret,omitempty"` // Секрет интеграции
	GrantType    string `json:"grant_type,omitempty"`    // refresh_token или authorization_code
	Code         string `json:"code,omitempty"`          // Код авторизации(используется 1 раз grant_type=authorization_code)
	RefreshToken string `json:"refresh_token,omitempty"` // Токен обновления (используется для grant_type=refresh_token)
	RedirectURI  string `json:"redirect_uri,omitempty"`  // URI указанный в настройках интеграции
}
