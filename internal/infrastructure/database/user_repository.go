package database

import (
	"context"
	"database/sql"
	"git.savvycom.vn/go-services/go-boilerplate/internal/domain/users/entities"
	"git.savvycom.vn/go-services/go-boilerplate/internal/domain/users/repositories"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/extra/bundebug"
	"time"
)

type (
	defaultUserRepository struct {
		db *bun.DB
	}

	// User represents a user that map 1-1 to the database table User.
	User struct {
		bun.BaseModel `bun:"user,alias:u"`
		Id            int64      `bun:"id,pk,autoincrement"`
		UserName      string     `bun:"username,notnull,nullzero"`
		Password      string     `bun:"password,notnull,nullzero"`
		Email         string     `bun:"email,notnull,nullzero"`
		Status        string     `bun:"status,notnull,nullzero,default:1"`
		CreatedAt     *time.Time `bun:"created_at,notnull,nullzero,default:current_timestamp"`
		UpdatedAt     *time.Time `bun:"updated_at,nullzero,default:current_timestamp"`
	}
)

func (d defaultUserRepository) CreateUser(ctx context.Context, user *entities.User) (int64, error) {
	//TODO implement me
	panic("implement me")
}

func (d defaultUserRepository) UpdateUser(ctx context.Context, user *entities.User) error {
	//TODO implement me
	panic("implement me")
}

func (d defaultUserRepository) DeleteUser(ctx context.Context, userID int64) error {
	//TODO implement me
	panic("implement me")
}

func (d defaultUserRepository) GetUser(ctx context.Context, userID int64) (*entities.User, error) {
	//TODO implement me
	panic("implement me")
}

var _ bun.BeforeAppendModelHook = (*User)(nil)

func (u *User) BeforeAppendModel(ctx context.Context, query bun.Query) error {
	switch query.(type) {
	case *bun.InsertQuery:
		now := time.Now()
		u.CreatedAt = &now
		u.UpdatedAt = &now
	case *bun.UpdateQuery:
		now := time.Now()
		u.UpdatedAt = &now
	}
	return nil
}

func NewUserRepository(dsn string, isDebug bool) repositories.UserRepository {
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))

	db := bun.NewDB(sqldb, pgdialect.New())
	db.AddQueryHook(bundebug.NewQueryHook(
		bundebug.WithEnabled(isDebug),
		bundebug.FromEnv(""),
	))
	return &defaultUserRepository{
		db: db,
	}
}
