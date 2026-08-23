package uuid
import "testing"
func TestGoletaUUID016(t *testing.T){nu:=NullUUID{UUID:Max,Valid:true};if e:=nu.Scan(nil);e!=nil||nu.Valid||nu.UUID!=Nil{t.Fatalf("nu=%+v err=%v",nu,e)}}

func TestGoletaUUID016AdjacentBoundary(t *testing.T) {
	// Exercise a distinct adjacent boundary of the same public contract.
	nu:=NullUUID{};if e:=nu.Scan(nil);e!=nil||nu.Valid{t.Fatalf("nu=%+v",nu)}
}
