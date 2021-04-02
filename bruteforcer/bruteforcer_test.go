package bruteforcer

import (
	"reflect"
	"testing"

	"github.com/afjoseph/commongo/print"
	"github.com/stretchr/testify/require"
)

func Test_Functional(t *testing.T) {
	expectedBuff := []uint8{0x00, 20, 30}
	actualBuff := []uint8{0x00, 0xAA, 0xAA}
	Run(actualBuff, 1, func(buff []uint8) bool {
		print.Infof("%+v\n", buff)
		// XXX This is quite costly, but that's not an issue
		if reflect.DeepEqual(actualBuff, expectedBuff) {
			return false
		}
		return true
	})
	require.Equal(t, expectedBuff, actualBuff)
}
