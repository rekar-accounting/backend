package repositories
type  UserRepository interface{
	Create(models.User) (models.User, error)
	Read(int) (models.User, error)
	Update(models.User, int) (models.User, error)
	Delete(int) error

} 