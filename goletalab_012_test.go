package uuid
import "testing"
func TestGoletaUUID012(t *testing.T){u:=MustParse("00112233-4455-6677-8899-aabbccddeeff");b,e:=u.MarshalText();if e!=nil||string(b)!=u.String(){t.Fatalf("b=%q err=%v",b,e)}}

func TestGoletaUUID012AdjacentBoundary(t *testing.T) {
	// Exercise a distinct adjacent boundary of the same public contract.
	b,e:=Max.MarshalText();if e!=nil||len(b)!=36{t.Fatalf("len=%d err=%v",len(b),e)}
}
