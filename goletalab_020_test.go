package uuid
import "testing"
func TestGoletaUUID020(t *testing.T){u:=NewMD5(NameSpaceDNS,[]byte("example"));if u.Version()!=3{t.Fatalf("version=%d",u.Version())}}

func TestGoletaUUID020AdjacentBoundary(t *testing.T) {
	// Exercise a distinct adjacent boundary of the same public contract.
	u:=NewSHA1(NameSpaceDNS,[]byte("example"));if u.Version()!=5{t.Fatalf("version=%d",u.Version())}
}
