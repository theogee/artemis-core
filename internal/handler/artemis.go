package artemis

import (
	artemis "github.com/theogee/artemis-core/internal/usecase"
	"github.com/theogee/artemis-core/pkg/config"
)

type (
	ArtemisHandler struct {
		cfg            *config.Config
		artemisUsecase *artemis.ArtemisUsecase
	}
)

func NewHandler(cfg *config.Config, artemisUsecase *artemis.ArtemisUsecase) *ArtemisHandler {
	return &ArtemisHandler{
		cfg:            cfg,
		artemisUsecase: artemisUsecase,
	}
}
