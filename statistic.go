package urlshortener

type Statistic struct {
	Id       int    `json:"-" db:"id"`
	UrlId    int    `json:"url_id" db:"url_id" binding:"required"`
	Ip       string `json:"ip" db:"ip" binding:"required"`
	Device   string `json:"device" db:"device" binding:"required"`
	LastDate string `json:"last_date" db:"last_date" binding:"required"`
}
