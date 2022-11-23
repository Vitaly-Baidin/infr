package user

type Service interface {
	Save(u *UserGrade) error
	FindByKey(k string) (*UserGrade, error)
}

type Serv struct {
	repo Repository
}

func NewService(r Repository) *Serv {
	return &Serv{repo: r}
}

func (s *Serv) Save(u *UserGrade) error {
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
