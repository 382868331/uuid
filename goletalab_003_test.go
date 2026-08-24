package uuid
import "testing"
func TestGoletaUUID003(t *testing.T){if err:=Validate("00112233-4455-6677-8899-aabbccddeeff");err!=nil{t.Fatalf("valid rejected: %v",err)}}
