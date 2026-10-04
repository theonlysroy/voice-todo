package todo

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/theonlysroy/voice-todo/go-backend/internal/db"
)

var ErrNotFound = errors.New("todo not found")

type Service struct {
	q *db.Queries
}

func NewService(q *db.Queries) *Service {
	return &Service{q: q}
}

func (s *Service) Create(ctx context.Context, title string) (db.Todo, error) {
	return s.q.CreateTodo(ctx, title)
}

func (s *Service) Get(ctx context.Context, id int64) (db.Todo, error) {
	t, err := s.q.GetTodo(ctx, id)
	if errors.Is(err, pgx.ErrNoRows) {
		return db.Todo{}, ErrNotFound
	}
	return t, err
}

func (s *Service) List(ctx context.Context, limit, offset int32) ([]db.Todo, error) {
	return s.q.ListTodos(ctx, db.ListTodosParams{Limit: limit, Offset: offset})
}

func (s *Service) Update(ctx context.Context, id int64, title string, done bool) (db.Todo, error) {
	t, err := s.q.UpdateTodo(ctx, db.UpdateTodoParams{ID: id, Title: title, IsCompleted: done})
	if errors.Is(err, pgx.ErrNoRows) {
		return db.Todo{}, ErrNotFound
	}
	return t, err
}

func (s *Service) Delete(ctx context.Context, id int64) error {
	_, err := s.q.DeleteTodo(ctx, id)
	return err
}
