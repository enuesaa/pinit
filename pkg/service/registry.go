package service

import (
	"context"
	"path/filepath"

	q "github.com/enuesaa/pinit/pkg/ent/appconf"
	"github.com/enuesaa/pinit/pkg/repository"
)

type AppConfig struct {
	chatgpttoken string
}

func NewRegistrySrv(repos repository.Repos) RegistrySrv {
	return RegistrySrv{
		repos: repos,
	}
}

type RegistrySrv struct {
	repos repository.Repos
}

func (srv *RegistrySrv) GetPath() (string, error) {
	homedir, err := srv.repos.Fs.HomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(homedir, ".pinit"), nil
}


func (srv *RegistrySrv) IsExist() bool {
	path, err := srv.GetPath()
	if err != nil {
		return false
	}
	return srv.repos.Fs.IsExist(path)
}

func (srv *RegistrySrv) Create() error {
	if srv.IsExist() {
		return nil
	}
	path, err := srv.GetPath()
	if err != nil {
		return err
	}
	return srv.repos.Fs.CreateDir(path)
}

func (srv *RegistrySrv) DeleteBufDir() error {
	if !srv.IsExist() {
		return nil
	}
	path, err := srv.GetPath()
	if err != nil {
		return err
	}
	return srv.repos.Fs.Remove(path)
}

func (srv *RegistrySrv) GetOpenAiApiToken() string {
	record, err := srv.repos.Db.Appconf().Query().Where(q.KeyEQ("openaiApiToken")).First(context.Background())
	if err != nil {
		return ""
	}
	return record.Value
}

func (srv *RegistrySrv) SetOpenAiApiToken(token string) error {
	record, err := srv.repos.Db.Appconf().Query().Where(q.KeyEQ("openaiApiToken")).First(context.Background())
	if err != nil {
		_, err := srv.repos.Db.Appconf().Create().
			SetKey("openaiApiToken").
			SetValue(token).
			Save(context.Background())
		if err != nil {
			return err
		}
		return nil
	}

	_, err = record.Update().
		SetValue(token).
		Save(context.Background())
	if err != nil {
		return err
	}
	return nil
}
