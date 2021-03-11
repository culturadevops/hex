package backupPolicy

type BackupPolicyService struct {
	r BackupPolicyRepository
}

func (t *BackupPolicyService) NewService(r Repository) Service {
	return &UserService{
		r,
	}
}
func (t *BackupPolicyService) ListAllFor(v interface{}, fields interface{}) error {
	return nil
}
func (t *BackupPolicyService) GetField(v interface{}) error {
	return nil
}
func (t *BackupPolicyService) CreateField(filter interface{}, v interface{}) error {
	return t.r.CreateField(filter, v)
}
func (t *BackupPolicyService) Get(v interface{}, where string) error {
	return nil
}
func (t *BackupPolicyService) Create(v interface{}, where string) error {
	return nil
}
func (t *BackupPolicyService) Update(filter interface{}, update interface{}) error {
	return nil
}
