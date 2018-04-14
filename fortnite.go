package fortnite

import (
	"fmt"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/parnurzeal/gorequest"
)

//OauthTokenRequest holds the required fields to initiate the authentication process
type OauthTokenRequest struct {
	GrantType    string `json:"grant_type"`
	Username     string `json:"username"`
	Password     string `json:"password"`
	IncludePerms bool   `json:"includePerms"`
}

//OauthTokenRequestResponse is used to unmarshal the JSON response received after initiating
//the authentication process. It gives us access to the access token required to continue the
//authentication flow.
type OauthTokenRequestResponse struct {
	AccessToken string `json:"access_token"`
}

//OauthRequestCodeResponse is used to unmarshal the JSON response received after using the
//access token to request an authentication code.
type OauthRequestCodeResponse struct {
	ExchangeCode string `json:"code"`
}

//OauthTokenExchangeRequest is used to marshal a JSON payload to finalise the authentication
//process.
type OauthTokenExchangeRequest struct {
	GrantType    string `json:"grant_type"`
	ExchangeCode string `json:"exchange_code"`
	IncludePerms bool   `json:"includePerms"`
	TokenType    string `json:"token_type"`
}

//OauthTokenResponse is used to unmarshal the JSON response received after successfully
//completing the authentication process.
type OauthTokenResponse struct {
	ExpiresAt    string `json:"expires_at"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

//User represents the state of the user retrieved from the Fortnite API
type User struct {
	ID          string `json:"id"`
	DisplayName string `json:"displayName"`
}

//StatusResponse is used the unmarshal the JSON response received after successfully
//querying the status endpoint of their API.
type StatusResponse []struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

//NewsResponse is used the unmarshal the JSON response received after successfully
//querying the news endpoint of their API.
type NewsResponse struct {
	Survivalmessage struct {
		Overrideablemessage struct {
			Message struct {
				Type  string `json:"_type"`
				Title string `json:"title"`
				Body  string `json:"body"`
			} `json:"message"`
			Messages []struct {
				Image  string `json:"image"`
				Hidden bool   `json:"hidden"`
				Type   string `json:"_type"`
				Title  string `json:"title"`
				Body   string `json:"body"`
			} `json:"messages"`
		} `json:"overrideablemessage"`
	} `json:"survivalmessage"`
	Athenamessage struct {
		Overrideablemessage struct {
			Type    string `json:"_type"`
			Message struct {
				Image string `json:"image"`
				Type  string `json:"_type"`
				Title string `json:"title"`
				Body  string `json:"body"`
			} `json:"message"`
			Messages []struct {
				Image  string `json:"image"`
				Hidden bool   `json:"hidden"`
				Type   string `json:"_type"`
				Title  string `json:"title"`
				Body   string `json:"body"`
			} `json:"messages"`
		} `json:"overrideablemessage"`
	} `json:"athenamessage"`
	Savetheworldnews struct {
		News struct {
			Message struct {
				Image string `json:"image"`
				Type  string `json:"_type"`
				Title string `json:"title"`
				Body  string `json:"body"`
			} `json:"message"`
			Messages []struct {
				Image  string `json:"image"`
				Hidden bool   `json:"hidden"`
				Type   string `json:"_type"`
				Title  string `json:"title"`
				Body   string `json:"body"`
			} `json:"messages"`
		} `json:"news"`
	} `json:"savetheworldnews"`
	Battleroyalenews struct {
		News struct {
			Message struct {
				Image string `json:"image"`
				Type  string `json:"_type"`
				Title string `json:"title"`
				Body  string `json:"body"`
			} `json:"message"`
			Messages []struct {
				Image  string `json:"image"`
				Hidden bool   `json:"hidden"`
				Type   string `json:"_type"`
				Title  string `json:"title"`
				Body   string `json:"body"`
			} `json:"messages"`
		} `json:"news"`
	} `json:"battleroyalenews"`
	Loginmessage struct {
		Loginmessage struct {
			Message struct {
				Type  string `json:"_type"`
				Title string `json:"title"`
				Body  string `json:"body"`
			} `json:"message"`
			Messages []struct {
				Image  string `json:"image"`
				Hidden bool   `json:"hidden"`
				Type   string `json:"_type"`
				Title  string `json:"title"`
				Body   string `json:"body"`
			} `json:"messages"`
		} `json:"loginmessage"`
	} `json:"loginmessage"`
}

//StoreResponse is used the unmarshal the JSON response received after successfully
//querying the store endpoint of their API.
type StoreResponse struct {
	RefreshIntervalHrs int       `json:"refreshIntervalHrs"`
	DailyPurchaseHrs   int       `json:"dailyPurchaseHrs"`
	Expiration         time.Time `json:"expiration"`
	Storefronts        []struct {
		Name           string `json:"name"`
		CatalogEntries []struct {
			OfferID   string `json:"offerId"`
			DevName   string `json:"devName"`
			OfferType string `json:"offerType"`
			Prices    []struct {
				CurrencyType    string    `json:"currencyType"`
				CurrencySubType string    `json:"currencySubType"`
				RegularPrice    int       `json:"regularPrice"`
				FinalPrice      int       `json:"finalPrice"`
				SaleExpiration  time.Time `json:"saleExpiration"`
				BasePrice       int       `json:"basePrice"`
			} `json:"prices"`
			Categories   []interface{} `json:"categories"`
			DailyLimit   int           `json:"dailyLimit"`
			WeeklyLimit  int           `json:"weeklyLimit"`
			MonthlyLimit int           `json:"monthlyLimit"`
			AppStoreID   []string      `json:"appStoreId"`
			Requirements []interface{} `json:"requirements"`
			MetaInfo     []struct {
				Key   string `json:"key"`
				Value string `json:"value"`
			} `json:"metaInfo"`
			CatalogGroup         string        `json:"catalogGroup"`
			CatalogGroupPriority int           `json:"catalogGroupPriority"`
			SortPriority         int           `json:"sortPriority"`
			Title                string        `json:"title"`
			ShortDescription     string        `json:"shortDescription"`
			Description          string        `json:"description"`
			DisplayAssetPath     string        `json:"displayAssetPath"`
			ItemGrants           []interface{} `json:"itemGrants"`
			GiftInfo             struct {
				BIsEnabled              bool          `json:"bIsEnabled"`
				ForcedGiftBoxTemplateID string        `json:"forcedGiftBoxTemplateId"`
				PurchaseRequirements    []interface{} `json:"purchaseRequirements"`
			} `json:"giftInfo,omitempty"`
		} `json:"catalogEntries"`
	} `json:"storefronts"`
}

//PveInfoResponse is used the unmarshal the JSON response received after successfully
//querying the PVE info endpoint of their API.
type PveInfoResponse struct {
	Theaters []struct {
		DisplayName                      string `json:"displayName"`
		UniqueID                         string `json:"uniqueId"`
		TheaterSlot                      int    `json:"theaterSlot"`
		BIsTestTheater                   bool   `json:"bIsTestTheater"`
		BHideLikeTestTheater             bool   `json:"bHideLikeTestTheater"`
		RequiredEventFlag                string `json:"requiredEventFlag"`
		MissionRewardNamedWeightsRowName string `json:"missionRewardNamedWeightsRowName"`
		Description                      string `json:"description"`
		RuntimeInfo                      struct {
			TheaterType string `json:"theaterType"`
			TheaterTags struct {
				GameplayTags []struct {
					TagName string `json:"tagName"`
				} `json:"gameplayTags"`
			} `json:"theaterTags"`
			TheaterVisibilityRequirements struct {
				CommanderLevel         int    `json:"commanderLevel"`
				PersonalPowerRating    int    `json:"personalPowerRating"`
				MaxPersonalPowerRating int    `json:"maxPersonalPowerRating"`
				PartyPowerRating       int    `json:"partyPowerRating"`
				MaxPartyPowerRating    int    `json:"maxPartyPowerRating"`
				ActiveQuestDefinition  string `json:"activeQuestDefinition"`
				QuestDefinition        string `json:"questDefinition"`
				ObjectiveStatHandle    struct {
					DataTable string `json:"dataTable"`
					RowName   string `json:"rowName"`
				} `json:"objectiveStatHandle"`
				UncompletedQuestDefinition string `json:"uncompletedQuestDefinition"`
				ItemDefinition             string `json:"itemDefinition"`
			} `json:"theaterVisibilityRequirements"`
			Requirements struct {
				CommanderLevel         int    `json:"commanderLevel"`
				PersonalPowerRating    int    `json:"personalPowerRating"`
				MaxPersonalPowerRating int    `json:"maxPersonalPowerRating"`
				PartyPowerRating       int    `json:"partyPowerRating"`
				MaxPartyPowerRating    int    `json:"maxPartyPowerRating"`
				ActiveQuestDefinition  string `json:"activeQuestDefinition"`
				QuestDefinition        string `json:"questDefinition"`
				ObjectiveStatHandle    struct {
					DataTable string `json:"dataTable"`
					RowName   string `json:"rowName"`
				} `json:"objectiveStatHandle"`
				UncompletedQuestDefinition string `json:"uncompletedQuestDefinition"`
				ItemDefinition             string `json:"itemDefinition"`
			} `json:"requirements"`
			RequiredSubGameForVisibility  string `json:"requiredSubGameForVisibility"`
			BOnlyMatchLinkedQuestsToTiles bool   `json:"bOnlyMatchLinkedQuestsToTiles"`
			WorldMapPinClass              string `json:"worldMapPinClass"`
			TheaterImage                  string `json:"theaterImage"`
			TheaterImages                 struct {
				BrushXXS struct {
					ImageSize struct {
						X int `json:"x"`
						Y int `json:"y"`
					} `json:"imageSize"`
					Margin struct {
						Left   int `json:"left"`
						Top    int `json:"top"`
						Right  int `json:"right"`
						Bottom int `json:"bottom"`
					} `json:"margin"`
					TintColor struct {
						SpecifiedColor struct {
							R int `json:"r"`
							G int `json:"g"`
							B int `json:"b"`
							A int `json:"a"`
						} `json:"specifiedColor"`
						ColorUseRule string `json:"colorUseRule"`
					} `json:"tintColor"`
					ResourceObject string `json:"resourceObject"`
					ResourceName   string `json:"resourceName"`
					UVRegion       struct {
						Min struct {
							X int `json:"x"`
							Y int `json:"y"`
						} `json:"min"`
						Max struct {
							X int `json:"x"`
							Y int `json:"y"`
						} `json:"max"`
						BIsValid int `json:"bIsValid"`
					} `json:"uVRegion"`
					DrawAs               string `json:"drawAs"`
					Tiling               string `json:"tiling"`
					Mirroring            string `json:"mirroring"`
					ImageType            string `json:"imageType"`
					BIsDynamicallyLoaded bool   `json:"bIsDynamicallyLoaded"`
				} `json:"brush_XXS"`
				BrushXS struct {
					ImageSize struct {
						X int `json:"x"`
						Y int `json:"y"`
					} `json:"imageSize"`
					Margin struct {
						Left   int `json:"left"`
						Top    int `json:"top"`
						Right  int `json:"right"`
						Bottom int `json:"bottom"`
					} `json:"margin"`
					TintColor struct {
						SpecifiedColor struct {
							R int `json:"r"`
							G int `json:"g"`
							B int `json:"b"`
							A int `json:"a"`
						} `json:"specifiedColor"`
						ColorUseRule string `json:"colorUseRule"`
					} `json:"tintColor"`
					ResourceObject string `json:"resourceObject"`
					ResourceName   string `json:"resourceName"`
					UVRegion       struct {
						Min struct {
							X int `json:"x"`
							Y int `json:"y"`
						} `json:"min"`
						Max struct {
							X int `json:"x"`
							Y int `json:"y"`
						} `json:"max"`
						BIsValid int `json:"bIsValid"`
					} `json:"uVRegion"`
					DrawAs               string `json:"drawAs"`
					Tiling               string `json:"tiling"`
					Mirroring            string `json:"mirroring"`
					ImageType            string `json:"imageType"`
					BIsDynamicallyLoaded bool   `json:"bIsDynamicallyLoaded"`
				} `json:"brush_XS"`
				BrushS struct {
					ImageSize struct {
						X int `json:"x"`
						Y int `json:"y"`
					} `json:"imageSize"`
					Margin struct {
						Left   int `json:"left"`
						Top    int `json:"top"`
						Right  int `json:"right"`
						Bottom int `json:"bottom"`
					} `json:"margin"`
					TintColor struct {
						SpecifiedColor struct {
							R int `json:"r"`
							G int `json:"g"`
							B int `json:"b"`
							A int `json:"a"`
						} `json:"specifiedColor"`
						ColorUseRule string `json:"colorUseRule"`
					} `json:"tintColor"`
					ResourceObject string `json:"resourceObject"`
					ResourceName   string `json:"resourceName"`
					UVRegion       struct {
						Min struct {
							X int `json:"x"`
							Y int `json:"y"`
						} `json:"min"`
						Max struct {
							X int `json:"x"`
							Y int `json:"y"`
						} `json:"max"`
						BIsValid int `json:"bIsValid"`
					} `json:"uVRegion"`
					DrawAs               string `json:"drawAs"`
					Tiling               string `json:"tiling"`
					Mirroring            string `json:"mirroring"`
					ImageType            string `json:"imageType"`
					BIsDynamicallyLoaded bool   `json:"bIsDynamicallyLoaded"`
				} `json:"brush_S"`
				BrushM struct {
					ImageSize struct {
						X int `json:"x"`
						Y int `json:"y"`
					} `json:"imageSize"`
					Margin struct {
						Left   int `json:"left"`
						Top    int `json:"top"`
						Right  int `json:"right"`
						Bottom int `json:"bottom"`
					} `json:"margin"`
					TintColor struct {
						SpecifiedColor struct {
							R int `json:"r"`
							G int `json:"g"`
							B int `json:"b"`
							A int `json:"a"`
						} `json:"specifiedColor"`
						ColorUseRule string `json:"colorUseRule"`
					} `json:"tintColor"`
					ResourceObject string `json:"resourceObject"`
					ResourceName   string `json:"resourceName"`
					UVRegion       struct {
						Min struct {
							X int `json:"x"`
							Y int `json:"y"`
						} `json:"min"`
						Max struct {
							X int `json:"x"`
							Y int `json:"y"`
						} `json:"max"`
						BIsValid int `json:"bIsValid"`
					} `json:"uVRegion"`
					DrawAs               string `json:"drawAs"`
					Tiling               string `json:"tiling"`
					Mirroring            string `json:"mirroring"`
					ImageType            string `json:"imageType"`
					BIsDynamicallyLoaded bool   `json:"bIsDynamicallyLoaded"`
				} `json:"brush_M"`
				BrushL struct {
					ImageSize struct {
						X int `json:"x"`
						Y int `json:"y"`
					} `json:"imageSize"`
					Margin struct {
						Left   int `json:"left"`
						Top    int `json:"top"`
						Right  int `json:"right"`
						Bottom int `json:"bottom"`
					} `json:"margin"`
					TintColor struct {
						SpecifiedColor struct {
							R int `json:"r"`
							G int `json:"g"`
							B int `json:"b"`
							A int `json:"a"`
						} `json:"specifiedColor"`
						ColorUseRule string `json:"colorUseRule"`
					} `json:"tintColor"`
					ResourceObject string `json:"resourceObject"`
					ResourceName   string `json:"resourceName"`
					UVRegion       struct {
						Min struct {
							X int `json:"x"`
							Y int `json:"y"`
						} `json:"min"`
						Max struct {
							X int `json:"x"`
							Y int `json:"y"`
						} `json:"max"`
						BIsValid int `json:"bIsValid"`
					} `json:"uVRegion"`
					DrawAs               string `json:"drawAs"`
					Tiling               string `json:"tiling"`
					Mirroring            string `json:"mirroring"`
					ImageType            string `json:"imageType"`
					BIsDynamicallyLoaded bool   `json:"bIsDynamicallyLoaded"`
				} `json:"brush_L"`
				BrushXL struct {
					ImageSize struct {
						X int `json:"x"`
						Y int `json:"y"`
					} `json:"imageSize"`
					Margin struct {
						Left   int `json:"left"`
						Top    int `json:"top"`
						Right  int `json:"right"`
						Bottom int `json:"bottom"`
					} `json:"margin"`
					TintColor struct {
						SpecifiedColor struct {
							R int `json:"r"`
							G int `json:"g"`
							B int `json:"b"`
							A int `json:"a"`
						} `json:"specifiedColor"`
						ColorUseRule string `json:"colorUseRule"`
					} `json:"tintColor"`
					ResourceObject string `json:"resourceObject"`
					ResourceName   string `json:"resourceName"`
					UVRegion       struct {
						Min struct {
							X int `json:"x"`
							Y int `json:"y"`
						} `json:"min"`
						Max struct {
							X int `json:"x"`
							Y int `json:"y"`
						} `json:"max"`
						BIsValid int `json:"bIsValid"`
					} `json:"uVRegion"`
					DrawAs               string `json:"drawAs"`
					Tiling               string `json:"tiling"`
					Mirroring            string `json:"mirroring"`
					ImageType            string `json:"imageType"`
					BIsDynamicallyLoaded bool   `json:"bIsDynamicallyLoaded"`
				} `json:"brush_XL"`
			} `json:"theaterImages"`
			TheaterColorInfo struct {
				BUseDifficultyToDetermineColor bool `json:"bUseDifficultyToDetermineColor"`
				Color                          struct {
					SpecifiedColor struct {
						R float64 `json:"r"`
						G int     `json:"g"`
						B float64 `json:"b"`
						A int     `json:"a"`
					} `json:"specifiedColor"`
					ColorUseRule string `json:"colorUseRule"`
				} `json:"color"`
			} `json:"theaterColorInfo"`
			Socket                   string `json:"socket"`
			MissionAlertRequirements struct {
				CommanderLevel         int    `json:"commanderLevel"`
				PersonalPowerRating    int    `json:"personalPowerRating"`
				MaxPersonalPowerRating int    `json:"maxPersonalPowerRating"`
				PartyPowerRating       int    `json:"partyPowerRating"`
				MaxPartyPowerRating    int    `json:"maxPartyPowerRating"`
				ActiveQuestDefinition  string `json:"activeQuestDefinition"`
				QuestDefinition        string `json:"questDefinition"`
				ObjectiveStatHandle    struct {
					DataTable string `json:"dataTable"`
					RowName   string `json:"rowName"`
				} `json:"objectiveStatHandle"`
				UncompletedQuestDefinition string `json:"uncompletedQuestDefinition"`
				ItemDefinition             string `json:"itemDefinition"`
			} `json:"missionAlertRequirements"`
			MissionAlertCategoryRequirements []struct {
				MissionAlertCategoryName string `json:"missionAlertCategoryName"`
				BRespectTileRequirements bool   `json:"bRespectTileRequirements"`
				BAllowQuickplay          bool   `json:"bAllowQuickplay"`
			} `json:"missionAlertCategoryRequirements"`
		} `json:"runtimeInfo"`
		Tiles []struct {
			TileType     string `json:"tileType"`
			ZoneTheme    string `json:"zoneTheme"`
			Requirements struct {
				CommanderLevel         int    `json:"commanderLevel"`
				PersonalPowerRating    int    `json:"personalPowerRating"`
				MaxPersonalPowerRating int    `json:"maxPersonalPowerRating"`
				PartyPowerRating       int    `json:"partyPowerRating"`
				MaxPartyPowerRating    int    `json:"maxPartyPowerRating"`
				ActiveQuestDefinition  string `json:"activeQuestDefinition"`
				QuestDefinition        string `json:"questDefinition"`
				ObjectiveStatHandle    struct {
					DataTable string `json:"dataTable"`
					RowName   string `json:"rowName"`
				} `json:"objectiveStatHandle"`
				UncompletedQuestDefinition string `json:"uncompletedQuestDefinition"`
				ItemDefinition             string `json:"itemDefinition"`
			} `json:"requirements"`
			LinkedQuests           []interface{} `json:"linkedQuests"`
			XCoordinate            int           `json:"xCoordinate"`
			YCoordinate            int           `json:"yCoordinate"`
			MissionWeightOverrides []struct {
				Weight           float64 `json:"weight"`
				MissionGenerator string  `json:"missionGenerator"`
			} `json:"missionWeightOverrides"`
			DifficultyWeightOverrides []interface{} `json:"difficultyWeightOverrides"`
			CanBeMissionAlert         bool          `json:"canBeMissionAlert"`
			TileTags                  struct {
				GameplayTags []interface{} `json:"gameplayTags"`
			} `json:"tileTags"`
		} `json:"tiles"`
		Regions []struct {
			DisplayName string `json:"displayName"`
			RegionTags  struct {
				GameplayTags []struct {
					TagName string `json:"tagName"`
				} `json:"gameplayTags"`
			} `json:"regionTags"`
			TileIndices     []int  `json:"tileIndices"`
			RegionThemeIcon string `json:"regionThemeIcon"`
			MissionData     struct {
				MissionWeights []struct {
					Weight           float64 `json:"weight"`
					MissionGenerator string  `json:"missionGenerator"`
				} `json:"missionWeights"`
				DifficultyWeights []struct {
					Weight         float64 `json:"weight"`
					DifficultyInfo struct {
						DataTable string `json:"dataTable"`
						RowName   string `json:"rowName"`
					} `json:"difficultyInfo"`
				} `json:"difficultyWeights"`
				NumMissionsAvailable   int     `json:"numMissionsAvailable"`
				NumMissionsToChange    int     `json:"numMissionsToChange"`
				MissionChangeFrequency float64 `json:"missionChangeFrequency"`
			} `json:"missionData"`
			Requirements struct {
				CommanderLevel         int    `json:"commanderLevel"`
				PersonalPowerRating    int    `json:"personalPowerRating"`
				MaxPersonalPowerRating int    `json:"maxPersonalPowerRating"`
				PartyPowerRating       int    `json:"partyPowerRating"`
				MaxPartyPowerRating    int    `json:"maxPartyPowerRating"`
				ActiveQuestDefinition  string `json:"activeQuestDefinition"`
				QuestDefinition        string `json:"questDefinition"`
				ObjectiveStatHandle    struct {
					DataTable string `json:"dataTable"`
					RowName   string `json:"rowName"`
				} `json:"objectiveStatHandle"`
				UncompletedQuestDefinition string `json:"uncompletedQuestDefinition"`
				ItemDefinition             string `json:"itemDefinition"`
			} `json:"requirements"`
			MissionAlertRequirements []struct {
				CategoryName string `json:"categoryName"`
				Requirements struct {
					CommanderLevel         int    `json:"commanderLevel"`
					PersonalPowerRating    int    `json:"personalPowerRating"`
					MaxPersonalPowerRating int    `json:"maxPersonalPowerRating"`
					PartyPowerRating       int    `json:"partyPowerRating"`
					MaxPartyPowerRating    int    `json:"maxPartyPowerRating"`
					ActiveQuestDefinition  string `json:"activeQuestDefinition"`
					QuestDefinition        string `json:"questDefinition"`
					ObjectiveStatHandle    struct {
						DataTable string `json:"dataTable"`
						RowName   string `json:"rowName"`
					} `json:"objectiveStatHandle"`
					UncompletedQuestDefinition string `json:"uncompletedQuestDefinition"`
					ItemDefinition             string `json:"itemDefinition"`
				} `json:"requirements"`
			} `json:"missionAlertRequirements"`
		} `json:"regions"`
	} `json:"theaters"`
	Missions []struct {
		TheaterID         string `json:"theaterId"`
		AvailableMissions []struct {
			MissionGUID    string `json:"missionGuid"`
			MissionRewards struct {
				TierGroupName string `json:"tierGroupName"`
				Items         []struct {
					ItemType string `json:"itemType"`
					Quantity int    `json:"quantity"`
				} `json:"items"`
			} `json:"missionRewards"`
			MissionGenerator      string `json:"missionGenerator"`
			MissionDifficultyInfo struct {
				DataTable string `json:"dataTable"`
				RowName   string `json:"rowName"`
			} `json:"missionDifficultyInfo"`
			TileIndex           int       `json:"tileIndex"`
			AvailableUntil      time.Time `json:"availableUntil"`
			BonusMissionRewards struct {
				TierGroupName string `json:"tierGroupName"`
				Items         []struct {
					ItemType string `json:"itemType"`
					Quantity int    `json:"quantity"`
				} `json:"items"`
			} `json:"bonusMissionRewards,omitempty"`
		} `json:"availableMissions"`
		NextRefresh time.Time `json:"nextRefresh"`
	} `json:"missions"`
	MissionAlerts []struct {
		TheaterID              string `json:"theaterId"`
		AvailableMissionAlerts []struct {
			Name                 string    `json:"name"`
			CategoryName         string    `json:"categoryName"`
			SpreadDataName       string    `json:"spreadDataName"`
			MissionAlertGUID     string    `json:"missionAlertGuid"`
			TileIndex            int       `json:"tileIndex"`
			AvailableUntil       time.Time `json:"availableUntil"`
			TotalSpreadRefreshes int       `json:"totalSpreadRefreshes"`
			MissionAlertRewards  struct {
				TierGroupName string `json:"tierGroupName"`
				Items         []struct {
					ItemType string `json:"itemType"`
					Quantity int    `json:"quantity"`
				} `json:"items"`
			} `json:"missionAlertRewards"`
			MissionAlertModifiers struct {
				TierGroupName string `json:"tierGroupName"`
				Items         []struct {
					ItemType string `json:"itemType"`
					Quantity int    `json:"quantity"`
				} `json:"items"`
			} `json:"missionAlertModifiers"`
		} `json:"availableMissionAlerts"`
		NextRefresh time.Time `json:"nextRefresh"`
	} `json:"missionAlerts"`
}

//RawBRStatsResponse is used the unmarshal the JSON response received after successfully
//querying the leaderboard endpoint of their API.
type RawBRStatsResponse []struct {
	Name      string  `json:"name"`
	Value     float64 `json:"value"`
	Window    string  `json:"window"`
	OwnerType int     `json:"ownerType"`
}

//FormattedBRStats is used to store the BR stats from a RawBRStatsResponse after they have been
//transformed into a more readable and meaningful state
type FormattedBRStats struct {
	Group struct {
		Solo struct {
			Wins                float64 `json:"wins"`
			Top3                float64 `json:"top3"`
			Top5                float64 `json:"top5"`
			Top6                float64 `json:"top6"`
			Top10               float64 `json:"top10"`
			Top12               float64 `json:"top12"`
			Top25               float64 `json:"top25"`
			KdRatio             float64 `json:"kd_ratio"`
			WinPercentage       float64 `json:"win_percentage"`
			Matches             float64 `json:"matches"`
			Kills               float64 `json:"kills"`
			TimePlayed          float64 `json:"time_played"`
			TimePlayedFormatted string  `json:"time_played_formatted"`
			KillsPerMatch       float64 `json:"kills_per_match"`
			KillsPerMin         float64 `json:"kills_per_min"`
			Score               float64 `json:"score"`
		}
		Duo struct {
			Wins                float64 `json:"wins"`
			Top3                float64 `json:"top3"`
			Top5                float64 `json:"top5"`
			Top6                float64 `json:"top6"`
			Top10               float64 `json:"top10"`
			Top12               float64 `json:"top12"`
			Top25               float64 `json:"top25"`
			KdRatio             float64 `json:"kd_ratio"`
			WinPercentage       float64 `json:"win_percentage"`
			Matches             float64 `json:"matches"`
			Kills               float64 `json:"kills"`
			TimePlayed          float64 `json:"time_played"`
			TimePlayedFormatted string  `json:"time_played_formatted"`
			KillsPerMatch       float64 `json:"kills_per_match"`
			KillsPerMin         float64 `json:"kills_per_min"`
			Score               float64 `json:"score"`
		}
		Squad struct {
			Wins                float64 `json:"wins"`
			Top3                float64 `json:"top3"`
			Top5                float64 `json:"top5"`
			Top6                float64 `json:"top6"`
			Top10               float64 `json:"top10"`
			Top12               float64 `json:"top12"`
			Top25               float64 `json:"top25"`
			KdRatio             float64 `json:"kd_ratio"`
			WinPercentage       float64 `json:"win_percentage"`
			Matches             float64 `json:"matches"`
			Kills               float64 `json:"kills"`
			TimePlayed          float64 `json:"time_played"`
			TimePlayedFormatted string  `json:"time_played_formatted"`
			KillsPerMatch       float64 `json:"kills_per_match"`
			KillsPerMin         float64 `json:"kills_per_min"`
			Score               float64 `json:"score"`
		}
	}
	Info struct {
		AccountID string `json:"account_id"`
		Username  string `json:"username"`
		Platform  string `json:"platform"`
	}
	LifetimeStats struct {
		Wins                float64 `json:"wins"`
		Top3                float64 `json:"top3"`
		Top5                float64 `json:"top5"`
		Top6                float64 `json:"top6"`
		Top10               float64 `json:"top10"`
		Top12               float64 `json:"top12"`
		Top25               float64 `json:"top25"`
		KdRatio             float64 `json:"kd_ratio"`
		WinPercentage       float64 `json:"win_percentage"`
		Matches             float64 `json:"matches"`
		Kills               float64 `json:"kills"`
		TimePlayed          float64 `json:"time_played"`
		TimePlayedFormatted string  `json:"time_played_formatted"`
		KillsPerMatch       float64 `json:"kills_per_match"`
		KillsPerMin         float64 `json:"kills_per_min"`
		Score               float64 `json:"score"`
	}
}

//Client represents the Fortnite Client and is used as the access point to query any of the API
//endpoints.
type Client struct {
	Email                string
	Password             string
	ClientLauncherToken  string
	FortniteClientToken  string
	AccessToken          string
	AccessTokenExpiresAt time.Time
	RefreshToken         string
	Request              *gorequest.SuperAgent
	Mutex                sync.Mutex
}

//NewClient instantiates an instance of Client that can then be used to make queries to the Fortnite
//API.
func NewClient(email string, password string, clientLauncherToken string, fortniteClientToken string) *Client {
	c := &Client{
		Email:               email,
		Password:            password,
		ClientLauncherToken: clientLauncherToken,
		FortniteClientToken: fortniteClientToken,
		Request:             gorequest.New(),
	}

	return c
}

//Login completes the OAuth authentication process, which is required to make calls to the Fortnite API
func (c *Client) Login() {
	tokenConfig := OauthTokenRequest{
		GrantType:    "password",
		Username:     c.Email,
		Password:     c.Password,
		IncludePerms: true,
	}

	var accessTokenResponse OauthTokenRequestResponse

	_, _, err := c.Request.Post(oauthTokenEndpoint).
		SendStruct(tokenConfig).
		Type("form").
		Set("Authorization", fmt.Sprintf("basic %v", c.FortniteClientToken)).
		EndStruct(&accessTokenResponse)

	if err != nil {
		log.Fatal("Error with login request 1:", err)
	}

	var codeResponse OauthRequestCodeResponse

	_, _, err = c.Request.Get(oauthExchangeEndpoint).
		Set("Authorization", fmt.Sprintf("bearer %v", accessTokenResponse.AccessToken)).
		EndStruct(&codeResponse)

	if err != nil {
		log.Fatal("Error with login request 2:", err)
	}

	exchangeRequest := OauthTokenExchangeRequest{
		GrantType:    "exchange_code",
		ExchangeCode: codeResponse.ExchangeCode,
		IncludePerms: true,
		TokenType:    "egl",
	}

	var tokenResponse OauthTokenResponse

	_, _, err = c.Request.Post(oauthTokenEndpoint).
		SendStruct(exchangeRequest).
		Type("form").
		Set("Authorization", fmt.Sprintf("basic %v", c.FortniteClientToken)).
		EndStruct(&tokenResponse)

	if err != nil {
		log.Fatal("Error with login request 3:", err)
	}

	c.Mutex.Lock()
	c.AccessTokenExpiresAt, _ = time.Parse(time.RFC3339Nano, tokenResponse.ExpiresAt)
	c.AccessToken = tokenResponse.AccessToken
	c.RefreshToken = tokenResponse.RefreshToken
	c.Mutex.Unlock()
}

//Lookup returns an instance of User which is the information received from the Fortnite API.
func (c *Client) Lookup(username string) User {
	var response User

	c.Mutex.Lock()
	_, _, err := c.Request.Get(lookupURLEndpoint(username)).
		Set("Authorization", fmt.Sprintf("bearer %v", c.AccessToken)).
		EndStruct(&response)
	c.Mutex.Unlock()

	if err != nil {
		log.Fatal(err)
	}

	return response
}

//CheckPlayer indicates whether a requested player exists and has played on the requested
//platform.
func (c *Client) CheckPlayer(username string, platform string) bool {

	if !(platform == "pc" || platform == "ps4" || platform == "xb1") {
		fmt.Println("Bad platform provided;", platform)
		return false
	}

	account := c.Lookup(username)

	if account.ID == "" {
		return false
	}

	var response RawBRStatsResponse

	c.Mutex.Lock()
	c.Request.Get(statsBattleRoyaleEndpoint(account.ID)).
		Set("Authorization", fmt.Sprintf("bearer %v", c.AccessToken)).
		EndStruct(&response)
	c.Mutex.Unlock()

	for _, stat := range response {
		if strings.Contains(stat.Name, fmt.Sprintf("_%v_", platform)) {
			return true
		}
	}

	return false
}

//GetStatsBR performs the necessary lookups and transformation to return a meaningful
//representation of your current Battle Royale stats.
//It will return the stats for the requested platform; useful if the player is active on
//more than one platform.
func (c *Client) GetStatsBR(username string, platform string) FormattedBRStats {

	if !(platform == "pc" || platform == "ps4" || platform == "xb1") {
		fmt.Println("Bad platform provided;", platform)
		return FormattedBRStats{}
	}

	account := c.Lookup(username)

	var response RawBRStatsResponse

	c.Mutex.Lock()
	c.Request.Get(statsBattleRoyaleEndpoint(account.ID)).
		Set("Authorization", fmt.Sprintf("bearer %v", c.AccessToken)).
		EndStruct(&response)
	c.Mutex.Unlock()

	return processBRStats(response, account, platform)
}

//GetStatsBRFromID is an alternative to GetStatsBR through which you can retrieve the stats
//for an account where you already know the Epic/Fortnite Account ID.
//It will return the stats for the requested platform; useful if the player is active on
//more than one platform.
func (c *Client) GetStatsBRFromID(accountID string, platform string) FormattedBRStats {

	if !(platform == "pc" || platform == "ps4" || platform == "xb1") {
		fmt.Println("Bad platform provided;", platform)
		return FormattedBRStats{}
	}

	account := User{
		ID:          accountID,
		DisplayName: "No Username",
	}

	var response RawBRStatsResponse

	c.Mutex.Lock()
	c.Request.Get(statsBattleRoyaleEndpoint(accountID)).
		Set("Authorization", fmt.Sprintf("bearer %v", c.AccessToken)).
		EndStruct(&response)
	c.Mutex.Unlock()

	return processBRStats(response, account, platform)
}

//GetFortniteNews returns a variety of news messages displayed in Fortnite.
//It includes news for Survival, STW, BR, and Login.
func (c *Client) GetFortniteNews(lang string) NewsResponse {
	languageHeader := ""

	switch lang {
	case "fr":
		languageHeader = "fr-FR"
	case "en":
		languageHeader = "en"
	default:
		languageHeader = "en"
	}

	var response NewsResponse

	c.Mutex.Lock()
	c.Request.Get(fortniteNewsEndpoint).
		Set("Authorization", fmt.Sprintf("bearer %v", c.AccessToken)).
		Set("Accept-Language", languageHeader).
		EndStruct(&response)
	c.Mutex.Unlock()

	return response
}

//CheckFortniteStatus checks their status endpoint and will return a bool to
//indicate whether Fortnite is up or not and if not then the message Fortnite
//have provided for why it is down.
func (c *Client) CheckFortniteStatus() (bool, string) {
	var response StatusResponse

	c.Mutex.Lock()
	c.Request.Get(fortniteStatusEndpoint).
		Set("Authorization", fmt.Sprintf("bearer %v", c.AccessToken)).
		EndStruct(&response)
	c.Mutex.Unlock()

	if len(response) > 0 {
		if response[0].Status == "UP" {
			return true, ""
		}

		return false, response[0].Message
	}

	return false, "No data returned from status endpoint"
}

//GetFortnitePVEInfo returns a variety of information specific to PVE.
func (c *Client) GetFortnitePVEInfo(lang string) PveInfoResponse {
	languageHeader := ""

	switch lang {
	case "fr":
		languageHeader = "fr-FR"
	case "en":
		languageHeader = "en"
	default:
		languageHeader = "en"
	}

	var response PveInfoResponse

	c.Mutex.Lock()
	c.Request.Get(fortnitePVEInfoEndpoint).
		Set("Authorization", fmt.Sprintf("bearer %v", c.AccessToken)).
		Set("X-EpicGames-Language", languageHeader).
		EndStruct(&response)
	c.Mutex.Unlock()

	return response
}

//GetStore returns all the items currently available for purchase for the user.
//This matches what you would see in the shop within Fortnite.
func (c *Client) GetStore(lang string) StoreResponse {
	languageHeader := ""

	switch lang {
	case "fr":
		languageHeader = "fr-FR"
	case "en":
		languageHeader = "en"
	default:
		languageHeader = "en"
	}

	var response StoreResponse

	c.Mutex.Lock()
	c.Request.Get(fortniteStoreEndpoint).
		Set("Authorization", fmt.Sprintf("bearer %v", c.AccessToken)).
		Set("X-EpicGames-Language", languageHeader).
		EndStruct(&response)
	c.Mutex.Unlock()

	return response
}

//KillSession is responsible for invalidating your OAuth Token.
func (c *Client) KillSession() {
	c.Mutex.Lock()
	c.Request.Delete(killSessionEndpoint(c.AccessToken)).
		Set("Authorization", fmt.Sprintf("bearer %v", c.AccessToken)).
		End()

	c.AccessToken = ""
	c.RefreshToken = ""
	c.AccessTokenExpiresAt = time.Now()
	c.Mutex.Unlock()
}

//CheckToken will check whether the current OAuth access token has expired, and
//if it has then it will refresh the token.
func (c *Client) CheckToken() {
	if c.AccessToken == "" {
		return
	}

	if time.Now().After(c.AccessTokenExpiresAt) {
		var response OauthTokenResponse

		_, _, err := c.Request.Post(oauthTokenEndpoint).
			Type("multipart").
			Send(`{"grant_type": "refresh_token"}`).
			Send(fmt.Sprintf(`{"refresh_token": %v}`, c.RefreshToken)).
			Send(`{"includePerms": true}`).
			Set("Authorization", fmt.Sprintf("basic %v", c.ClientLauncherToken)).
			EndStruct(&response)

		if err != nil {
			log.Fatal("Error with check token request:", err)
		}

		c.Mutex.Lock()
		c.AccessTokenExpiresAt, _ = time.Parse(time.RFC3339Nano, response.ExpiresAt)
		c.AccessToken = response.AccessToken
		c.RefreshToken = response.RefreshToken
		c.Mutex.Unlock()
	}
}
