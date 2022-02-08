package repo

import (
	"github.com/furee/backend/infra"
	m "github.com/furee/backend/repo/master"
	"github.com/furee/backend/repo/user"
	"github.com/sirupsen/logrus"
)

type Repo struct {
	Master m.MasterRepo
	User   user.UserRepo
}

func NewRepo(db *infra.DatabaseList, logger *logrus.Logger) Repo {
	return Repo{
		Master: m.NewMasterRepo(db, logger),
		User:   user.NewMasterRepo(db, logger),
	}
}
