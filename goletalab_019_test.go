package uuid
import "testing"
func TestGoletaUUID019(t *testing.T){nu:=NullUUID{};b,e:=nu.MarshalJSON();if e!=nil||string(b)!="null"{t.Fatalf("b=%q err=%v",b,e)}}
