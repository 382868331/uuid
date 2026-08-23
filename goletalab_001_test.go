package uuid
import "testing"
func TestGoletaUUID001(t *testing.T){a:=UUID{1};b:=UUID{2};if Compare(a,b)>=0||Compare(b,a)<=0{t.Fatalf("ab=%d ba=%d",Compare(a,b),Compare(b,a))}}

func TestGoletaUUID001AdjacentBoundary(t *testing.T) {
	// Exercise a distinct adjacent boundary of the same public contract.
	a:=UUID{};if Compare(a,a)!=0{t.Fatalf("equal=%d",Compare(a,a))}
}
