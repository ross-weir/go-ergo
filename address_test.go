package ergo

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddress_Base58_Mainnet(t *testing.T) {
	addr, _ := NewAddress("9hdxkYakTHWXR992umPcvh8bAEGG9Sdoi7uW8TKXk1enXCDFBVJ")

	assert.Equal(t, "9hdxkYakTHWXR992umPcvh8bAEGG9Sdoi7uW8TKXk1enXCDFBVJ", addr.Base58(MainnetPrefix))
}

func TestAddress_Base58_Testnet(t *testing.T) {
	addr, _ := NewAddress("3WwqxmeXRWpfaH9YMLQFye7Y6ddsg1anS9hFN2EQs1P6uNMjt9tK")

	assert.Equal(t, "3WwqxmeXRWpfaH9YMLQFye7Y6ddsg1anS9hFN2EQs1P6uNMjt9tK", addr.Base58(TestnetPrefix))
}

func TestAddress_Base58_Error(t *testing.T) {
	// invalid comma
	_, err := NewAddress("9hdxkYakTHWXR992umPcvh8bAEGG9Sdoi7uW8TKXk1enXCDFBVJ,")

	assert.Error(t, err)
}

func TestAddress_TypePrefix_P2Pk(t *testing.T) {
	addr, _ := NewAddress("3WwqxmeXRWpfaH9YMLQFye7Y6ddsg1anS9hFN2EQs1P6uNMjt9tK")

	assert.Equal(t, P2PkPrefix, addr.TypePrefix())
}
