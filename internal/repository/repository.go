package repository

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	"monitoring-app/internal/domain"
	"monitoring-app/internal/repository/db"
	"time"
)

type Repository struct {
	Queries *db.Queries
	DB      *pgxpool.Pool
}

func NewRepository(dbPool *pgxpool.Pool) *Repository {
	return &Repository{
		Queries: db.New(dbPool),
		DB:      dbPool,
	}
}

func (r *Repository) SaveResult(ctx context.Context, website domain.Website) error {
	return r.Queries.InsertResult(ctx, db.InsertResultParams{
		Url:        website.URL,
		StatusCode: int32(website.StatusCode),
		DurationMs: int32(website.Duration.Milliseconds()),
		Error:      pgxNullString(website.Error),
	})
}

func (r *Repository) GetLastResults(ctx context.Context, limit int) ([]domain.Website, error) {
	rows, err := r.Queries.GetLastResults(ctx, int32(limit))
	if err != nil {
		return nil, err
	}

	var results []domain.Website
	for _, row := range rows {
		results = append(results, domain.Website{
			URL:        row.Url,
			StatusCode: int(row.StatusCode),
			Duration:   time.Duration(row.DurationMs) * time.Millisecond,
			CheckedAt:  pgxTimestampToTime(pgtype.Timestamptz(row.CreatedAt)),
			Error:      pgxStringToError(row.Error),
		})
	}

	return results, nil
}

func (r *Repository) GetResultByURL(ctx context.Context, url string) (domain.Website, error) {
	row, err := r.Queries.GetResultByURL(ctx, url)
	if err != nil {
		return domain.Website{}, err
	}

	return domain.Website{
		URL:        row.Url,
		StatusCode: int(row.StatusCode),
		Duration:   time.Duration(row.DurationMs) * time.Millisecond,
		CheckedAt:  pgxTimestampToTime(pgtype.Timestamptz(row.CreatedAt)),
		Error:      pgxStringToError(row.Error),
	}, nil
}

func (r *Repository) DeleteOldResults(ctx context.Context) error {
	return r.Queries.DeleteOldResults(ctx)
}

func (r *Repository) CountResults(ctx context.Context) (int, error) {
	count, err := r.Queries.CountResults(ctx)
	if err != nil {
		return 0, err
	}
	return int(count), nil
}

func (r *Repository) Close() {
	r.DB.Close()
}

func pgxTimestampToTime(ts pgtype.Timestamptz) time.Time {
	if ts.Valid {
		return ts.Time
	}
	return time.Time{}
}

func pgxNullString(err error) pgtype.Text {
	if err != nil {
		return pgtype.Text{String: err.Error(), Valid: true}
	}
	return pgtype.Text{Valid: false}
}

func pgxStringToError(text pgtype.Text) error {
	if text.Valid {
		return fmt.Errorf("%s", text.String)
	}
	return nil
}
