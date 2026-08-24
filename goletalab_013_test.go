package btree

import ("fmt";"testing")

var _ = fmt.Sprint
func TestGoletaBTree013(t *testing.T){tr:=NewOrderedG[int](4); for _,v:=range []int{2,4,6}{tr.ReplaceOrInsert(v)}; old,ok:=tr.ReplaceOrInsert(2); if !ok||old!=2||tr.Len()!=3{t.Fatalf("old=%d ok=%v len=%d",old,ok,tr.Len())}}

func TestGoletaBTree013Boundary(t *testing.T) {
 tr:=NewOrderedG[int](3)
 tr.ReplaceOrInsert(-1); tr.ReplaceOrInsert(5)
 _,ok:=tr.ReplaceOrInsert(-1); if !ok||tr.Len()!=2{t.Fatalf("ok=%v len=%d",ok,tr.Len())}
}
