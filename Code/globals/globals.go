package globals

import (
	"crypto/rand"
	"errors"

	"github.com/gorilla/schema"
	_ "github.com/gorilla/schema"
	"github.com/sirupsen/logrus"
)

var (
	Decoder      *schema.Decoder
	Log          *logrus.Logger
	TokenSignKey []byte
)

func Init() {
	var err error

	Decoder = schema.NewDecoder()
	Log = logrus.New()

	if TokenSignKey, err = GenSymmetricKey(64); err != nil {
		panic(err)
	}
}

// GenSymmetricKey generates a key for the JWT encryption
// https://github.com/northbright/Notes/blob/master/jwt/generate_hmac_secret_key_for_jwt.md
func GenSymmetricKey(bits int) (k []byte, err error) {
	if bits <= 0 || bits%8 != 0 {
		return nil, errors.New("Key size error")
	}

	size := bits / 8
	k = make([]byte, size)
	if _, err = rand.Read(k); err != nil {
		return nil, err
	}

	return k, nil
}
