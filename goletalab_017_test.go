package uuid
import "testing"
func TestGoletaUUID017(t *testing.T){nu:=NullUUID{};v,e:=nu.Value();if e!=nil||v!=nil{t.Fatalf("value=%#v err=%v",v,e)}}

func TestGoletaUUID017AdjacentBoundary(t *testing.T) {
	// Exercise a distinct adjacent boundary of the same public contract.
	nu:=NullUUID{UUID:MustParse("00112233-4455-6677-8899-aabbccddeeff"),Valid:true};v,e:=nu.Value();if e!=nil||v==nil{t.Fatalf("value=%#v err=%v",v,e)}
}
