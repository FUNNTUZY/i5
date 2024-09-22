package usecase

import (
	"context"
	"interactions/internal/domain"
	"interactions/internal/domain/entity"
	"interactions/internal/dto"
	"time"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

// InteractionUsecase - интерфейс, описывающий бизнес-логику работы с взаимодействиями

type InteractionUsecaseImpl struct {
	repo domain.InteractionRepository
}

// Определение структуры InteractionUsecaseImpl
func NewInteractionUsecase(repo domain.InteractionRepository) *InteractionUsecaseImpl {
	return &InteractionUsecaseImpl{
		repo: repo,
	}
}

// AddInteraction создает новое взаимодействие
func (uc *InteractionUsecaseImpl) AddInteraction(ctx context.Context, req *dto.InteractionDTO) (*dto.InteractionResponseDTO, error) {
	// Преобразуем DTO в доменную модель (entity.Interaction)
	interaction := entity.Interaction{
		ID:              uuid.NewString(), // Генерируем новый UUID для сущности
		UserID:          req.UserID,
		AdID:            req.AdID,
		SellerID:        req.SellerID,
		InteractionType: req.InteractionType,
		CreatedAt:       time.Now(),
	}

	// Добавляем взаимодействие через репозиторий
	err := uc.repo.CreateInteraction(ctx, interaction)
	if err != nil {
		log.Error().Err(err).Msg("Ошибка при добавлении взаимодействия")
		return &dto.InteractionResponseDTO{
			Success: false,
		}, err
	}

	// Возвращаем успешный результат
	return &dto.InteractionResponseDTO{
		Success: true,
	}, nil
}

func (uc *InteractionUsecaseImpl) GetInteraction(ctx context.Context, req *dto.GetInteractionRequestDTO) (*dto.GetInteractionResponseDTO, error) {
	// Получаем взаимодействия из репозитория
	interactions, err := uc.repo.GetInteractions(ctx, req.ID)
	if err != nil {
		log.Error().Err(err).Msg("Ошибка при получении взаимодействий")
		return &dto.GetInteractionResponseDTO{}, err
	}

	var types []string

	for _, interaction := range interactions {
		types = append(types, interaction.InteractionType)
	}

	return &dto.GetInteractionResponseDTO{
		Types: types,
	}, nil
}
