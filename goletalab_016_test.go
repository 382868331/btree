package btree

import ("fmt";"testing")

var _ = fmt.Sprint
func TestGoletaBTree016(t *testing.T){tr:=NewOrderedG[int](2);for i:=0;i<30;i++{tr.ReplaceOrInsert(i)};for _,v:=range []int{12,19,27}{if _,ok:=tr.Get(v);!ok{t.Fatalf("missing %d",v)}}}

func TestGoletaBTree016Boundary(t *testing.T) {
 tr:=NewOrderedG[int](3)
 for _,v:=range []int{100,20,180,60,140,10,30,170,190}{tr.ReplaceOrInsert(v)}
 if !tr.Has(170){t.Fatal("right subtree lookup failed")}
}
