package responses

type UserInfoCard struct {
	IconPath          string `json:"iconPath"`
	CrossSaveOverride int    `json:"crossSaveOverride"`
	IsPublic          bool   `json:"isPublic"`
	MembershipType    int    `json:"membershipType"`
	MembershipID      string `json:"membershipId"`
	DisplayName       string `json:"displayName"`
}
