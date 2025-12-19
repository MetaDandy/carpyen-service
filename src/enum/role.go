package enum

type RoleEnum string

const (
	RoleAdmin          RoleEnum = "ADMIN"
	RoleDesigner       RoleEnum = "DESIGNER"
	RoleSeller         RoleEnum = "SELLER"
	RoleChiefInstaller RoleEnum = "CHIEF_INSTALLER"
	RoleInstaller      RoleEnum = "INSTALLER"
)
