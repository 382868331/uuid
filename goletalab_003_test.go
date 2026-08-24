package uuid
import "testing"
func TestGoletaUUID003(t *testing.T){if err:=Validate("00112233-4455-6677-8899-aabbccddeeff");err!=nil{t.Fatalf("valid rejected: %v",err)}}

func TestGoletaUUID003AdjacentBoundary(t *testing.T) {
	// Exercise a distinct adjacent boundary of the same public contract.
	if err:=Validate("ffffffff-ffff-ffff-ffff-ffffffffffff");err!=nil{t.Fatalf("valid rejected: %v",err)}
}
