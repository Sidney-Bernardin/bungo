package bungo

type ItemActionRequest struct {
	ItemId         string `json:"itemId"`
	CharacterId    string `json:"characterId"`
	MembershipType string `json:"membershipType"`
}
