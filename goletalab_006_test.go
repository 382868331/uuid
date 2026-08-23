package uuid
import "testing"
func TestGoletaUUID006(t *testing.T){u:=MustParse("00112233-4455-6677-8899-aabbccddeeff");if got:=u.String();got[8]!='-'||got[7]=='-'{t.Fatalf("got=%q",got)}}
