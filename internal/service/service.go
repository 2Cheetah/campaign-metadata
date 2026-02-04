package service

import "context"

type Repository interface {
	AddTags([]string) error
	ReadTags(context.Context, string) ([]string, error)
}

type Service struct {
	Repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{
		Repo: repo,
	}
}

func (s *Service) GetCampaignTags(ctx context.Context, campaignId string) ([]string, error) {
	return s.Repo.ReadTags(ctx, campaignId)
}
