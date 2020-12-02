package bungo

import "time"

type GetEntityDefinitionResponse struct {
	Response struct {
		DisplayProperties struct {
			Description string `json:"description"`
			Name        string `json:"name"`
			Icon        string `json:"icon"`
			HasIcon     bool   `json:"hasIcon"`
		} `json:"displayProperties"`
		TooltipNotifications []interface{} `json:"tooltipNotifications"`
		CollectibleHash      int64         `json:"collectibleHash"`
		BackgroundColor      struct {
			Red   int `json:"red"`
			Green int `json:"green"`
			Blue  int `json:"blue"`
			Alpha int `json:"alpha"`
		} `json:"backgroundColor"`
		Screenshot                 string `json:"screenshot"`
		ItemTypeDisplayName        string `json:"itemTypeDisplayName"`
		UIItemDisplayStyle         string `json:"uiItemDisplayStyle"`
		ItemTypeAndTierDisplayName string `json:"itemTypeAndTierDisplayName"`
		DisplaySource              string `json:"displaySource"`
		Action                     struct {
			VerbName                string        `json:"verbName"`
			VerbDescription         string        `json:"verbDescription"`
			IsPositive              bool          `json:"isPositive"`
			RequiredCooldownSeconds int           `json:"requiredCooldownSeconds"`
			RequiredItems           []interface{} `json:"requiredItems"`
			ProgressionRewards      []interface{} `json:"progressionRewards"`
			ActionTypeLabel         string        `json:"actionTypeLabel"`
			RewardSheetHash         int           `json:"rewardSheetHash"`
			RewardItemHash          int           `json:"rewardItemHash"`
			RewardSiteHash          int           `json:"rewardSiteHash"`
			RequiredCooldownHash    int           `json:"requiredCooldownHash"`
			DeleteOnAction          bool          `json:"deleteOnAction"`
			ConsumeEntireStack      bool          `json:"consumeEntireStack"`
			UseOnAcquire            bool          `json:"useOnAcquire"`
		} `json:"action"`
		Inventory struct {
			MaxStackSize                             int    `json:"maxStackSize"`
			BucketTypeHash                           int    `json:"bucketTypeHash"`
			RecoveryBucketTypeHash                   int    `json:"recoveryBucketTypeHash"`
			TierTypeHash                             int64  `json:"tierTypeHash"`
			IsInstanceItem                           bool   `json:"isInstanceItem"`
			NonTransferrableOriginal                 bool   `json:"nonTransferrableOriginal"`
			TierTypeName                             string `json:"tierTypeName"`
			TierType                                 int    `json:"tierType"`
			ExpirationTooltip                        string `json:"expirationTooltip"`
			ExpiredInActivityMessage                 string `json:"expiredInActivityMessage"`
			ExpiredInOrbitMessage                    string `json:"expiredInOrbitMessage"`
			SuppressExpirationWhenObjectivesComplete bool   `json:"suppressExpirationWhenObjectivesComplete"`
		} `json:"inventory"`
		Stats struct {
			DisablePrimaryStatDisplay bool  `json:"disablePrimaryStatDisplay"`
			StatGroupHash             int64 `json:"statGroupHash"`
			Stats                     struct {
				Num155624089 struct {
					StatHash       int `json:"statHash"`
					Value          int `json:"value"`
					Minimum        int `json:"minimum"`
					Maximum        int `json:"maximum"`
					DisplayMaximum int `json:"displayMaximum"`
				} `json:"155624089"`
				Num447667954 struct {
					StatHash       int `json:"statHash"`
					Value          int `json:"value"`
					Minimum        int `json:"minimum"`
					Maximum        int `json:"maximum"`
					DisplayMaximum int `json:"displayMaximum"`
				} `json:"447667954"`
				Num943549884 struct {
					StatHash       int `json:"statHash"`
					Value          int `json:"value"`
					Minimum        int `json:"minimum"`
					Maximum        int `json:"maximum"`
					DisplayMaximum int `json:"displayMaximum"`
				} `json:"943549884"`
				Num1345609583 struct {
					StatHash       int `json:"statHash"`
					Value          int `json:"value"`
					Minimum        int `json:"minimum"`
					Maximum        int `json:"maximum"`
					DisplayMaximum int `json:"displayMaximum"`
				} `json:"1345609583"`
				Num1480404414 struct {
					StatHash       int `json:"statHash"`
					Value          int `json:"value"`
					Minimum        int `json:"minimum"`
					Maximum        int `json:"maximum"`
					DisplayMaximum int `json:"displayMaximum"`
				} `json:"1480404414"`
				Num1591432999 struct {
					StatHash       int `json:"statHash"`
					Value          int `json:"value"`
					Minimum        int `json:"minimum"`
					Maximum        int `json:"maximum"`
					DisplayMaximum int `json:"displayMaximum"`
				} `json:"1591432999"`
				Num1885944937 struct {
					StatHash       int `json:"statHash"`
					Value          int `json:"value"`
					Minimum        int `json:"minimum"`
					Maximum        int `json:"maximum"`
					DisplayMaximum int `json:"displayMaximum"`
				} `json:"1885944937"`
				Num1931675084 struct {
					StatHash       int `json:"statHash"`
					Value          int `json:"value"`
					Minimum        int `json:"minimum"`
					Maximum        int `json:"maximum"`
					DisplayMaximum int `json:"displayMaximum"`
				} `json:"1931675084"`
				Num1935470627 struct {
					StatHash       int `json:"statHash"`
					Value          int `json:"value"`
					Minimum        int `json:"minimum"`
					Maximum        int `json:"maximum"`
					DisplayMaximum int `json:"displayMaximum"`
				} `json:"1935470627"`
				Num2715839340 struct {
					StatHash       int64 `json:"statHash"`
					Value          int   `json:"value"`
					Minimum        int   `json:"minimum"`
					Maximum        int   `json:"maximum"`
					DisplayMaximum int   `json:"displayMaximum"`
				} `json:"2715839340"`
				Num2961396640 struct {
					StatHash       int64 `json:"statHash"`
					Value          int   `json:"value"`
					Minimum        int   `json:"minimum"`
					Maximum        int   `json:"maximum"`
					DisplayMaximum int   `json:"displayMaximum"`
				} `json:"2961396640"`
				Num3555269338 struct {
					StatHash       int64 `json:"statHash"`
					Value          int   `json:"value"`
					Minimum        int   `json:"minimum"`
					Maximum        int   `json:"maximum"`
					DisplayMaximum int   `json:"displayMaximum"`
				} `json:"3555269338"`
				Num4043523819 struct {
					StatHash       int64 `json:"statHash"`
					Value          int   `json:"value"`
					Minimum        int   `json:"minimum"`
					Maximum        int   `json:"maximum"`
					DisplayMaximum int   `json:"displayMaximum"`
				} `json:"4043523819"`
				Num4188031367 struct {
					StatHash       int64 `json:"statHash"`
					Value          int   `json:"value"`
					Minimum        int   `json:"minimum"`
					Maximum        int   `json:"maximum"`
					DisplayMaximum int   `json:"displayMaximum"`
				} `json:"4188031367"`
			} `json:"stats"`
			HasDisplayableStats bool `json:"hasDisplayableStats"`
			PrimaryBaseStatHash int  `json:"primaryBaseStatHash"`
		} `json:"stats"`
		EquippingBlock struct {
			UniqueLabelHash       int      `json:"uniqueLabelHash"`
			EquipmentSlotTypeHash int      `json:"equipmentSlotTypeHash"`
			Attributes            int      `json:"attributes"`
			EquippingSoundHash    int      `json:"equippingSoundHash"`
			HornSoundHash         int      `json:"hornSoundHash"`
			AmmoType              int      `json:"ammoType"`
			DisplayStrings        []string `json:"displayStrings"`
		} `json:"equippingBlock"`
		TranslationBlock struct {
			WeaponPatternHash int `json:"weaponPatternHash"`
			DefaultDyes       []struct {
				ChannelHash int `json:"channelHash"`
				DyeHash     int `json:"dyeHash"`
			} `json:"defaultDyes"`
			LockedDyes   []interface{} `json:"lockedDyes"`
			CustomDyes   []interface{} `json:"customDyes"`
			Arrangements []struct {
				ClassHash          int `json:"classHash"`
				ArtArrangementHash int `json:"artArrangementHash"`
			} `json:"arrangements"`
			HasGeometry bool `json:"hasGeometry"`
		} `json:"translationBlock"`
		Preview struct {
			ScreenStyle         string `json:"screenStyle"`
			PreviewVendorHash   int    `json:"previewVendorHash"`
			PreviewActionString string `json:"previewActionString"`
		} `json:"preview"`
		Quality struct {
			ItemLevels                      []interface{} `json:"itemLevels"`
			QualityLevel                    int           `json:"qualityLevel"`
			InfusionCategoryName            string        `json:"infusionCategoryName"`
			InfusionCategoryHash            int64         `json:"infusionCategoryHash"`
			InfusionCategoryHashes          []int64       `json:"infusionCategoryHashes"`
			ProgressionLevelRequirementHash int64         `json:"progressionLevelRequirementHash"`
			CurrentVersion                  int           `json:"currentVersion"`
			Versions                        []struct {
				PowerCapHash int `json:"powerCapHash"`
			} `json:"versions"`
			DisplayVersionWatermarkIcons []string `json:"displayVersionWatermarkIcons"`
		} `json:"quality"`
		AcquireRewardSiteHash int `json:"acquireRewardSiteHash"`
		AcquireUnlockHash     int `json:"acquireUnlockHash"`
		Sockets               struct {
			Detail        string `json:"detail"`
			SocketEntries []struct {
				SocketTypeHash                        int64         `json:"socketTypeHash"`
				SingleInitialItemHash                 int           `json:"singleInitialItemHash"`
				ReusablePlugItems                     []interface{} `json:"reusablePlugItems"`
				PreventInitializationOnVendorPurchase bool          `json:"preventInitializationOnVendorPurchase"`
				PreventInitializationWhenVersioning   bool          `json:"preventInitializationWhenVersioning"`
				HidePerksInItemTooltip                bool          `json:"hidePerksInItemTooltip"`
				PlugSources                           int           `json:"plugSources"`
				ReusablePlugSetHash                   int           `json:"reusablePlugSetHash"`
				OverridesUIAppearance                 bool          `json:"overridesUiAppearance"`
				DefaultVisible                        bool          `json:"defaultVisible"`
			} `json:"socketEntries"`
			IntrinsicSockets []struct {
				PlugItemHash   int64 `json:"plugItemHash"`
				SocketTypeHash int   `json:"socketTypeHash"`
				DefaultVisible bool  `json:"defaultVisible"`
			} `json:"intrinsicSockets"`
			SocketCategories []struct {
				SocketCategoryHash int64 `json:"socketCategoryHash"`
				SocketIndexes      []int `json:"socketIndexes"`
			} `json:"socketCategories"`
		} `json:"sockets"`
		TalentGrid struct {
			TalentGridHash   int    `json:"talentGridHash"`
			ItemDetailString string `json:"itemDetailString"`
			HudDamageType    int    `json:"hudDamageType"`
		} `json:"talentGrid"`
		InvestmentStats []struct {
			StatTypeHash          int  `json:"statTypeHash"`
			Value                 int  `json:"value"`
			IsConditionallyActive bool `json:"isConditionallyActive"`
		} `json:"investmentStats"`
		Perks []struct {
			RequirementDisplayString string `json:"requirementDisplayString"`
			PerkHash                 int64  `json:"perkHash"`
			PerkVisibility           int    `json:"perkVisibility"`
		} `json:"perks"`
		SummaryItemHash                   int64         `json:"summaryItemHash"`
		AllowActions                      bool          `json:"allowActions"`
		DoesPostmasterPullHaveSideEffects bool          `json:"doesPostmasterPullHaveSideEffects"`
		NonTransferrable                  bool          `json:"nonTransferrable"`
		ItemCategoryHashes                []interface{} `json:"itemCategoryHashes"`
		SpecialItemType                   int           `json:"specialItemType"`
		ItemType                          int           `json:"itemType"`
		ItemSubType                       int           `json:"itemSubType"`
		ClassType                         int           `json:"classType"`
		BreakerType                       int           `json:"breakerType"`
		Equippable                        bool          `json:"equippable"`
		DamageTypeHashes                  []int64       `json:"damageTypeHashes"`
		DamageTypes                       []int         `json:"damageTypes"`
		DefaultDamageType                 int           `json:"defaultDamageType"`
		DefaultDamageTypeHash             int64         `json:"defaultDamageTypeHash"`
		IsWrapper                         bool          `json:"isWrapper"`
		TraitIds                          []string      `json:"traitIds"`
		Hash                              int           `json:"hash"`
		Index                             int           `json:"index"`
		Redacted                          bool          `json:"redacted"`
		Blacklisted                       bool          `json:"blacklisted"`
	} `json:"Response"`
	ErrorCode       errorCode `json:"ErrorCode"`
	ThrottleSeconds int       `json:"ThrottleSeconds"`
	ErrorStatus     string    `json:"ErrorStatus"`
	Message         string    `json:"Message"`
	MessageData     struct {
	} `json:"MessageData"`
}

