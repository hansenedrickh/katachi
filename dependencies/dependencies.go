package dependencies

import (
	"github.com/hansenedrickh/katachi/config"
	"github.com/hansenedrickh/katachi/pkg/ping"
	"github.com/hansenedrickh/katachi/pkg/sample"
	"github.com/hansenedrickh/katachi/pkg/user"
)

type Dependencies struct {
	PingUsecase      *ping.Usecase
	SampleUsecase    *sample.Usecase
	SampleRepository *sample.Repository
	UserRepository   *user.Repository
	UserUsecase      *user.Usecase
}

func Setup(deps *Dependencies, cfg config.Config) error {
	db, err := SetupDB(cfg.Database)
	if err != nil {
		return err
	}

	deps.PingUsecase = ping.NewUsecase()

	deps.SampleRepository = sample.NewRepository(db)
	deps.SampleUsecase = sample.NewUsecase(deps.SampleRepository)

	deps.UserRepository = user.NewRepository(db)
	deps.UserUsecase = user.NewUsecase(cfg.JWTSecretToken, deps.UserRepository)

	return nil
}
