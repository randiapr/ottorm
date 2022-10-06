package thirdparty

type Clients struct {
	ClientId     string `gorm:"primaryKey"`
	ClientSecret string
	PublicKey    string
	ClientName   string
}

func (Clients) TableName() string {
	return "third_party.clients"
}