type SearchDestinyPlayerResponse struct {
	Response []struct {
		IconPath          string `json:"iconPath"`
		CrossSaveOverride int    `json:"crossSaveOverride"`
		IsPublic          bool   `json:"isPublic"`
		MembershipType    int    `json:"membershipType"`
		MembershipID      string `json:"membershipId"`
		DisplayName       string `json:"displayName"`
	} `json:"Response"`
	ErrorCode       errorCode `json:"ErrorCode"`
	ThrottleSeconds int       `json:"ThrottleSeconds"`
	ErrorStatus     string    `json:"ErrorStatus"`
	Message         string    `json:"Message"`
	MessageData     struct {
	} `json:"MessageData"`
}

type LinkedProfilesResponse struct {
	Response struct {
		Profiles []struct {
			DateLastPlayed            time.Time `json:"dateLastPlayed"`
			IsOverridden              bool      `json:"isOverridden"`
			IsCrossSavePrimary        bool      `json:"isCrossSavePrimary"`
			CrossSaveOverride         int       `json:"crossSaveOverride"`
			ApplicableMembershipTypes []int     `json:"applicableMembershipTypes"`
			IsPublic                  bool      `json:"isPublic"`
			MembershipType            int       `json:"membershipType"`
			MembershipID              string    `json:"membershipId"`
			DisplayName               string    `json:"displayName"`
		} `json:"profiles"`
		BnetMembership struct {
			SupplementalDisplayName string `json:"supplementalDisplayName"`
			IconPath                string `json:"iconPath"`
			CrossSaveOverride       int    `json:"crossSaveOverride"`
			IsPublic                bool   `json:"isPublic"`
			MembershipType          int    `json:"membershipType"`
			MembershipID            string `json:"membershipId"`
			DisplayName             string `json:"displayName"`
		} `json:"bnetMembership"`
		ProfilesWithErrors []interface{} `json:"profilesWithErrors"`
	}
	ErrorCode       errorCode `json:"ErrorCode"`
	ThrottleSeconds int       `json:"ThrottleSeconds"`
	ErrorStatus     string    `json:"ErrorStatus"`
	Message         string    `json:"Message"`
	MessageData     struct {
	} `json:"MessageData"`
}

