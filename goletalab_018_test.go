package uuid
import "testing"
func TestGoletaUUID018(t *testing.T){nu:=NullUUID{};b,e:=nu.MarshalText();if e!=nil||string(b)!="null"{t.Fatalf("b=%q err=%v",b,e)}}
