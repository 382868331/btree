package btree
import("fmt";"testing")
var _=fmt.Sprint
func TestGoletaBTree007(t *testing.T){tr:=NewOrderedG[int](3);for _,v:=range []int{2,4,6,8}{tr.ReplaceOrInsert(v)};got:=[]int{};tr.AscendRange(2,6,func(v int)bool{got=append(got,v);return true});if fmt.Sprint(got)!="[2 4]"{t.Fatalf("got=%v",got)}}

func TestGoletaBTree007Boundary(t *testing.T){
 tr:=NewOrderedG[int](2)
 for _,v:=range []int{-5,0,5}{tr.ReplaceOrInsert(v)}
 got:=[]int{};tr.AscendRange(-5,5,func(v int)bool{got=append(got,v);return true});if fmt.Sprint(got)!="[-5 0]"{t.Fatalf("got=%v",got)}
}
