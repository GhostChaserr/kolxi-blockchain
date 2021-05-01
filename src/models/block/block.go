package block

type Block struct {
	Id       int    `json:"id"`
	PrevHash string `json:"PrevHash"`
	Data     string `json:"data"`
	Nonce    int    `json:"nonce"`
}
