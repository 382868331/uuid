package uuid
import "testing"
func TestGoletaUUID018(t *testing.T){nu:=NullUUID{};b,e:=nu.MarshalText();if e!=nil||string(b)!="null"{t.Fatalf("b=%q err=%v",b,e)}}

func TestGoletaUUID018AdjacentBoundary(t *testing.T) {
	// Exercise a distinct adjacent boundary of the same public contract.
	nu:=NullUUID{UUID:MustParse("00112233-4455-6677-8899-aabbccddeeff"),Valid:true};b,e:=nu.MarshalText();if e!=nil||string(b)!=nu.UUID.String(){t.Fatalf("b=%q err=%v",b,e)}
}
