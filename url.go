package urlshortener

type Url struct {
	Id      int    `json:"-" db:"id"`
	Url     string `json:"url" db:"url" binding:"required"`
	Alias   string `json:"alias" db:"alias"`
	Counter int    `json:"counter" db:"counter"`
}

type UserUrl struct {
	Id     int
	UserId int
	UrlId  int
}