type GetProifleResponse struct {
	Profiles SingleComponentResponseOfDestinyProfileComponent
}

type SingleComponentResponseOfDestinyProfileComponent struct {
	Response struct {
		Profile struct {
			Data struct {
				UserInfo struct {
					CrossSaveOverride         int    `json:"crossSaveOverride"`
					ApplicableMembershipTypes []int  `json:"applicableMembershipTypes"`
					IsPublic                  bool   `json:"isPublic"`
					MembershipType            int    `json:"membershipType"`
					MembershipID              string `json:"membershipId"`
					DisplayName               string `json:"displayName"`
				} `json:"userInfo"`
				DateLastPlayed              time.Time     `json:"dateLastPlayed"`
				VersionsOwned               int           `json:"versionsOwned"`
				CharacterIds                []string      `json:"characterIds"`
				SeasonHashes                []interface{} `json:"seasonHashes"`
				CurrentSeasonHash           int           `json:"currentSeasonHash"`
				CurrentSeasonRewardPowerCap int           `json:"currentSeasonRewardPowerCap"`
			} `json:"data"`
			Privacy int `json:"privacy"`
		} `json:"profile"`
	} `json:"Response"`
	ErrorCode       errorCode `json:"ErrorCode"`
	ThrottleSeconds int       `json:"ThrottleSeconds"`
	ErrorStatus     string    `json:"ErrorStatus"`
	Message         string    `json:"Message"`
	MessageData     struct {
	} `json:"MessageData"`
}

