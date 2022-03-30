package resource

type Wallet struct {
	ID        int64  `json:"id"`
	Balance   string `json:"balance"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
