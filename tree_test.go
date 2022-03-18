package ergo

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewTree(t *testing.T) {
	tree, _ := NewTree("0008cd0336100ef59ced80ba5f89c4178ebd57b6c1dd0f3d135ee1db9f62fc634d637041")

	assert.NotNil(t, tree)
}

func TestNewTree_Invalid(t *testing.T) {
	_, err := NewTree("1111108zzxczbkkk")

	assert.Error(t, err)
}

func TestTree_Base16(t *testing.T) {
	tree, _ := NewTree("0008cd0336100ef59ced80ba5f89c4178ebd57b6c1dd0f3d135ee1db9f62fc634d637041")
	s, _ := tree.Base16()

	assert.Equal(t, "0008cd0336100ef59ced80ba5f89c4178ebd57b6c1dd0f3d135ee1db9f62fc634d637041", *s)
}

func TestTree_Address(t *testing.T) {
	tree, _ := NewTree("0008cd0336100ef59ced80ba5f89c4178ebd57b6c1dd0f3d135ee1db9f62fc634d637041")
	a, _ := tree.Address()

	assert.Equal(t, "9gscej8Kyzvy7AE3DMhBVpkU2CEAAM9fC6zs5dVwjAmcszPCjEr", a.Base58(MainnetPrefix))
}
