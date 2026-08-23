package uuid
import "testing"
func TestGoletaUUID012(t *testing.T){u:=MustParse("00112233-4455-6677-8899-aabbccddeeff");b,e:=u.MarshalText();if e!=nil||string(b)!=u.String(){t.Fatalf("b=%q err=%v",b,e)}}
