package helpers

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"gitlab.ozon.dev/qa/classroom-4/act-device-api/internal/config"
	"gitlab.ozon.dev/qa/classroom-4/act-device-api/internal/database"
	"testing"
)

func NewDB(t *testing.T) {
	if err := config.ReadConfigYML("config.yml"); err != nil {
		log.Fatal().Err(err).Msg("Failed init configuration")
	}
	cfg := config.GetConfigInstance()

	dsn := fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=%v",
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.Name,
		cfg.Database.SslMode,
	)

	db, err := database.NewPostgres(dsn, cfg.Database.Driver)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed init postgres")
	}
	defer db.Close()
}
