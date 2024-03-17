package usecase

import (
	"context"
	"fmt"

	"go-clean/internal/entity"
)

// TranslationUseCase -.
type TranslationUseCase struct {
	repo TranslationRepo
}

// New -.
func New(r TranslationRepo) *TranslationUseCase {
	return &TranslationUseCase{
		repo: r,
	}
}

// History - getting translate history from store.
func (uc *TranslationUseCase) History(ctx context.Context) ([]entity.Translation, error) {
	translations, err := uc.repo.GetHistory(ctx)
	if err != nil {
		return nil, fmt.Errorf("TranslationUseCase - History - s.repo.GetHistory: %w", err)
	}

	return translations, nil
}
