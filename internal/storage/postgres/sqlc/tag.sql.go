// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: tag.sql

package sqlc

import (
	"context"
	"time"
)

const createTag = `-- name: CreateTag :exec
INSERT INTO tags (id, title, user_id, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5)
`

type CreateTagParams struct {
	ID        string    `db:"id"`
	Title     string    `db:"title"`
	UserID    string    `db:"user_id"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

func (q *Queries) CreateTag(ctx context.Context, arg CreateTagParams) error {
	_, err := q.db.Exec(ctx, createTag,
		arg.ID,
		arg.Title,
		arg.UserID,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
	return err
}

const getTagIDByTitle = `-- name: GetTagIDByTitle :one
SELECT id
FROM tags
WHERE title = $1
  AND user_id = $2
  AND deleted_at IS NULL
`

type GetTagIDByTitleParams struct {
	Title  string `db:"title"`
	UserID string `db:"user_id"`
}

func (q *Queries) GetTagIDByTitle(ctx context.Context, arg GetTagIDByTitleParams) (string, error) {
	row := q.db.QueryRow(ctx, getTagIDByTitle, arg.Title, arg.UserID)
	var id string
	err := row.Scan(&id)
	return id, err
}

const getTagsByTaskID = `-- name: GetTagsByTaskID :many
SELECT tags.id, tags.title, tags.updated_at
FROM tags
    JOIN tasks_tags
        ON tags.id = tasks_tags.tag_id
WHERE tasks_tags.task_id = $1
  AND tags.deleted_at IS NULL
`

type GetTagsByTaskIDRow struct {
	ID        string    `db:"id"`
	Title     string    `db:"title"`
	UpdatedAt time.Time `db:"updated_at"`
}

func (q *Queries) GetTagsByTaskID(ctx context.Context, taskID string) ([]GetTagsByTaskIDRow, error) {
	rows, err := q.db.Query(ctx, getTagsByTaskID, taskID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetTagsByTaskIDRow{}
	for rows.Next() {
		var i GetTagsByTaskIDRow
		if err := rows.Scan(&i.ID, &i.Title, &i.UpdatedAt); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getTagsByUserID = `-- name: GetTagsByUserID :many
SELECT id, title, updated_at
FROM tags
WHERE user_id = $1
  AND deleted_at IS NULL
`

type GetTagsByUserIDRow struct {
	ID        string    `db:"id"`
	Title     string    `db:"title"`
	UpdatedAt time.Time `db:"updated_at"`
}

func (q *Queries) GetTagsByUserID(ctx context.Context, userID string) ([]GetTagsByUserIDRow, error) {
	rows, err := q.db.Query(ctx, getTagsByUserID, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetTagsByUserIDRow{}
	for rows.Next() {
		var i GetTagsByUserIDRow
		if err := rows.Scan(&i.ID, &i.Title, &i.UpdatedAt); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const linkTagToTask = `-- name: LinkTagToTask :exec
INSERT INTO tasks_tags (task_id, tag_id)
VALUES ($1, (SELECT id
             FROM tags
             WHERE title = $2
             AND user_id = $3)
)
`

type LinkTagToTaskParams struct {
	TaskID string `db:"task_id"`
	Title  string `db:"title"`
	UserID string `db:"user_id"`
}

func (q *Queries) LinkTagToTask(ctx context.Context, arg LinkTagToTaskParams) error {
	_, err := q.db.Exec(ctx, linkTagToTask, arg.TaskID, arg.Title, arg.UserID)
	return err
}

const unlinkTagFromTask = `-- name: UnlinkTagFromTask :exec
DELETE FROM tasks_tags
WHERE task_id = $1
  AND tag_id = (SELECT id
                FROM tags
                WHERE title = $2
                AND user_id = $3
)
`

type UnlinkTagFromTaskParams struct {
	TaskID string `db:"task_id"`
	Title  string `db:"title"`
	UserID string `db:"user_id"`
}

func (q *Queries) UnlinkTagFromTask(ctx context.Context, arg UnlinkTagFromTaskParams) error {
	_, err := q.db.Exec(ctx, unlinkTagFromTask, arg.TaskID, arg.Title, arg.UserID)
	return err
}
