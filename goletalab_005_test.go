package btree
import("fmt";"testing")
var _=fmt.Sprint
func TestGoletaBTree005(t *testing.T){tr:=NewOrderedG[int](3);tr.ReplaceOrInsert(8);if !tr.Has(8)||tr.Len()!=1{t.Fatalf("has=%v len=%d",tr.Has(8),tr.Len())}}
