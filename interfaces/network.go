package interfaces

import (
	"github.com/google/uuid"
)

type ItemType interface {
	GetName() string
	GetID() uuid.UUID
}

type Item interface {
	GetID() uuid.UUID
	GetType() ItemType
	GetCustomName() string
}

type AdminKey string

type Node interface {
	GetUserID() (UserID, error)
	GetTerm() (Terminal, error)
	GetPermissions() ([]FinePermission, error)
}

type SubnetID string

type Post interface {
	GetTitle() (string, error)
	GetBriefDescription() (string, error)
	GetFullPost() (string, error)
}

type Network interface {
	GetUsers() []*User
	Connect(SubnetID) error
	Disconnect(SubnetID) error
	CreateSubnet(SubnetID) error
	DeleteSubnet(SubnetID) error
	InviteUser(SubnetID, UserID) error
	GetSubnets() ([]SubnetID, error)
	GetSubnetPermissions() (map[SubnetID]FinePermission, error)
	ValidateItem(Item) (bool, error)
	RegisterItem(Item, AdminKey) error
	RegisterItemType(ItemType, AdminKey) error
	CreateAdminKey() (AdminKey, error)
	GetPosts() ([]Post, error)
	GetMOTD() (string, error)
}
