package comment

import (
	"context"
	"errors"
	"fmt"
)

var (
	ErrFetchComment   = errors.New("failed to fetch comment by id")
	ErrNotImplemented = errors.New("not implemented")
)

// Comment - a representation of the Comment
// structure for our service
type Comment struct {
	ID     string
	Slug   string
	Body   string
	Author string
}

// Store - this interface defines all of the methods
// that our service needs to operate
type Store interface {
	GetComment(context.Context, string) (Comment, error)
	PostComment(context.Context, Comment) (Comment, error)
	UpdateComment(context.Context, string, Comment) (Comment, error)
	DeleteComment(context.Context, string) error
}

// Service - struct containing app logic
type Service struct {
	Store Store
}

// NewService - returns a pointer to a new service
func NewService(store Store) *Service {
	return &Service{
		Store: store,
	}
}

func (s *Service) GetComment(ctx context.Context, id string) (Comment, error) {
	fmt.Println("retrieving comment")
	cmt, err := s.Store.GetComment(ctx, id)
	if err != nil {
		fmt.Println(err)
		return Comment{}, ErrFetchComment
	}
	return cmt, nil
}

func (s *Service) UpdateComment(ctx context.Context, ID string, updatedCmt Comment) (Comment, error) {
	cmt, err := s.Store.UpdateComment(ctx, ID, updatedCmt)
	if err != nil {
		fmt.Println("error updating comment")
		return Comment{}, err
	}
	return cmt, nil
}

func (s *Service) PostComment(ctx context.Context, cmt Comment) (Comment, error) {

	insertedCmt, err := s.Store.PostComment(ctx, cmt)
	if err != nil {
		return Comment{}, err
	}
	return insertedCmt, nil
}

func (s *Service) DeleteComment(ctx context.Context, id string) error {
	return s.Store.DeleteComment(ctx, id)
}
