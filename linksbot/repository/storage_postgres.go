package repository

import (
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
	"math/rand"
	"time"
)

type StoragePostgres struct {
	db *sqlx.DB
}

func NewStoragePostgres(db *sqlx.DB) *StoragePostgres {
	return &StoragePostgres{db: db}
}

func (s *StoragePostgres) Save(p *Page) error {
	if s.IsPresent(p) == false {
		query := fmt.Sprintf("INSERT INTO links (link, username) VALUES ($1, $2)")
		_, err := s.db.Exec(query, p.URL, p.UserName)
		if err != nil {
			log.Fatal(err)
		}
	}
	return nil
}

func (s *StoragePostgres) Remove(p *Page) error {
	if s.IsPresent(p) == false {
		query := fmt.Sprintf("DELETE FROM links WHERE link = $1")
		_, err := s.db.Exec(query, p.URL)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		return errors.New("file does not exist")
	}
	return nil
}

func (s *StoragePostgres) IsPresent(p *Page) bool {
	var res string
	query := fmt.Sprintf("SELECT username FROM links WHERE link = $1")
	_ = s.db.Get(&res, query, p.URL)
	return res != ""
}

func (s *StoragePostgres) PickRandom(userName string) (*Page, error) {
	var link string
	var links []string

	query := fmt.Sprintf("SELECT link FROM links WHERE username = $1")
	rows, err := s.db.Query(query, userName)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		err := rows.Scan(&link)
		links = append(links, link)
		if err != nil {
			log.Fatal(err)
		}
	}

	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(len(links))

	return &Page{links[n], userName}, nil
}
