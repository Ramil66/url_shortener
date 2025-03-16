package repository

import (
	"github.com/jmoiron/sqlx"
	urlshortener "github.com/ramil66/url-shortener"
)

type Authorization interface {
	CreateUser(user urlshortener.User) (int, error)
	GetUser(email, password string) (urlshortener.User, error)
}

type Url interface {
	SaveUrl(url string, alias string) (string, error)
	SaveUrlUsers(userId int, url string, alias string) (string, error)
	GetAllUrls(userId int) ([]urlshortener.Url, error)
	GetUrl(alias string) (string, error)
	CheckUrl(url string) int
	CheckUserUrl(url string, userId int) bool
	CheckAlias(alias string) bool
	DeleteUrl(alias string) error
	GetIdUrl(alias string) (int, error)
	CheckLinkUrlUser(alias string) bool
	IncrementCounter(alias string) error
}

type Statistic interface {
	GetMetric(alias string) ([]urlshortener.Statistic, error)
	SaveStatistic(st urlshortener.Statistic) error
	CheckStat(ip string, device string, urlId int) bool
}

type Repository struct {
	Authorization
	Url
	Statistic
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Url:           NewUrlPostgres(db),
		Statistic:     NewStatisticPostgres(db),
	}
}
