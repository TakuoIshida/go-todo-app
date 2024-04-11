// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package tenant

const TableNameTodo = "todos"

// Todo mapped from table <todos>
type Todo struct {
	CreateUserID string `gorm:"column:create_user_id;not null" json:"create_user_id"`
	UpdateUserID string `gorm:"column:update_user_id;not null" json:"update_user_id"`
	TenantID     string `gorm:"column:tenant_id;not null" json:"tenant_id"`
	ID           string `gorm:"column:id;primaryKey" json:"id"`
	Title        string `gorm:"column:title;not null" json:"title"`
	Description  string `gorm:"column:description" json:"description"`
	IsDeleted    bool   `gorm:"column:is_deleted" json:"is_deleted"`
}

// TableName Todo's table name
func (*Todo) TableName() string {
	return TableNameTodo
}
