package gotoorbit

type SearchDestinyPlayerResponse struct {
	Response []struct {
		IconPath          string `json:"iconPath"`
		CrossSaveOverride int    `json:"crossSaveOverride"`
		IsPublic          bool   `json:"isPublic"`
		MembershipType    int    `json:"membershipType"`
		MembershipID      string `json:"membershipId"`
		DisplayName       string `json:"displayName"`
	} `json:"Response"`
	ErrorCode       int    `json:"ErrorCode"`
	ThrottleSeconds int    `json:"ThrottleSeconds"`
	ErrorStatus     string `json:"ErrorStatus"`
	Message         string `json:"Message"`
	MessageData     struct {
	} `json:"MessageData"`
}
