package genesisblock

type Balance struct {
	Username string `json:"username"`
}

type GenesisBlock struct {
	Id        string `json:"id"`
	CreatedAt string `json:"createdAt"`
}
