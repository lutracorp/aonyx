package token

import (
	"bytes"

	tokenpb "github.com/lutracorp/aonyx/api/protocol/pkg/token"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// Sign generates signed token and encodes it.
func Sign(payload []byte, secret []byte) (string, error) {
	data := &tokenpb.Data{
		Payload:   payload,
		Timestamp: timestamppb.Now(),
	}

	sign, err := SignData(data, secret)
	if err != nil {
		return "", err
	}

	return Marshal(sign)
}

// SignData generates signed token data.
func SignData(data *tokenpb.Data, secret []byte) (*tokenpb.Data, error) {
	sign, err := computeSignature(data, secret)
	if err != nil {
		return nil, err
	}

	data.Signature = sign
	return data, nil
}

// Verify decodes and verifies token for validity.
func Verify(token string, secret []byte) (bool, error) {
	data := &tokenpb.Data{}
	if err := Unmarshal(token, data); err != nil {
		return false, err
	}

	return VerifyData(data, secret)
}

// VerifyData verifies token data for validity.
func VerifyData(data *tokenpb.Data, secret []byte) (bool, error) {
	sign, err := computeSignature(data, secret)
	if err != nil {
		return false, err
	}

	return bytes.Equal(sign, data.Signature), nil
}
