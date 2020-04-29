package models

import(
    "github.com/eifandevs/amby/repo"
    "github.com/thoas/go-funk"
	"github.com/jinzhu/gorm"
)

type AccessTokenItem struct {
    Token string
    Expire string
}

type AccessToken struct {
    gorm.Model
    Token string
    Expire string
}

type GetAccessTokenResponse struct {
    BaseResponse
    Items  []AccessTokenItem `json:"data"`
}

func GetAccessToken() GetAccessTokenResponse {
    db := repo.Connect("development")
    defer db.Close()

    accessTokens := []AccessToken{}
    if err := db.Find(&accessTokens).Error; err != nil {
        return GetAccessTokenResponse{BaseResponse: BaseResponse{Result: "NG", ErrorCode: ""}, Items: nil}
    }

    items := funk.Map(accessTokens, func(accessToken AccessToken) AccessTokenItem {
        return AccessTokenItem{Token: accessToken.Token, Expire: accessToken.Expire}
    })
    
    if castedItems, ok := items.([]AccessTokenItem); ok {
        return GetAccessTokenResponse{BaseResponse: BaseResponse{Result: "OK", ErrorCode: ""}, Items: castedItems}
    } else {
        panic("cannot cast accessToken item.")
    }
}