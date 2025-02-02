package token

import (
	tokenpb "github.com/lutracorp/aonyx/api/protocol/pkg/token"
	"golang.org/x/crypto/poly1305"
	"google.golang.org/protobuf/proto"
)

// computeSignature computes token signature.
func computeSignature(data *tokenpb.Data, secret []byte) ([]byte, error) {
	td := proto.Clone(data).(*tokenpb.Data)
	td.Signature = nil

	mrs, err := proto.Marshal(td)
	if err != nil {
		return nil, err
	}

	var key [32]byte
	var sign [16]byte

	copy(key[:], secret)

	poly1305.Sum(&sign, mrs, &key)

	return sign[:], nil
}
