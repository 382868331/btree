package btree
import("fmt";"testing")
var _=fmt.Sprint
func TestGoletaBTree006(t *testing.T){tr:=NewOrderedG[int](2);for _,v:=range []int{4,1,3,2}{tr.ReplaceOrInsert(v)};got:=[]int{};tr.Ascend(func(v int)bool{got=append(got,v);return len(got)<2});if fmt.Sprint(got)!="[1 2]"{t.Fatalf("got=%v",got)}}

func TestGoletaBTree006Boundary(t *testing.T){
 tr:=NewOrderedG[int](3)
 for _,v:=range []int{9,7,8}{tr.ReplaceOrInsert(v)}
 n:=0;tr.Ascend(func(int)bool{n++;return false});if n!=1{t.Fatalf("callbacks=%d",n)}
}
