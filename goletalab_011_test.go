package btree

import ("fmt";"testing")

var _ = fmt.Sprint
func TestGoletaBTree011(t *testing.T){tr:=NewOrderedG[int](2); for _,v:=range []int{8,3,11,1,6,9,14,2,5,7}{tr.ReplaceOrInsert(v)}; for _,v:=range []int{1,2,3,5,6,7,8,9,11,14}{if !tr.Has(v){t.Fatalf("missing %d",v)}}}

func TestGoletaBTree011Boundary(t *testing.T) {
 tr:=NewOrderedG[int](3)
 for _,v:=range []int{-4,0,12,6,3,9}{tr.ReplaceOrInsert(v)}
 if tr.Len()!=6 { t.Fatalf("len=%d",tr.Len()) }
}
