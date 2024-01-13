package service

import (
	"path/filepath"

	"github.com/enuesaa/pinit/pkg/repository"
)

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
