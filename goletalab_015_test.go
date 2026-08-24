package btree

import ("fmt";"testing")

var _ = fmt.Sprint
func TestGoletaBTree015(t *testing.T){tr:=NewOrderedG[int](2);for _,v:=range []int{0,-8,8,-4,4,-6,-2,2,6,-7,-5,-3,-1,1,3,5,7}{tr.ReplaceOrInsert(v)};for v:=-8;v<=8;v++{if !tr.Has(v){t.Fatalf("missing %d",v)}}}

func TestGoletaBTree015Boundary(t *testing.T) {
 tr:=NewOrderedG[int](3)
 for i:=50;i>=1;i--{tr.ReplaceOrInsert(i)}
 if _,ok:=tr.Get(37);!ok{t.Fatal("37 not reachable")}
}
