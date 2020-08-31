package model

const (
	CVarTypeAll = iota
	CVarTypeGame
	CVarTypeMode
	CVarTypeServer
)

type CVar struct {
	ID         int
	Key        string
	Value      string
	Type       int
	TypeArea   int32
	NeedUpdate bool
}

// TableName return table name
func (*CVar) TableName() string {
	return "np_cvar"
}

func (c *CVar) IsType(gameID int32, modID int32, serverID int32) bool {
	switch c.Type {
	case CVarTypeAll:
		return true
	case CVarTypeGame:
		return c.TypeArea == gameID
	case CVarTypeMode:
		return c.TypeArea == modID
	case CVarTypeServer:
		return c.TypeArea == serverID
	default:
		return false
	}
}
