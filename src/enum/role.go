package enum

type RoleEnum string

const (
	RoleAdmin          RoleEnum = "ADMIN"
	RoleDesigner       RoleEnum = "DESIGNER"
	RoleSeller         RoleEnum = "SELLER"
	RoleChiefInstaller RoleEnum = "CHIEF_INSTALLER"
	RoleInstaller      RoleEnum = "INSTALLER"
)

func (r RoleEnum) String() string {
	return string(r)
}

func IsValidRole(role string) bool {
	switch RoleEnum(role) {
	case RoleAdmin, RoleDesigner, RoleSeller, RoleChiefInstaller, RoleInstaller:
		return true
	}
	return false
}
