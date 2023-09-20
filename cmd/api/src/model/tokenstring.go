package model

import (
	"crypto/rand"
	"database/sql/driver"
	"encoding/base64"
	"errors"
	"fmt"
	"hash/crc32"
	"io"
	"math/big"
	"strings"
)

func generateTokenValue(prng io.Reader) (string, error) {
	var val strings.Builder
	big_range := big.NewInt(62) // a-zA-Z0-9 gives a range of 62 chars
	val.Grow(64)
	for i := 0; i < 64; i++ {
		if n, err := rand.Int(prng, big_range); err != nil {
			return "", fmt.Errorf("error getting random value: %w", err)
		} else {
			if _, err := val.WriteString(n.Text(62)); err != nil {
				return "", fmt.Errorf("error appending character to string: %w", err)
			}
		}
	}
	return val.String(), nil
}

type TokenString struct {
	Prefix    string
	value     string
	cksum     uint32
	is_legacy bool
}

func GenerateTokenString(prefix string) (TokenString, error) {
	if len(prefix) <= 0 {
		return TokenString{}, fmt.Errorf("prefix must not be empty")
	}
	if val, err := generateTokenValue(rand.Reader); err != nil {
		return TokenString{}, fmt.Errorf("error generating token string value: %w", err)
	} else {
		return CreateTokenStringWithValue(prefix, val)
	}
}

func CreateTokenStringWithValue(prefix, value string) (TokenString, error) {
	if len(prefix) <= 0 {
		return TokenString{}, fmt.Errorf("prefix must not be empty")
	}
	if len(value) != 64 {
		return TokenString{}, fmt.Errorf("value must be of length 64")
	}
	new := TokenString{Prefix: strings.ToUpper(prefix), value: value}
	new.cksum = crc32.ChecksumIEEE([]byte(new.Prefix + new.value))
	return new, nil
}

func formatChecksum(cksum uint32) string {
	return strings.ReplaceAll(fmt.Sprintf("%6s", big.NewInt(int64(cksum)).Text(62)), " ", "0")
}

// This method isn't really necessary anymore, but leaving it in case
// we want to modify what part of a token is hashed in the future
func (s TokenString) DigestableValue() ([]byte, error) {
	if s.value == "" {
		return []byte{}, errors.New("token value is not set")
	} else {
		return []byte(s.String()), nil
	}
}

func (s TokenString) String() string {
	if s.value == "" {
		return ""
	} else if s.is_legacy {
		return s.value
	}
	return fmt.Sprintf("%s_%s%s", s.Prefix, s.value, formatChecksum(s.cksum))
}

func (s TokenString) MarshalText() ([]byte, error) {
	return []byte(s.String()), nil
}

func (s *TokenString) loadFromString(src string) error {
	if src == "" {
		*s = TokenString{}
		return nil
	}
	ts, err := ParseTokenString(src)
	if err != nil {
		return fmt.Errorf("unable to parse value: %w", err)
	}
	*s = ts
	return nil
}

func (s *TokenString) UnmarshalText(text []byte) error {
	return s.loadFromString(string(text))
}

func (s TokenString) Value() (driver.Value, error) {
	return s.String(), nil
}

func (s *TokenString) Scan(src any) error {
	src_str, ok := src.(string)
	if !ok {
		return errors.New("expected value of type string")
	}
	return s.loadFromString(src_str)
}

func isValidBase62(val string) bool {
	for _, v := range []byte(val) {
		if (v >= '0' && v <= '9') || (v >= 'A' && v <= 'Z') || (v >= 'a' && v <= 'z') {
			continue
		}
		return false
	}
	return true
}

func isValidBase64Chars(val string) bool {
	vlen := len(val)
	if vlen == 0 {
		return true
	}

	val = strings.TrimRight(val, "=")
	if val == "" || vlen-len(val) > 2 {
		return false
	}

	for _, v := range []byte(val) {
		if (v >= '0' && v <= '9') || (v >= 'A' && v <= 'Z') || (v >= 'a' && v <= 'z') {
			continue
		}
		if v == '+' || v == '/' {
			continue
		}
		return false
	}
	return true
}

func isValidBase64(val string) bool {
	// DecodeString() will accept some character sequences that contain
	// non base64 characters, as well as whitespace. This filters those out.
	if !isValidBase64Chars(val) {
		return false
	}
	_, err := base64.StdEncoding.DecodeString(val)
	return err == nil
}

var ErrTokenStringFormat = errors.New("invalid token format")
var ErrTokenStringChecksum = errors.New("token checksum is invalid")

func ParseTokenString(src string) (TokenString, error) {
	// if legacy style token
	if len(src) == 64 && !strings.Contains(src, "_") && isValidBase64(src) {
		return TokenString{value: src, is_legacy: true}, nil
	}

	// min length is prefix (1) + value (64) + cksum (6)
	if len(src) < 1+64+6 {
		return TokenString{}, fmt.Errorf("%w: token string is too short", ErrTokenStringFormat)
	}
	var src_prefix, src_val, src_cksum string
	src_cksum = src[len(src)-6:]
	src = src[:len(src)-6]
	if i := strings.LastIndex(src, "_"); i < 0 {
		return TokenString{}, fmt.Errorf("%w: token string missing prefix separator", ErrTokenStringFormat)
	} else {
		src_prefix = src[:i]
		src_val = src[i+1:]
	}

	if !isValidBase62(src_cksum) {
		return TokenString{}, fmt.Errorf("%w: token string checksum contains invalid characters", ErrTokenStringFormat)
	}
	tok, err := CreateTokenStringWithValue(src_prefix, src_val)
	if err != nil {
		return TokenString{}, err
	}
	if formatChecksum(tok.cksum) != src_cksum {
		return TokenString{}, ErrTokenStringChecksum
	}
	return tok, nil
}
