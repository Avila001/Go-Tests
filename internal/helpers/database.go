package helpers

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"gitlab.ozon.dev/qa/classroom-4/act-device-api/internal/config"
	"gitlab.ozon.dev/qa/classroom-4/act-device-api/internal/database"
	"testing"
)

type Store struct {
	DB *sqlx.DB
}

func ConnectPostgres(dsn string) (*Store, error) {
	db, err := database.NewPostgres(dsn, "pgx")
	if err != nil {
		return nil, err
	}
	return &Store{DB: db}, nil
}

//func (r Store) CreateDeviceFromBD(ctx context.Context, deviceID int) (*model.Device, error) {
//	var (
//		event models.DeviceEvent
//	)
//	if err := db.Q
//	query := sq.Select("id", "device_id", "type", "status", "payload", "created_at", "updated_at").
//		PlaceholderFormat(sq.Dollar).
//		From("devices_events").
//		Where(sq.Eq{"device_id": deviceID}).
//		OrderBy("id DESC").
//		Limit(1)
//
//	s, args, err := query.ToSql()
//	if err != nil {
//		return nil, err
//	}
//
//	err = r.DB.GetContext(ctx, &event, s, args...)
//
//	return &event, err
//}

func DBConnection(t *testing.T) *Store {
	if err := config.ReadConfigYML("../../config.yml"); err != nil {
		t.Fatal("Failed to init configuration")
	}
	cfg := config.GetConfigInstance()
	dsn := fmt.Sprintf("postgres://%v:%v@%v:%v/%v",
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.Name)
	fmt.Println(dsn)
	db, err := ConnectPostgres(dsn)
	if err != nil {
		t.Fatal("Failed to init postgres")
	}
	return db
}

//func (r Store) ByDeviceId(ctx context.Context, deviceID int) (*models.DeviceEvent, error) {
//	var (
//		event models.DeviceEvent
//	)
//	query := sq.Select("id", "device_id", "type", "status", "payload", "created_at", "updated_at").
//		PlaceholderFormat(sq.Dollar).
//		From("devices_events").
//		Where(sq.Eq{"device_id": deviceID}).
//		OrderBy("id DESC").
//		Limit(1)
//
//	s, args, err := query.ToSql()
//	if err != nil {
//		return nil, err
//	}
//
//	err = r.DB.GetContext(ctx, &event, s, args...)
//
//	return &event, err
//}
