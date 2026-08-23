package uuid
import "testing"
func TestGoletaUUID010(t *testing.T){if got:=Microsoft.String();got!="Microsoft"{t.Fatalf("got=%q",got)}}

func TestGoletaUUID010AdjacentBoundary(t *testing.T) {
	// Exercise a distinct adjacent boundary of the same public contract.
	if got:=RFC4122.String();got!="RFC4122"{t.Fatalf("got=%q",got)}
}
