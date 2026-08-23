package uuid
import "testing"
func TestGoletaUUID009(t *testing.T){if got:=Version(15).String();got!="VERSION_15"{t.Fatalf("got=%q",got)}}
