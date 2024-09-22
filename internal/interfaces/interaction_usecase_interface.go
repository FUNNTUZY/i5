package domain

import (
	"context"
	"interactions/internal/dto"
)

type InteractionUsecase interface {
	AddInteraction(ctx context.Context, req *dto.InteractionDTO) (*dto.InteractionResponseDTO, error)
	GetInteraction(ctx context.Context, req *dto.GetInteractionRequestDTO) (*dto.GetInteractionResponseDTO, error)
}
