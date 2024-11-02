package service

import (
	"time"

	"github.com/enuesaa/pinit/internal/repository"
)

type AppConfig struct {
	InternalBinderName string `dynamo:"BinderName"`
	InternalNoteName   string `dynamo:"NoteName"`
	OpenaiToken string `dynamo:"openaiToken"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
type AppConfigCreation struct {
	OpenaiToken string
}

func NewRegistrySrv(repos repository.Repos) RegistrySrv {
	return RegistrySrv{
		repos: repos,
	}
}

type RegistrySrv struct {
	repos repository.Repos
}

func (srv *RegistrySrv) Get() AppConfig {
	data := AppConfig{
		OpenaiToken: "",
	}
	if err := srv.repos.Db.Get("@app", "config", &data); err != nil {
		// return default config
		return data
	}
	return data
}

func (srv *RegistrySrv) Update(creation AppConfigCreation) error {
	config := AppConfig{
		InternalBinderName: "@app",
		InternalNoteName: "config",
		OpenaiToken: creation.OpenaiToken,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	return srv.repos.Db.Put(config)
}
