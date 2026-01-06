package dto

type CreateNovelaDTO struct {
	Title             string `validate:"required,min=2,max=255"`
	Description       string `validate:"required,max=5000"`
	Type              string `validate:"required,oneof=Ранобэ Веб-новелла Оригинал"`
	AgeRating         string `validate:"required,oneof=0+ 6+ 12+ 16+ 18+"`
	ReleaseYear       string `validate:"required,len=4,numeric"`
	Status            string `validate:"required"`
	TranslationStatus string `validate:"required"`
	
	Genres            []string `validate:"required,min=1"`
	Categories        []string `validate:"required,min=1"`
}