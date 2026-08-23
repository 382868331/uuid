package uuid
import "testing"
func TestGoletaUUID014(t *testing.T){u:=MustParse("00112233-4455-6677-8899-aabbccddeeff");b,e:=u.MarshalBinary();if e!=nil||len(b)!=16||b[15]!=0xff{t.Fatalf("len=%d last=%x err=%v",len(b),b[len(b)-1],e)}}
