package backupPolicy

import (
	"errors"

	"github.com/jinzhu/gorm"
)

type BackupPolicyRepository struct {
	client *gorm.DB
}

const (
	ModelErrRecordNotFound  string = "no existe el registro"
	ModelErrRecordNotCreate string = "Registro ya existe no se puede crear nuevamente"
)

func (t *BackupPolicyRepository) SetDriver(d Driver) error {
	t.client = d
	return nil
}
func (t *BackupPolicyRepository) CreateField(filter interface{}, v interface{}) error {
	err := t.GetField(filter)
	if err == nil {
		return errors.New(ModelErrRecordNotCreate)
	}
	if gorm.ErrRecordNotFound == err {
		if err := t.client.Create(v).Error; err != nil {
			return err
		}
	}
	return err

}
func (t *BackupPolicyRepository) GetField(v interface{}) error {
	result := t.client.Where(v).First(v)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected <= 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
