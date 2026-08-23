package uuid
import "testing"
func TestGoletaUUID005(t *testing.T){u:=MustParse("00112233-4455-6677-8899-aabbccddeeff");if got:=u.URN();got!="urn:uuid:00112233-4455-6677-8899-aabbccddeeff"{t.Fatalf("got=%q",got)}}
