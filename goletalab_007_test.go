package btree
import("fmt";"testing")
var _=fmt.Sprint
func TestGoletaBTree007(t *testing.T){tr:=NewOrderedG[int](3);for _,v:=range []int{2,4,6,8}{tr.ReplaceOrInsert(v)};got:=[]int{};tr.AscendRange(2,6,func(v int)bool{got=append(got,v);return true});if fmt.Sprint(got)!="[2 4]"{t.Fatalf("got=%v",got)}}
