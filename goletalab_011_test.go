package uuid
import "testing"
func TestGoletaUUID011(t *testing.T){in:=UUIDs{Nil,Max};got:=in.Strings();if got[0]!=Nil.String()||got[1]!=Max.String(){t.Fatalf("got=%v",got)}}

func TestGoletaUUID011AdjacentBoundary(t *testing.T) {
	// Exercise a distinct adjacent boundary of the same public contract.
	in:=UUIDs{MustParse("00112233-4455-6677-8899-aabbccddeeff")};got:=in.Strings();if len(got)!=1||got[0]!=in[0].String(){t.Fatalf("got=%v",got)}
}
