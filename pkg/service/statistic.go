package service

import (
	urlshortener "github.com/ramil66/url-shortener"
	"github.com/ramil66/url-shortener/pkg/repository"
)

type StatisticService struct {
	repo repository.Statistic
}

func NewStatisticService(repo repository.Statistic) *StatisticService {
	return &StatisticService{repo: repo}
}

func (s *StatisticService) GetMetric(alias string) ([]urlshortener.Statistic, error) {
	return s.repo.GetMetric(alias)
}

func (s *StatisticService) SaveStatistic(st urlshortener.Statistic) error {
	return s.repo.SaveStatistic(st)
}
