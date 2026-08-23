package uuid
import "testing"
func TestGoletaUUID002(t *testing.T){raw:=[]byte{0,1,2,3,4,5,6,7,8,9,10,11,12,13,14,15};u,e:=FromBytes(raw);if e!=nil||u[15]!=15{t.Fatalf("u=%v err=%v",u,e)}}
