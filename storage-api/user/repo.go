package user

type Repository interface {
	Store(k string, u *UserGrade)
	GetByKey(k string) (*UserGrade, error)
}
