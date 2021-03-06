package addrutils

import (
	"errors"
	"github.com/antimoth/logger"
)

var utilLogger = logger.NewLogger("INFO", "addrutils")

var (
	ErrHash160Size        = errors.New("decoded address is of unknown size")
	ErrUnknownAddrPrefix  = errors.New("unknown addr prefix")
	ErrEncodeSizeOverflow = errors.New("encoded size out of valid range")
	ErrBchEncodePadding   = errors.New("encoding padding error")
	ErrEthNetID           = errors.New("error eth net id")
	ErrUnknownChainCode   = errors.New("unknown chain code")
	ErrUnknownNetID       = errors.New("unknown netID")
)
