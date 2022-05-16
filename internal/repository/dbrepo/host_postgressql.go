package dbrepo

import (
	"context"
	"log"
	"time"

	"github.com/techarm/go-vigilate/internal/models"
)

func (m *postgresDBRepo) InsertHost(h models.Hosts) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := `insert into hosts(host_name, canonical_name, url, ip, ipv6, location, os, active, created_at, updated_at)
			 values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) returning id`

	var newID int
	err := m.DB.QueryRowContext(ctx, stmt,
		h.HostName,
		h.CanonicalHost,
		h.URL,
		h.IP,
		h.IPv6,
		h.Location,
		h.OS,
		h.Active,
		time.Now(),
		time.Now(),
	).Scan(&newID)

	if err != nil {
		log.Println(err)
		return newID, err
	}

	return newID, nil
}
