package uuid
import "testing"
func TestGoletaUUID013(t *testing.T){var got UUID;if e:=got.UnmarshalText([]byte("00112233-4455-6677-8899-aabbccddeeff"));e!=nil||got==Nil{t.Fatalf("got=%v err=%v",got,e)}}
