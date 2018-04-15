package storage

import (
	"context"
	"net/url"

	_ "github.com/lib/pq"

	sqlx "github.com/jmoiron/sqlx"

	"github.com/athega/flockflow-server/flockflow"
)

const schema = `
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS logins (
	key UUID NOT NULL DEFAULT uuid_generate_v4(),
	email text UNIQUE NOT NULL,
	timestamp timestamp NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS profiles (
	id SERIAL PRIMARY KEY,
	email text UNIQUE NOT NULL,
	name text NOT NULL DEFAULT '',
	link text NOT NULL DEFAULT '',
	phone text NOT NULL DEFAULT ''
);
`

type Store struct {
	db *sqlx.DB
}

func ConnectAndSetupSchema(dataSourceName string) (*Store, error) {
	db, err := sqlx.Connect("postgres", dataSourceName)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(16)

	if _, err := db.Exec(schema); err != nil {
		return nil, err
	}

	s := &Store{db}

	if _, err := s.LoginKey(context.Background(), "peter.hellberg@athega.se"); err != nil {
		return nil, err
	}

	return s, nil
}

func New(db *sqlx.DB) *Store {
	return &Store{db}
}

func (s *Store) Close() error {
	return s.db.Close()
}

func (s *Store) LoginKey(ctx context.Context, email string) (string, error) {
	if _, err := s.db.NamedExecContext(ctx, `INSERT INTO logins (email) 
		VALUES (:email) 
		ON CONFLICT (email) DO UPDATE
		SET key = uuid_generate_v4()`,
		&flockflow.Login{
			Email: email,
		},
	); err != nil {
		return "", err
	}

	var key string

	if err := s.db.GetContext(ctx, &key, `SELECT key FROM logins WHERE email=$1`, email); err != nil {
		return "", err
	}

	return key, nil
}

func (s *Store) ProfileID(ctx context.Context, key string) (string, error) {
	var l flockflow.Login

	if err := s.db.GetContext(ctx, &l, `SELECT * FROM logins WHERE key=$1`, key); err != nil {
		return "", flockflow.ErrInvalidLoginKey
	}

	if _, err := s.db.NamedExecContext(ctx,
		`INSERT INTO profiles (email) VALUES (:email) ON CONFLICT DO NOTHING`,
		&flockflow.Profile{Email: l.Email},
	); err != nil {
		return "", err
	}

	var id string

	if err := s.db.GetContext(ctx, &id,
		`SELECT id FROM profiles WHERE email=$1`,
		l.Email,
	); err != nil {
		return "", flockflow.ErrProfileNotFound
	}

	return id, nil
}

func (s *Store) Profile(ctx context.Context, subject string) (*flockflow.Profile, error) {
	var p flockflow.Profile

	if err := s.db.GetContext(ctx, &p, "SELECT * FROM profiles WHERE id=1"); err != nil {
		return nil, err
	}

	return &p, nil
}

func (s *Store) UpdateProfile(ctx context.Context, subject string, v url.Values) error {
	_, err := s.db.NamedExecContext(ctx,
		`UPDATE profiles SET name=:name, link=:link, phone=:phone WHERE id=:id RETURNING *`,
		&flockflow.Profile{
			ID:    subject,
			Name:  v.Get("name"),
			Link:  v.Get("link"),
			Phone: v.Get("phone"),
		},
	)

	return err
}
