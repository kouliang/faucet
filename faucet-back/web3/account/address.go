package account

import (
	"errors"
	"reflect"

	"github.com/ethereum/go-ethereum/common"
)

func IsZeroAddress(iaddress interface{}) bool {
	var address common.Address
	switch v := iaddress.(type) {
	case string:
		address = common.HexToAddress(v)
	case common.Address:
		address = v
	default:
		return false
	}

	zeroAddressBytes := common.FromHex("0x0000000000000000000000000000000000000000")
	addressBytes := address.Bytes()
	return reflect.DeepEqual(addressBytes, zeroAddressBytes)
}

func IsAvailableAddress(hexString string) bool {
	hexCheck := common.IsHexAddress(hexString)
	isZero := IsZeroAddress(hexString)
	return hexCheck && !isZero
}

func CommonAddress(hexString string) (common.Address, error) {
	result := common.HexToAddress(hexString)
	if !IsAvailableAddress(hexString) {
		return result, errors.New("unAvailable")
	}
	return result, nil
}
