package uuid
import "testing"
func TestGoletaUUID007(t *testing.T){u:=UUID{};u[8]=0x80;if got:=u.Variant();got!=RFC4122{t.Fatalf("variant=%v",got)}}

func TestGoletaUUID007AdjacentBoundary(t *testing.T) {
	// Exercise a distinct adjacent boundary of the same public contract.
	u:=UUID{};u[8]=0xbf;if got:=u.Variant();got!=RFC4122{t.Fatalf("variant=%v",got)}
}
