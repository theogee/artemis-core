package artemis

import (
	artemis "github.com/theogee/artemis-core/internal/repo"
	"github.com/theogee/artemis-core/pkg/config"
)

type (
	ArtemisUsecase struct {
		cfg         *config.Config
		artemisRepo *artemis.ArtemisRepo
	}
)

func NewUsecase(cfg *config.Config, artemisRepo *artemis.ArtemisRepo) *ArtemisUsecase {
	return &ArtemisUsecase{
		cfg:         cfg,
		artemisRepo: artemisRepo,
	}
}
