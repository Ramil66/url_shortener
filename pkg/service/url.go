package service

import (
	"fmt"
	"math/rand"
	"time"

	urlshortener "github.com/ramil66/url-shortener"
	"github.com/ramil66/url-shortener/pkg/repository"
)

type UrlService struct {
	repo repository.Url
}

func NewUrlService(repo repository.Url) *UrlService {
	return &UrlService{repo: repo}
}

func GenerateShortUrl(size int) string {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))

	chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
		"abcdefghijklmnopqrstuvwxyz" +
		"0123456789")

	b := make([]rune, size)
	for i := range b {
		b[i] = chars[rnd.Intn(len(chars))]
	}
	fmt.Println(string(b))
	return string(b)
}

func (s *UrlService) Shortening(url string) (string, error) {
	alias := GenerateShortUrl(5)

	for !s.repo.CheckAlias(alias) {
		alias = GenerateShortUrl(5)
	}

	return s.repo.SaveUrl(url, alias)
}

func (s *UrlService) IncrementCounter(alias string) error {
	return s.repo.IncrementCounter(alias)
}
func (s *UrlService) ShorteningUrlUsers(userId int, url string) (string, error) {
	alias := GenerateShortUrl(5)

	for !s.repo.CheckAlias(alias) {
		alias = GenerateShortUrl(5)
	}
	fmt.Println("Service: " + alias)
	return s.repo.SaveUrlUsers(userId, url, alias)
}

func (s *UrlService) GetAllUrls(userId int) ([]urlshortener.Url, error) {
	return s.repo.GetAllUrls(userId)
}
func (s *UrlService) GetUrl(alias string) (string, error) {
	return s.repo.GetUrl(alias)
}

func (s *UrlService) DeleteUrl(alias string) error {
	return s.repo.DeleteUrl(alias)
}

func (s *UrlService) CheckAlias(alias string) bool {
	return s.repo.CheckAlias(alias)
}

func (s *UrlService) CheckLinkUrlUser(alias string) bool {
	return s.repo.CheckLinkUrlUser(alias)
}

func (s *UrlService) CustomUrl(userId int, url string, alias string) (string, error) {
	return s.repo.SaveUrlUsers(userId, url, alias)
	//return s.repo.CustomUrl(userId, url, alias)
}

func (s *UrlService) GetIdUrl(alias string) (int, error) {
	return s.repo.GetIdUrl(alias)
}
