package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	urlshortener "github.com/ramil66/url-shortener"
)

type StatisticPostgres struct {
	db *sqlx.DB
}

func NewStatisticPostgres(db *sqlx.DB) *StatisticPostgres {
	return &StatisticPostgres{db: db}
}

func (r *StatisticPostgres) GetMetric(alias string) ([]urlshortener.Statistic, error) {
	var statistics []urlshortener.Statistic
	query := fmt.Sprintf("SELECT st.id,st.url_id,st.ip,st.device,st.last_date FROM %s st INNER JOIN %s url ON st.url_id=url.id WHERE url.alias=$1", statTable, urlTable)
	err := r.db.Select(&statistics, query, alias)

	return statistics, err
}

func (r *StatisticPostgres) CheckStat(ip string, device string, urlId int) bool {
	id := 0
	query := fmt.Sprintf("SELECT st.id FROM %s st INNER JOIN %s url ON st.url_id=url.id WHERE st.ip=$1 AND st.device=$2 AND url.id=$3 ", statTable, urlTable)
	r.db.Get(&id, query, ip, device, urlId)
	return id == 0
}

func (r *StatisticPostgres) SaveStatistic(st urlshortener.Statistic) error {
	// if !r.CheckStat(st.Ip, st.Device, st.UrlId) {
	// 	fmt.Println("update")
	// 	query := fmt.Sprintf("UPDATE %s SET last_date=$1 WHERE ip=$2 AND device=$3", statTable)
	// 	_, err := r.db.Exec(query, st.LastDate, st.Ip, st.Device)
	// 	return err
	// }
	fmt.Println("insert")
	query := fmt.Sprintf("INSERT INTO %s (url_id,ip,device,last_date) values ($1,$2,$3,$4) RETURNING id", statTable)
	_, err := r.db.Exec(query, st.UrlId, st.Ip, st.Device, st.LastDate)
	return err
}
