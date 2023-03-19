package forge

import (
	"testing"
)

func TestEncryptDecrypt(t *testing.T) {
    c := Encrypt("hello", "world", "")
    res, err := Decrypt("hello", c)
    if err != nil {
        t.Errorf("Decryption failed: %v", err) 
    }
    if res != "world" {
        t.Errorf("Expected world, received: " + res)
    }
    
}
