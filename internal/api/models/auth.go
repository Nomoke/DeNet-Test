package models

type AuthTokensResp struct {
	AccessToken  string `json:"accessToken"`
}

type AuthUserReq struct {
	Nickname string `json:"nickname" validate:"required,gte=1"`
	Password string `json:"password" validate:"required,gte=1"`
}