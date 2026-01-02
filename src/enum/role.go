package enum

type Role string

const (
	RoleAdmin          Role = "ADMIN"
	RoleDesigner       Role = "DESIGNER"
	RoleSeller         Role = "SELLER"
	RoleChiefInstaller Role = "CHIEF_INSTALLER"
	RoleInstaller      Role = "INSTALLER"
	RoleClient         Role = "CLIENT"
)

func (r Role) String() string {
	return string(r)
}

func IsValidRole(role string) bool {
	switch Role(role) {
	case RoleAdmin, RoleDesigner, RoleSeller, RoleChiefInstaller, RoleInstaller, RoleClient:
		return true
	}
	return false
}
