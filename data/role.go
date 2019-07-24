package data

//Permission 定义权限
type Permission uint

//权限定义
const (
	Follow   Permission = 1
	Comment  Permission = 2
	Write    Permission = 4
	Moderate Permission = 8
	Admin    Permission = 16
)

//Role 定义用户角色权限
type Role struct {
	Model
	Name        string     `gorm:"unique;not null" json:"name"`
	Default     bool       `gorm:"default:false" json:"default"`
	Permissions Permission `gorm:"default:0"`
	Users       []User
}

func (r *Role) resetPermission() {
	r.Permissions = 0
}

func (r *Role) addPermission(perm Permission) {
	r.Permissions += perm
}

func (r *Role) removePermission(perm Permission) {
	r.Permissions -= perm
}

func (r *Role) hasPermission(perm Permission) bool {
	return r.Permissions&perm == perm
}

// func (r *Role) BeforeCreate(scope *gorm.Scope) error {
// 	uuid := uuid.NewV4()
// 	return scope.SetColumn("ID", uuid.String())
// }

var roles = map[string][]Permission{"User": {Follow, Comment, Write},
	"Moderator":     {Follow, Comment, Write, Moderate},
	"Administrator": {Follow, Comment, Write, Moderate, Admin}}

//InsertRoles 初始化角色权限
func InsertRoles() {
	Db.CreateTable(&Role{})

	defaultRole := "User"
	for key, value := range roles {
		role := Role{}
		Db.Where("name = ?", key).First(&role)
		if role.Name == "" {
			role = Role{Name: key}
			role.resetPermission()
			for _, perm := range value {
				role.addPermission(perm)
			}
			role.Default = (role.Name == defaultRole)
			Db.Create(&role)
		}
	}
}
