package crypto

import (
	"bytes"
	"crypto/ed25519"
	"crypto/rand"
	"fmt"
)

type ed25519key struct {
	priv ed25519.PrivateKey
}

func checkSequenceIsNil(seq *uint32) {
	if seq != nil {
		panic("Ed25519 keys do not support account families")
	}
}

func (e *ed25519key) Id(seq *uint32) []byte {
	checkSequenceIsNil(seq)
	return Sha256RipeMD160(e.Public(seq))
}

func (e *ed25519key) Public(seq *uint32) []byte {
	checkSequenceIsNil(seq)
	return append([]byte{0xED}, e.priv[32:]...)
}

func (e *ed25519key) Private(seq *uint32) []byte {
	checkSequenceIsNil(seq)
	return e.priv[:]
}

func (e *ed25519key) PrivateShort(seq *uint32) []byte {
	checkSequenceIsNil(seq)
	return e.priv[:32]
}

func NewEd25519Key(seed []byte) (*ed25519key, error) {
	r := rand.Reader
	if seed != nil {
		r = bytes.NewReader(Sha512Half(seed))
	}
	_, priv, err := ed25519.GenerateKey(r)
	if err != nil {
		return nil, err
	}
	return &ed25519key{priv: priv}, nil
}

func NewEd25519KeyFromRaw(raw []byte) (*ed25519key, error) {
	if raw == nil || len(raw) != 32 {
		return nil, fmt.Errorf("raw is wrong size, expected: 32 got: %d", len(raw))
	}
	r := rand.Reader
	r = bytes.NewReader(raw)
	_, priv, err := ed25519.GenerateKey(r)
	if err != nil {
		return nil, err
	}
	return &ed25519key{priv: priv}, nil
}
