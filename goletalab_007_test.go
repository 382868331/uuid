package uuid
import "testing"
func TestGoletaUUID007(t *testing.T){u:=UUID{};u[8]=0x80;if got:=u.Variant();got!=RFC4122{t.Fatalf("variant=%v",got)}}
