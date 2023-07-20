package helpers

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"gitlab.ozon.dev/qa/classroom-4/act-device-api/internal/config"
	"gitlab.ozon.dev/qa/classroom-4/act-device-api/internal/database"
	"testing"
)

type DeviceRepository struct {
	store *Store
	//FindByEmail(string) (*model.User, error)
}

type Storage struct {
	DB *sqlx.DB
}

func NewPostgres(dsn string) (*Storage, error) {
	db, err := database.NewPostgres(dsn, "pgx")
	if err != nil {
		return nil, err
	}
	return &Storage{DB: db}, nil
}

func GetDBConnection(t *testing.T) *Storage {
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
	db, err := NewPostgres(dsn)
	if err != nil {
		t.Fatal("Failed to init postgres")
	}
	return db
}

//// Store ...
//type Store struct {
//	db               *sql.DB
//	deviceRepository *DeviceRepository
//}
//
//// New ...
//func New(db *sql.DB) *Store {
//	return &Store{
//		db: db,
//	}
//}

//func (r *DeviceRepository) FindDevice(id int) (*model.Device, error) {
//	u := &model.Device{}
//	if err := r.store.db.QueryRow(
//		"SELECT id, platform, user_id, entered_id, removed, created_at, updated_at FROM devices WHERE device_id = $1",
//		id,
//	).Scan(
//		&u.id,
//		&u.platform,
//		&u.user_id,
//		&u.entered_id,
//		&u.Removed
//	); err != nil {
//		if err == sql.ErrNoRows {
//			return nil, store.ErrRecordNotFound
//		}
//
//		return nil, err
//	}
//
//	return u, nil
//}

// //////=============================================================================
//func ByDeviceId(ctx context.Context, deviceID int) (*models.DeviceEvent, error) {
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
