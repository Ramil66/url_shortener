package service

import (
	urlshortener "github.com/ramil66/url-shortener"
	"github.com/ramil66/url-shortener/pkg/repository"
)

type Authorization interface {
	CreateUser(user urlshortener.User) (int, error)
	GenerateToken(email, password string) (string, error)
	ParseToken(token string) (int, error)
}

type Url interface {
	Shortening(url string) (string, error)
	ShorteningUrlUsers(userId int, url string) (string, error)
	GetAllUrls(userId int) ([]urlshortener.Url, error)
	GetUrl(alias string) (string, error)
	GetIdUrl(alias string) (int, error)
	CheckAlias(alias string) bool
	CheckLinkUrlUser(alias string) bool
	DeleteUrl(alias string) error
	CustomUrl(userId int, url string, alias string) (string, error)
	IncrementCounter(alias string) error
}

type Statistic interface {
	GetMetric(alias string) ([]urlshortener.Statistic, error)
	SaveStatistic(st urlshortener.Statistic) error
}

type Service struct {
	Authorization
	Url
	Statistic
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Url:           NewUrlService(repos.Url),
		Statistic:     NewStatisticService(repos.Statistic),
	}
}
