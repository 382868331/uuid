package uuid
import "testing"
func TestGoletaUUID008(t *testing.T){u:=UUID{};u[6]=0x4a;if got:=u.Version();got!=4{t.Fatalf("version=%d",got)}}