type DestinyCharacterResponse struct {
	Characters           SingleComponentResponseOfDestinyCharacterComponent
	CharacterInventories SingleComponentResponseOfDestinyInventoryComponent
	CharacterEquipment   SingleComponentResponseOfDestinyEquipmentComponent
}

type SingleComponentResponseOfDestinyCharacterComponent struct {
	Response struct {
		Character struct {
			Data struct {
				MembershipID             string    `json:"membershipId"`
				MembershipType           int       `json:"membershipType"`
				CharacterID              string    `json:"characterId"`
				DateLastPlayed           time.Time `json:"dateLastPlayed"`
				MinutesPlayedThisSession string    `json:"minutesPlayedThisSession"`
				MinutesPlayedTotal       string    `json:"minutesPlayedTotal"`
				Light                    int       `json:"light"`
				Stats                    struct {
					Num144602215  int `json:"144602215"`
					Num392767087  int `json:"392767087"`
					Num1735777505 int `json:"1735777505"`
					Num1935470627 int `json:"1935470627"`
					Num1943323491 int `json:"1943323491"`
					Num2996146975 int `json:"2996146975"`
					Num4244567218 int `json:"4244567218"`
				} `json:"stats"`
				RaceHash             int    `json:"raceHash"`
				GenderHash           int64  `json:"genderHash"`
				ClassHash            int64  `json:"classHash"`
				RaceType             int    `json:"raceType"`
				ClassType            int    `json:"classType"`
				GenderType           int    `json:"genderType"`
				EmblemPath           string `json:"emblemPath"`
				EmblemBackgroundPath string `json:"emblemBackgroundPath"`
				EmblemHash           int    `json:"emblemHash"`
				EmblemColor          struct {
					Red   int `json:"red"`
					Green int `json:"green"`
					Blue  int `json:"blue"`
					Alpha int `json:"alpha"`
				} `json:"emblemColor"`
				LevelProgression struct {
					ProgressionHash     int `json:"progressionHash"`
					DailyProgress       int `json:"dailyProgress"`
					DailyLimit          int `json:"dailyLimit"`
					WeeklyProgress      int `json:"weeklyProgress"`
					WeeklyLimit         int `json:"weeklyLimit"`
					CurrentProgress     int `json:"currentProgress"`
					Level               int `json:"level"`
					LevelCap            int `json:"levelCap"`
					StepIndex           int `json:"stepIndex"`
					ProgressToNextLevel int `json:"progressToNextLevel"`
					NextLevelAt         int `json:"nextLevelAt"`
				} `json:"levelProgression"`
				BaseCharacterLevel float64 `json:"baseCharacterLevel"`
				PercentToNextLevel float64 `json:"percentToNextLevel"`
			} `json:"data"`
			Privacy int `json:"privacy"`
		} `json:"character"`
	} `json:"Response"`
	ErrorCode       errorCode `json:"ErrorCode"`
	ThrottleSeconds int       `json:"ThrottleSeconds"`
	ErrorStatus     string    `json:"ErrorStatus"`
	Message         string    `json:"Message"`
	MessageData     struct {
	} `json:"MessageData"`
}

