package uuid
import "testing"
func TestGoletaUUID011(t *testing.T){in:=UUIDs{Nil,Max};got:=in.Strings();if got[0]!=Nil.String()||got[1]!=Max.String(){t.Fatalf("got=%v",got)}}
