package fortnite

import (
	"fmt"
	"net/url"
)

//OAuth URLs
var oauthTokenEndpoint = "https://account-public-service-prod03.ol.epicgames.com/account/api/oauth/token"
var oauthExchangeEndpoint = "https://account-public-service-prod03.ol.epicgames.com/account/api/oauth/exchange"
var oauthVerifyEndpoint = "https://account-public-service-prod03.ol.epicgames.com/account/api/oauth/verify?includePerms=true"

//Require Authentication
var fortniteStatusEndpoint = "https://lightswitch-public-service-prod06.ol.epicgames.com/lightswitch/api/service/bulk/status?serviceId=Fortnite"
var fortniteNewsEndpoint = "https://fortnitecontent-website-prod07.ol.epicgames.com/content/api/pages/fortnite-game"

//Do Not Require Authentication
var fortnitePVEInfoEndpoint = "https://fortnite-public-service-prod11.ol.epicgames.com/fortnite/api/game/v2/world/info"
var fortniteStoreEndpoint = "https://fortnite-public-service-prod11.ol.epicgames.com/fortnite/api/storefront/v2/catalog"

func lookupURLEndpoint(username string) string {
	return fmt.Sprintf("https://persona-public-service-prod06.ol.epicgames.com/persona/api/public/account/lookup?q=%v", url.QueryEscape(username))
}

func statsBattleRoyaleEndpoint(accountID string) string {
	return fmt.Sprintf("https://fortnite-public-service-prod11.ol.epicgames.com/fortnite/api/stats/accountId/%v/bulk/window/alltime", accountID)
}

func statsPVEEndpoint(accountID string) string {
	return fmt.Sprintf("https://fortnite-public-service-prod11.ol.epicgames.com/fortnite/api/game/v2/profile/%v/public/QueryProfile?profileId=profile0&rvn=-1", accountID)
}

func killSessionEndpoint(token string) string {
	return fmt.Sprintf("https://account-public-service-prod03.ol.epicgames.com/account/api/oauth/sessions/kill/%v", token)
}