type SingleComponentResponseOfDestinyInventoryComponent struct {
	Response struct {
		Inventory struct {
			Data struct {
				Items []struct {
					ItemHash              int    `json:"itemHash"`
					ItemInstanceID        string `json:"itemInstanceId"`
					Quantity              int    `json:"quantity"`
					BindStatus            int    `json:"bindStatus"`
					Location              int    `json:"location"`
					BucketHash            int    `json:"bucketHash"`
					TransferStatus        int    `json:"transferStatus"`
					Lockable              bool   `json:"lockable"`
					State                 int    `json:"state"`
					DismantlePermission   int    `json:"dismantlePermission"`
					IsWrapper             bool   `json:"isWrapper"`
					VersionNumber         int    `json:"versionNumber,omitempty"`
					OverrideStyleItemHash int    `json:"overrideStyleItemHash,omitempty"`
					MetricHash            int    `json:"metricHash,omitempty"`
					MetricObjective       struct {
						ObjectiveHash   int64 `json:"objectiveHash"`
						Progress        int   `json:"progress"`
						CompletionValue int   `json:"completionValue"`
						Complete        bool  `json:"complete"`
						Visible         bool  `json:"visible"`
					} `json:"metricObjective,omitempty"`
				} `json:"items"`
			} `json:"data"`
			Privacy int `json:"privacy"`
		} `json:"inventory"`
	} `json:"Response"`
	ErrorCode       errorCode `json:"ErrorCode"`
	ThrottleSeconds int       `json:"ThrottleSeconds"`
	ErrorStatus     string    `json:"ErrorStatus"`
	Message         string    `json:"Message"`
	MessageData     struct {
	} `json:"MessageData"`
}

