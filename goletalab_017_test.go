package btree

import ("fmt";"testing")

var _ = fmt.Sprint
func TestGoletaBTree017(t *testing.T){tr:=NewOrderedG[int](3);for _,v:=range []int{1,2,3}{tr.ReplaceOrInsert(v)};if _,ok:=tr.Delete(2);!ok{t.Fatal("delete failed")};if tr.Len()!=2{t.Fatalf("len=%d",tr.Len())};tr.Delete(99);if tr.Len()!=2{t.Fatalf("missing delete len=%d",tr.Len())}}

func TestGoletaBTree017Boundary(t *testing.T) {
 tr:=NewOrderedG[int](2)
 tr.ReplaceOrInsert(8); tr.Delete(8)
 if tr.Len()!=0{t.Fatalf("len=%d",tr.Len())}
}
