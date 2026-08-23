package uuid
import "testing"
func TestGoletaUUID017(t *testing.T){nu:=NullUUID{};v,e:=nu.Value();if e!=nil||v!=nil{t.Fatalf("value=%#v err=%v",v,e)}}
