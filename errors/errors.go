package errors

import "errors"

var (
	ErrInvalidAdminKey          error = errors.New("hait: Invalid admin key")
	ErrInvalidItem              error = errors.New("hait: Invalid item")
	ErrInvalidItemId            error = errors.New("hait: Invalid item UUID")
	ErrInvalidItemType          error = errors.New("hait: Invalid item type")
	ErrInvalidItemTypeId        error = errors.New("hait: Invalid item type UUID")
	ErrDisallowedItemCustomName error = errors.New("hait: Invalid item custom name, try changing it to something else")
	ErrCannotReachNetwork       error = errors.New("hait: Cannot reach network")
	ErrCannotReachNode          error = errors.New("hait: Cannot reach node")
	ErrCannotReachNodeTerminal  error = errors.New("hait: Cannot reach node's terminal")
	ErrMissingPermissions       error = errors.New("hait: Missing permissions")
	ErrBanned                   error = errors.New("hait: Banned")
	ErrDatabaseKeyDoesNotExist  error = errors.New("hait: Database key does not exist")
)
