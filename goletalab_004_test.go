package uuid
import "testing"
func TestGoletaUUID004(t *testing.T){u:=MustParse("00112233-4455-6677-8899-aabbccddeeff");got:=u.String();if got!="00112233-4455-6677-8899-aabbccddeeff"{t.Fatalf("got=%q",got)}}

func TestGoletaUUID004AdjacentBoundary(t *testing.T) {
	// Exercise a distinct adjacent boundary of the same public contract.
	u:=Max;if len(u.String())!=36{t.Fatalf("len=%d",len(u.String()))}
}
