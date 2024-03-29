// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: queries.sql

package db

import (
	"context"

	"github.com/google/uuid"
)

const createDestination = `-- name: CreateDestination :one
INSERT INTO destinations (message_id, receiver) VALUES ($1, $2) RETURNING id, message_id, receiver, created_at, updated_at
`

type CreateDestinationParams struct {
	MessageID uuid.UUID
	Receiver  string
}

func (q *Queries) CreateDestination(ctx context.Context, arg CreateDestinationParams) (Destination, error) {
	row := q.db.QueryRowContext(ctx, createDestination, arg.MessageID, arg.Receiver)
	var i Destination
	err := row.Scan(
		&i.ID,
		&i.MessageID,
		&i.Receiver,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const createMessage = `-- name: CreateMessage :one
INSERT INTO messages (service, title, body) VALUES ($1, $2, $3) RETURNING id, service, title, body, created_at, updated_at
`

type CreateMessageParams struct {
	Service string
	Title   string
	Body    string
}

func (q *Queries) CreateMessage(ctx context.Context, arg CreateMessageParams) (Message, error) {
	row := q.db.QueryRowContext(ctx, createMessage, arg.Service, arg.Title, arg.Body)
	var i Message
	err := row.Scan(
		&i.ID,
		&i.Service,
		&i.Title,
		&i.Body,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
