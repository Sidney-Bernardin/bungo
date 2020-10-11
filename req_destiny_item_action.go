package bungo

type DestinyItemActionRequest struct {
	ItemId         string `json:"itemId"`
	CharacterId    string `json:"characterId"`
	MembershipType string `json:"membershipType"`
}
