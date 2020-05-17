package identity

import 	"github.com/qq1060656096/err"

var (
	ErrBitType18 = err.NewErr("tools.identity.bitType18", 422)
	ErrBitType18LastBit = err.NewErr("tools.identity.bitType18.lastBit", 422)
	ErrBitType15 = err.NewErr("tools.identity.bitType15", 422)
	ErrBitTypeNotFound = err.NewErr("tools.identity.bitType.notFound", 422)
	ErrBirthday = err.NewErr("tools.identity.birthday", 422)
)
