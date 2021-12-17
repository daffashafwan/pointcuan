package randomizer

import (
    "crypto/rand"
)

var chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890-"

func Randomize() string{
	ll := len(chars)
    b := make([]byte, 20)
    rand.Read(b) // generates len(b) random bytes
    for i := 0; i < 20; i++ {
        b[i] = chars[int(b[i])%ll]
    }
    return string(b)
}