type SingleComponentResponseOfDestinyEquipmentComponent struct {
	Response struct {
		Equipment struct {
			Data struct {
				Items []struct {
					ItemHash              int    `json:"itemHash"`
					ItemInstanceID        string `json:"itemInstanceId"`
					Quantity              int    `json:"quantity"`
					BindStatus            int    `json:"bindStatus"`
					Location              int    `json:"location"`
					BucketHash            int    `json:"bucketHash"`
					TransferStatus        int    `json:"transferStatus"`
					Lockable              bool   `json:"lockable"`
					State                 int    `json:"state"`
					DismantlePermission   int    `json:"dismantlePermission"`
					IsWrapper             bool   `json:"isWrapper"`
					VersionNumber         int    `json:"versionNumber,omitempty"`
					OverrideStyleItemHash int    `json:"overrideStyleItemHash,omitempty"`
					MetricHash            int    `json:"metricHash,omitempty"`
					MetricObjective       struct {
						ObjectiveHash   int64 `json:"objectiveHash"`
						Progress        int   `json:"progress"`
						CompletionValue int   `json:"completionValue"`
						Complete        bool  `json:"complete"`
						Visible         bool  `json:"visible"`
					} `json:"metricObjective,omitempty"`
				} `json:"items"`
			} `json:"data"`
			Privacy int `json:"privacy"`
		} `json:"equipment"`
	} `json:"Response"`
	ErrorCode       errorCode `json:"ErrorCode"`
	ThrottleSeconds int       `json:"ThrottleSeconds"`
	ErrorStatus     string    `json:"ErrorStatus"`
	Message         string    `json:"Message"`
	MessageData     struct {
	} `json:"MessageData"`
}

type EquipItemResponse struct {
	Response        int       `json:"Response"`
	ErrorCode       errorCode `json:"ErrorCode"`
	ThrottleSeconds int       `json:"ThrottleSeconds"`
	ErrorStatus     string    `json:"ErrorStatus"`
	Message         string    `json:"Message"`
	MessageData     struct {
	} `json:"MessageData"`
}
