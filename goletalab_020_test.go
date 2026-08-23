package uuid
import "testing"
func TestGoletaUUID020(t *testing.T){u:=NewMD5(NameSpaceDNS,[]byte("example"));if u.Version()!=3{t.Fatalf("version=%d",u.Version())}}
