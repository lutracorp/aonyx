package token

import (
	"encoding/base64"

	tokenpb "github.com/lutracorp/aonyx/api/protocol/pkg/token"
	"google.golang.org/protobuf/proto"
)

// Marshal encodes token data to strings.
func Marshal(data *tokenpb.Data) (string, error) {
	pm, err := proto.Marshal(data)
	if err != nil {
		return "", err
	}

	return base64.RawURLEncoding.EncodeToString(pm), nil
}

// Unmarshal decodes given string to token data.
func Unmarshal(src string, dst *tokenpb.Data) error {
	str, err := base64.RawURLEncoding.DecodeString(src)
	if err != nil {
		return err
	}

	return proto.Unmarshal(str, dst)
}
