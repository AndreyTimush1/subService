package repository

import (
    "context"
    "subscriptions-service/internal/models"

    "github.com/google/uuid"
    "github.com/jackc/pgx/v5/pgxpool"
)

type SubscriptionRepository struct {
    db *pgxpool.Pool
}

func NewSubscriptionRepository(db *pgxpool.Pool) *SubscriptionRepository {
    return &SubscriptionRepository{db: db}
}

func (r *SubscriptionRepository) Create(ctx context.Context, s models.Subscription) error {
    _, err := r.db.Exec(ctx,
        `INSERT INTO subscriptions (id, service_name, price, user_id, start_date, end_date)
         VALUES ($1, $2, $3, $4, $5, $6)`,
        s.ID, s.ServiceName, s.Price, s.UserID, s.StartDate, s.EndDate)
    return err
}

func (r *SubscriptionRepository) GetByID(ctx context.Context, id uuid.UUID) (*models.Subscription, error) {
    row := r.db.QueryRow(ctx,
        `SELECT id, service_name, price, user_id, start_date, end_date FROM subscriptions WHERE id=$1`, id)

    var s models.Subscription
    err := row.Scan(&s.ID, &s.ServiceName, &s.Price, &s.UserID, &s.StartDate, &s.EndDate)
    if err != nil {
        return nil, err
    }
    return &s, nil
}

func (r *SubscriptionRepository) Update(ctx context.Context, s models.Subscription) error {
    _, err := r.db.Exec(ctx,
        `UPDATE subscriptions SET service_name=$1, price=$2, user_id=$3, start_date=$4, end_date=$5 WHERE id=$6`,
        s.ServiceName, s.Price, s.UserID, s.StartDate, s.EndDate, s.ID)
    return err
}

func (r *SubscriptionRepository) Delete(ctx context.Context, id uuid.UUID) error {
    _, err := r.db.Exec(ctx, `DELETE FROM subscriptions WHERE id=$1`, id)
    return err
}

func (r *SubscriptionRepository) GetTotal(ctx context.Context, userID *uuid.UUID, serviceName *string) (int, error) {
    query := `SELECT COALESCE(SUM(price),0) FROM subscriptions WHERE 1=1`
    args := []interface{}{}
    idx := 1

    if userID != nil {
        query += " AND user_id=$1"
        args = append(args, *userID)
        idx++
    }
    if serviceName != nil {
        if idx == 1 {
            query += " AND service_name=$1"
        } else {
            query += " AND service_name=$2"
        }
        args = append(args, *serviceName)
    }

    row := r.db.QueryRow(ctx, query, args...)
    var total int
    if err := row.Scan(&total); err != nil {
        return 0, err
    }
    return total, nil
}
