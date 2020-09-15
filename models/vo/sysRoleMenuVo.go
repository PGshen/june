package vo

type RoleAuth struct {
	RoleId  int32   `json:"roleId"`
	MenuIds []int32 `json:"menuIds"`
}
