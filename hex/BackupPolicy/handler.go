package backupPolicy

import (
	"net/http"

	"github.com/labstack/echo"
)

type BackupPolicyHandler struct {
	s BackupPolicyService
}

func NewHandler(s *BackupPolicyService) *BackupPolicyHandler {
	return &BackupPolicyHandler{s: s}
}

func (t *BackupPolicyHandler) CreateFilterAccount(awsaccount_id uint, region string) *BackupPolicyDto {
	return &BackupPolicyDto{
		Awsaccount_id: awsaccount_id,
		Region:        region,
	}
}
func (t *BackupPolicyHandler) Create(c echo.Context) error {

	var (
		request      *BackupPolicyRequest
		filter, data *BackupPolicyDto
	)
	err := c.Bind(request)
	if err == nil {
		filter = t.CreateFilterAccount(request.Awsaccount_id, request.Region)
		data.Retentionendias = request.Retentionendias
		_, err = t.s.CreateField(filter, data)
	}
	if err == nil {
		return c.JSON(http.StatusOK, "listo")
	}
	return c.JSON(http.StatusUnauthorized, err)
}

func (t *BackupPolicyHandler) Update(c echo.Context) error {
	return nil
}
func (t *BackupPolicyHandler) ListAllFor(c echo.Context) error {
	return nil
}
func (t *BackupPolicyHandler) Get(c echo.Context) error {
	return nil
}
