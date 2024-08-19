package interfaces

type UserID string

type FinePermission uint32
type GlobalPermission uint32

const (
	PERM_NOTHING_FINE   = FinePermission(0)
	PERM_NOTHING_GLOBAL = GlobalPermission(0)
	PERM_MANAGE_SUBNET  = FinePermission(1 << (iota - 1))
	PERM_CREATE_SUBNETS = GlobalPermission(1 << (iota - 1))
	PERM_BAN_USERS
	PERM_CHANGE_PERMS
	PERM_DISABLE_NODES
	PERM_CREATE_ADMIN_KEYS
	PERM_SEE_USERS
	PERM_CONNECT_TERM = FinePermission(1 << (iota - 1))
)

type User interface {
	GetID() UserID
	GetGlobalPermissions() GlobalPermission
	GetBanned() (bool, error)
	GetStatus() (string, error)
	GetNodes() ([]Node, error)
}
