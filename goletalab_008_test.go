package uuid
import "testing"
func TestGoletaUUID008(t *testing.T){u:=UUID{};u[6]=0x4a;if got:=u.Version();got!=4{t.Fatalf("version=%d",got)}}

func TestGoletaUUID008AdjacentBoundary(t *testing.T) {
	// Exercise a distinct adjacent boundary of the same public contract.
	u:=UUID{};u[6]=0x7f;if got:=u.Version();got!=7{t.Fatalf("version=%d",got)}
}
