package backupPolicy

type BackuppolicyDto struct {
	Awsaccount_id   uint   `gorm:"type:int;not null;"`
	Region          string `gorm:"type:varchar(100);not null;"`
	Instanceid      string `gorm:"type:varchar(100);not null;"`
	Retentionendias int    `gorm:"type:int;not null;"`

	Repetition   string `gorm:"type:varchar(100);not null;"`
	Resource     string `gorm:"type:varchar(100);not null;"`
	Resourcename string `gorm:"type:varchar(100);not null;"`
	Iscritical   string `gorm:"type:varchar(100);not null;"`
	EnvId        int    `gorm:"type:int;not null;"`
	Reporte      bool   `gorm:"type:int;not null;"`
}

type BackupPolicyRequest struct {
	Awsaccount_id   uint   `json:"awsaccount_id" form:"awsaccount_id" query:"awsaccount_id"`
	Region          string `json:"region" form:"region" query:"region"`
	Instanceid      string `json:"instanceid" form:"instanceid" query:"instanceid"`
	Retentionendias int    `json:"retentionendias" form:"retentionendias" query:"retentionendias"`
	Repetition      string `json:"repetition" form:"repetition" query:"repetition"`
	Resource        string `json:"resource" form:"resource" query:"resource"`
	Resourcename    string `json:"resourcename" form:"resourcename" query:"resourcename"`
	Iscritical      string `json:"iscritical" form:"iscritical" query:"iscritical"`
	Reporte         bool   `json:"reporte" form:"reporte" query:"reporte"`
	EnvId           int    `json:"envid" form:"envid" query:"envid"`
}
