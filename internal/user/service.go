package user

import "github.com/go-playground/validator/v10"

type Service interface {
	Save(u *UserGrade) error
	FindByKey(k string) (*UserGrade, error)
}

type Serv struct {
	repo      Repository
	validator *validator.Validate
}

func NewService(r Repository) *Serv {
	v := validator.New()
	return &Serv{repo: r, validator: v}
}

func (s *Serv) Save(u *UserGrade) error {
	err := s.validator.Struct(u)
	if err != nil {
		return err
	}

	user, err := s.repo.GetByKey(u.UserId)
	if err == nil {
		updateUser(u, user)
	}
	s.repo.Store(u.UserId, u)

	return nil
}

func (s *Serv) FindByKey(k string) (*UserGrade, error) {
	return s.repo.GetByKey(k)
}

func updateUser(fromUser *UserGrade, toUser *UserGrade) {
	if fromUser.Spp != nil {
		toUser.Spp = fromUser.Spp
	}

	if fromUser.ShippingFee != nil {
		toUser.ShippingFee = fromUser.ShippingFee
	}

	if fromUser.ReturnFee != nil {
		toUser.ReturnFee = fromUser.ReturnFee
	}
}
