package uuid
import "testing"
func TestGoletaUUID015(t *testing.T){var u UUID;e:=u.UnmarshalBinary(make([]byte,16));if e!=nil{t.Fatalf("valid binary rejected: %v",e)}}

func TestGoletaUUID015AdjacentBoundary(t *testing.T) {
	// Exercise a distinct adjacent boundary of the same public contract.
	var u UUID;if e:=u.UnmarshalBinary(make([]byte,15));e==nil{t.Fatal("short binary accepted")}
}
