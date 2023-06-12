package artemis

import (
	"github.com/theogee/artemis-core/pkg/config"
	"github.com/theogee/artemis-core/pkg/connection/cache"
	"github.com/theogee/artemis-core/pkg/connection/database"
)

type (
	ArtemisRepo struct {
		cfg   *config.Config
		db    *database.Database
		cache *cache.Cache
	}
)

func NewRepo(cfg *config.Config, db *database.Database, c *cache.Cache) *ArtemisRepo {
	return &ArtemisRepo{
		cfg:   cfg,
		db:    db,
		cache: c,
	}
}
