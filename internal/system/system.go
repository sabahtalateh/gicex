package system

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/lib/pq"

	"github.com/sabahtalateh/gic"
	"github.com/sabahtalateh/gicex/internal/config"
)

func init() {
	gic.Add[*System](
		gic.WithInitE(func() (*System, error) {
			dbConf := gic.Get[config.Config]().DB
			psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
				dbConf.Host,
				dbConf.Port,
				dbConf.User,
				dbConf.Password,
				dbConf.DBName,
			)
			db, err := sql.Open("postgres", psqlconn)
			if err != nil {
				return nil, err
			}

			return &System{DB: db}, nil
		}),
		gic.WithStart(func(ctx context.Context, s *System) error {
			errC := make(chan error)
			go func() {
				if err := s.DB.Ping(); err != nil {
					errC <- errors.Join(fmt.Errorf("can not start database"), err)
				}
				errC <- nil
			}()

			select {
			case <-ctx.Done():
				return errors.Join(fmt.Errorf("not started"), ctx.Err())
			case err := <-errC:
				return err
			}
		}),
		gic.WithStop(func(ctx context.Context, s *System) error {
			errC := make(chan error)
			go func() {
				if err := s.DB.Close(); err != nil {
					errC <- errors.Join(fmt.Errorf("database not stopped"), err)
				}
				errC <- nil
			}()

			select {
			case <-ctx.Done():
				return errors.Join(fmt.Errorf("not stopped"), ctx.Err())
			case err := <-errC:
				return err
			}
		}),
	)
}

type System struct {
	DB *sql.DB
}
