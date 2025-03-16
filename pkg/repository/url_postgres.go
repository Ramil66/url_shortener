package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	urlshortener "github.com/ramil66/url-shortener"
)

type UrlPostgres struct {
	db *sqlx.DB
}

const BaseUrl = "http://localhost:8080/"

func NewUrlPostgres(db *sqlx.DB) *UrlPostgres {
	return &UrlPostgres{db: db}
}
func (r *UrlPostgres) CheckUrl(url string) int {
	id := 0
	query := fmt.Sprintf("SELECT id FROM %s WHERE url=$1", urlTable)
	r.db.Get(&id, query, url)
	return id
}
func (r *UrlPostgres) CheckUserUrl(url string, userId int) bool {
	id := 0
	urlId := r.CheckUrl(url)
	query := fmt.Sprintf("SELECT id FROM %s WHERE url_id=$1 AND user_id=$2", userUrlTable)
	r.db.Get(&id, query, urlId, userId)
	return id == 0
}
func (r *UrlPostgres) CheckAlias(alias string) bool {
	id := 0
	query := fmt.Sprintf("SELECT id FROM %s WHERE alias=$1", urlTable)
	r.db.Get(&id, query, alias)
	return id == 0
}
func (r *UrlPostgres) SaveUrl(url string, alias string) (string, error) {
	fmt.Println(url)
	fmt.Println(alias)
	fmt.Println("SaveUrl")
	urlId := 0
	if id := r.CheckUrl(url); id != 0 {
		fmt.Println(id)
		var url urlshortener.Url
		url, err := r.GetUrlById(id)
		if err != nil {
			return "", err
		}
		return url.Alias, nil
	}
	query := fmt.Sprintf("INSERT INTO %s (url,alias,counter) values ($1,$2,0) RETURNING id", urlTable)
	row := r.db.DB.QueryRow(query, url, alias)
	if err := row.Scan(&urlId); err != nil {
		return "", err
	}
	return (BaseUrl + alias), nil
}

func (r *UrlPostgres) SaveUrlUsers(userId int, url string, alias string) (string, error) {
	var urlId int
	fmt.Println(url + " " + alias)
	query := fmt.Sprintf("INSERT INTO %s (url,alias,counter) values ($1,$2,0) RETURNING id", urlTable)
	row := r.db.DB.QueryRow(query, url, alias)
	if err := row.Scan(&urlId); err != nil {
		return "", err
	}
	query = fmt.Sprintf("INSERT INTO %s (user_id,url_id) values ($1,$2) RETURNING id", userUrlTable)
	_, err := r.db.DB.Exec(query, userId, urlId)
	if err != nil {
		return "", err
	}
	return (BaseUrl + alias), nil
}

func (r *UrlPostgres) GetUrlById(urlId int) (urlshortener.Url, error) {
	var url urlshortener.Url
	query := fmt.Sprintf("SELECT url.id,url.url,url.alias,url.counter FROM %s url WHERE url.id=$1", urlTable)
	fmt.Println("getUrlBy id:1")
	err := r.db.Get(&url, query, urlId)
	fmt.Println(url.Alias)
	return url, err
}

func (r *UrlPostgres) GetAllUrls(userId int) ([]urlshortener.Url, error) {
	var urls []urlshortener.Url

	query := fmt.Sprintf("SELECT url.id,url.url,url.alias,url.counter FROM %s url INNER JOIN %s uu ON url.id=uu.url_id WHERE uu.user_id=$1", urlTable, userUrlTable)
	err := r.db.Select(&urls, query, userId)

	return urls, err
}

func (r *UrlPostgres) GetUrl(alias string) (string, error) {
	var url string
	fmt.Println(alias)
	query := fmt.Sprintf("SELECT url FROM %s WHERE alias=$1", urlTable)
	row := r.db.DB.QueryRow(query, alias)
	if err := row.Scan(&url); err != nil {
		return "", err
	}
	return url, nil
}

func (r *UrlPostgres) GetIdUrl(alias string) (int, error) {
	var urlId int
	query := fmt.Sprintf("SELECT id FROM %s WHERE alias=$1", urlTable)
	row := r.db.DB.QueryRow(query, alias)
	if err := row.Scan(&urlId); err != nil {
		return 0, err
	}

	return urlId, nil
}

func (r *UrlPostgres) DeleteUrl(alias string) error {
	var urlId int
	fmt.Println("Delete: 1")
	fmt.Println(alias)
	query := fmt.Sprintf("SELECT id FROM %s WHERE alias=$1", urlTable)
	row := r.db.DB.QueryRow(query, alias)
	if err := row.Scan(&urlId); err != nil {
		return err
	}
	fmt.Println("Delete: 2")
	query = fmt.Sprintf("DELETE FROM %s WHERE url_id=$1", userUrlTable)
	_, err := r.db.DB.Exec(query, urlId)

	if err != nil {
		return err
	}
	fmt.Println("Delete: 3")
	query = fmt.Sprintf("DELETE FROM %s WHERE id=$1", urlTable)

	_, err = r.db.DB.Exec(query, urlId)
	if err != nil {
		return err
	}

	return nil
}
func (r *UrlPostgres) IncrementCounter(alias string) error {
	query := fmt.Sprintf("UPDATE %s SET counter = counter + 1 WHERE alias=$1", urlTable)
	_, err := r.db.Exec(query, alias)
	return err
}
func (r *UrlPostgres) CheckLinkUrlUser(alias string) bool {
	id := 0
	query := fmt.Sprintf("SELECT uut.id FROM %s uut INNER JOIN %s ut ON uut.url_id=ut.id WHERE ut.alias=$1", userUrlTable, urlTable)
	r.db.Get(&id, query, alias)
	return id == 0
}
