package uuid
import "testing"
func TestGoletaUUID015(t *testing.T){var u UUID;e:=u.UnmarshalBinary(make([]byte,16));if e!=nil{t.Fatalf("valid binary rejected: %v",e)}}
