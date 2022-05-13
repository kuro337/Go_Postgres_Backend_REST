package models

// 1 role can have many permissions and 1 permission can also belong to multiple roles

// hence many to many
type Role struct {
	Id uint `json:"id"`
	Name string `json:"name"`
	Permissions []Permission `json:"permisssions" gorm:"many2many:role_permissions"`
}