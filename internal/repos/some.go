package repos

import (
	"database/sql"
	"github.com/sabahtalateh/gic"
	"github.com/sabahtalateh/gicex/internal/system"
)

func init() {
	gic.Add[*SomeRepo](
		gic.WithInit(func() *SomeRepo {
			return &SomeRepo{db: gic.MustGet[*system.System]().DB}
		}),
	)
}

type SomeRepo struct {
	db *sql.DB
}

func (s *SomeRepo) GetSome() (string, error) {
	rows, err := s.db.Query("SELECT 'SOME'")
	if err != nil {
		return "", err
	}
	defer rows.Close()

	for rows.Next() {
		var some string
		err = rows.Scan(&some)
		if err != nil {
			return "", err
		}
		return some, nil
	}

	return "", nil
}
