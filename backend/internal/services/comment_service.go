package service

import (
	"context"
	"time"

	"github.com/lanxre/kyokusulib/internal/models/db"
	"github.com/lanxre/kyokusulib/internal/models/dto"
	"github.com/lanxre/kyokusulib/internal/repository"
)

type CommentService struct {
	repo repository.CommentsRepository
}

func NewCommentService(repo *repository.CommentsRepository) *CommentService {
	return &CommentService{repo: *repo}
}

func (s *CommentService) GetNovelaComments(ctx context.Context, novelaID, userID int) ([]*dto.CommentResponse, error) {
	raw, err := s.repo.GetCommentsByNovelaID(ctx, novelaID, userID)
	if err != nil {
		return nil, err
	}

	nodes := make(map[int]*dto.CommentResponse)
	for _, c := range raw {
		nodes[c.ID] = &dto.CommentResponse{
			ID:        c.ID,
			ParentID:  c.ParentID,
			Content:   c.Content,
			CreatedAt: c.CreatedAt.Format(time.RFC3339),
			LikeCount: c.LikeCount,
			HasLike:   c.HasLike,
			HasReport: c.HasReport,
			User:      dto.CommentUserAuthor{ID: c.UserID, Name: c.UserName, Picture: c.UserImage},
			Replies:   []*dto.CommentResponse{},
		}
	}

	var root []*dto.CommentResponse
	for _, c := range raw {
		node := nodes[c.ID]
		if c.ParentID == nil {
			root = append(root, node)
		} else {
			if parent, ok := nodes[*c.ParentID]; ok {
				parent.Replies = append(parent.Replies, node)
			}
		}
	}

	if root == nil {
		return []*dto.CommentResponse{}, nil
	}

	return root, nil
}

func (s *CommentService) CreateComment(ctx context.Context, userID int, req *dto.CreateCommentRequest) (int, error) {
	return s.repo.CreateComment(ctx, userID, req)
}

func (s *CommentService) DeleteComment(ctx context.Context, commentID, userID int) error {
	return s.repo.DeleteComment(ctx, commentID, userID)
}

func (s *CommentService) UpdateComment(ctx context.Context, commentID, userID int, req *dto.UpdateCommentRequest) error {
	return s.repo.UpdateComment(ctx, commentID, userID, req)
}

func (s *CommentService) SetCommentLike(ctx context.Context, commentID, userID int) error {
	return s.repo.SetCommentLike(ctx, commentID, userID)
}

func (s *CommentService) DeleteCommentLike(ctx context.Context, commentID, userID int) error {
	return s.repo.DeleteCommentLike(ctx, commentID, userID)
}

func (s *CommentService) CreateCommentReport(ctx context.Context, commentID, userID int, reason string) error {
	return s.repo.CreateCommentReport(ctx, commentID, userID, reason)
}

func (s *CommentService) GetCommentsByUserID(ctx context.Context, userID int) ([]dto.ProfileUserCommentResponse, error) {
	userComments, err := s.repo.GetCommentsByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}
	converted := s.mapToProfileUserCommentResponse(userComments)
	return converted, nil
}

func (s *CommentService) mapToProfileUserCommentResponse(userComments []db.SelectNovelaComment) []dto.ProfileUserCommentResponse {
	res := make([]dto.ProfileUserCommentResponse, 0)
	for _, c := range userComments {
		res = append(res, dto.ProfileUserCommentResponse{
			ID: c.ID,
			NovelaID: c.NovelaID,
			NovelaTitle: c.NovelaTitle,
			Content: c.Content,
			CreatedAt: c.CreatedAt.Format(time.RFC3339),
			UpdatedAt: c.UpdatedAt.Format(time.RFC3339),
		})
	}

	return res
}