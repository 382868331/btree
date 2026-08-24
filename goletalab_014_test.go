package btree

import ("fmt";"testing")

var _ = fmt.Sprint
func TestGoletaBTree014(t *testing.T){tr:=NewOrderedG[int](2); for _,v:=range []int{5,1,9,3,7,2,8,4,6}{tr.ReplaceOrInsert(v)}; got:=[]int{};tr.Ascend(func(v int)bool{got=append(got,v);return true}); if fmt.Sprint(got)!="[1 2 3 4 5 6 7 8 9]"{t.Fatalf("got=%v",got)}}

func TestGoletaBTree014Boundary(t *testing.T) {
 tr:=NewOrderedG[int](2)
 for i:=20;i>=0;i--{tr.ReplaceOrInsert(i)}
 if tr.Len()!=21||!tr.Has(10){t.Fatalf("len=%d",tr.Len())}
}
