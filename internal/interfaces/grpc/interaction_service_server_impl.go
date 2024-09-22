package interfaces

import (
	"context"
	"interactions/api/proto"
	"interactions/internal/dto"
	domain "interactions/internal/interfaces"
)

type InteractionServiceServerImpl struct {
	proto.UnimplementedInteractionServiceServer
	Usecase domain.InteractionUsecase
}

// NewInteractionServiceServerImpl создает новый сервер взаимодействий
func NewInteractionServiceServerImpl(usecase domain.InteractionUsecase) proto.InteractionServiceServer {
	return &InteractionServiceServerImpl{Usecase: usecase}
}

// Преобразуем входящий proto запрос в DTO и вызываем бизнес-логику
func (s *InteractionServiceServerImpl) AddInteraction(ctx context.Context, req *proto.AddInteractionRequest) (*proto.InteractionResponse, error) {
	// Преобразуем proto в DTO
	dtoRequest := &dto.InteractionDTO{
		UserID:          req.UserId,
		AdID:            req.AdId,
		SellerID:        req.SellerId,
		InteractionType: req.Type.String(),
	}

	// Вызов бизнес-логики
	dtoResponse, err := s.Usecase.AddInteraction(ctx, dtoRequest)
	if err != nil {
		return nil, err
	}

	// Преобразуем результат из DTO в proto
	protoResponse := &proto.InteractionResponse{
		Success: dtoResponse.Success,
	}

	return protoResponse, nil
}

// Преобразуем входящий proto запрос в DTO и вызываем бизнес-логику
func (s *InteractionServiceServerImpl) GetInteraction(ctx context.Context, req *proto.GetInteractionRequest) (*proto.GetInteractionResponse, error) {
	// Преобразуем proto в DTO
	dtoRequest := &dto.GetInteractionRequestDTO{
		ID: req.AdId,
	}

	// Вызов бизнес-логики
	dtoResponse, err := s.Usecase.GetInteraction(ctx, dtoRequest)
	if err != nil {
		return nil, err
	}

	// Преобразуем результат из DTO в proto
	var protoTypes []proto.InteractionType
	for _, t := range dtoResponse.Types {
		switch t {
		case "message_sent":
			protoTypes = append(protoTypes, proto.InteractionType_message_sent)
		case "phone_revealed":
			protoTypes = append(protoTypes, proto.InteractionType_phone_revealed)
		}
	}

	protoResponse := &proto.GetInteractionResponse{
		Type: protoTypes,
	}

	return protoResponse, nil
}
