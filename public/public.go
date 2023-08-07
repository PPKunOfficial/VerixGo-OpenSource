package public

import (
	"embed"
	_ "embed"
)

//go:embed dist
var Public embed.FS

//go:embed key/public.pem
var PublicKey []byte

//go:embed key/private.pem
var PrivateKey []byte
