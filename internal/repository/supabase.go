package repository

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

type SupabaseDB struct {
	pool *pgxpool.Pool
}

func NewSupabaseDB(ctx context.Context) (*SupabaseDB, error) {
	pool, err := pgxpool.New(ctx, os.Getenv("DB_CONNECTION_STRING"))
	if err != nil {
		return nil, err
	}
	if err := pool.Ping(ctx); err != nil {
		pool.Close()
		return nil, err
	}
	return &SupabaseDB{
		pool: pool,
	}, nil
}

func (sDB *SupabaseDB) Close() {
	sDB.pool.Close()
}

func (sDB *SupabaseDB) AddTags(tags []string) error {
	return nil
}

func (sDB *SupabaseDB) ReadTags(ctx context.Context, campaign string) ([]string, error) {
	q := `
		SELECT tag
		FROM public.campaign_tags
		WHERE campaign = $1
	`
	rows, err := sDB.pool.Query(ctx, q, campaign)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tags []string
	for rows.Next() {
		var tag string
		if err := rows.Scan(&tag); err != nil {
			return nil, err
		}
		tags = append(tags, tag)
	}

	return tags, nil
}
