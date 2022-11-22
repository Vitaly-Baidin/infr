package user

import "github.com/Vitaly-Baidin/infr/pkg/storage"

type RepoStorage struct {
	storage *storage.Storage[string, UserGrade]
}

func NewRepoStorage(s *storage.Storage[string, UserGrade]) *RepoStorage {
	return &RepoStorage{storage: s}
}

func (r *RepoStorage) Store(k string, u *UserGrade) {
	r.storage.Set(k, u)
}

func (r *RepoStorage) GetByKey(k string) (*UserGrade, error) {
	result, err := r.storage.Get(k)
	if err != nil {
		return nil, err
	}

	return result, nil
}
