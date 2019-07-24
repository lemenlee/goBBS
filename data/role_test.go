package data

import (
	"testing"
)

func TestInserRole(t *testing.T) {
	Db.Delete(&Role{})
	InsertRoles()
	rolesTest := []*Role{}
	Db.Find(&rolesTest)
	// fmt.Println(rolesTest[0])
	if len(rolesTest) != len(roles) {
		t.Error("insert roles error")
	}
}
