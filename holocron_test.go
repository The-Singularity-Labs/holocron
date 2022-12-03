package holocron

import (
	"testing"
)

func TestTwoVariableSystem(t *testing.T) {
    c := Encrypt("hello", "world", "")
    if res := Decrypt("hello", c); res != "world" {
        t.Errorf("Expected world, received: " + res)
    }
    
}
