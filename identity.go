package p2pid

import (
	"bytes"
	"crypto/sha256"
	"io"
	"os"

	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p/core/crypto"
)

// Peer ID Seed Env String
const PID_SEED = "PID_SEED"

func PersistentIdentity() (libp2p.Option, error) {
	return PersistentIdentityFromEnv(PID_SEED)
}

func PersistentIdentityFromEnv(e string) (libp2p.Option, error) {
	seed := os.Getenv(e)
	// If the seed is empty, we return a nil option representing random identity
	if seed == "" {
		return nil, nil
	}

	// If the seed is not empty, we generate the identity from the seed
	reader := hashedReader(seed)
	return PersistentIdentityFromReader(reader)
}

func PersistentIdentityFromReader(r io.Reader) (libp2p.Option, error) {
	privKey, _, err := crypto.GenerateEd25519Key(r)
	if err != nil {
		return nil, err
	}

	return libp2p.Identity(privKey), nil
}

func hashedReader(s string) io.Reader {
	// Hash the seed to get a 32 byte seed using SHA256
	hash := sha256.New()
	hash.Write([]byte(s))
	return bytes.NewReader(hash.Sum(nil))
}
