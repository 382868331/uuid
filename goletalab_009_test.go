package uuid
import "testing"
func TestGoletaUUID009(t *testing.T){if got:=Version(15).String();got!="VERSION_15"{t.Fatalf("got=%q",got)}}

func TestGoletaUUID009AdjacentBoundary(t *testing.T) {
	// Exercise a distinct adjacent boundary of the same public contract.
	if got:=Version(16).String();got!="BAD_VERSION_16"{t.Fatalf("got=%q",got)}
}
