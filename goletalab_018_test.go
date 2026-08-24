package btree

import ("fmt";"testing")

var _ = fmt.Sprint
func TestGoletaBTree018(t *testing.T){tr:=NewOrderedG[int](2);func(){defer func(){if r:=recover();r!=nil{t.Fatalf("empty lookup panic: %v",r)}}();if _,ok:=tr.Get(7);ok{t.Fatal("unexpected hit")}}()}

func TestGoletaBTree018Boundary(t *testing.T) {
 tr:=NewOrderedG[int](4)
 if tr.Has(-3){t.Fatal("empty tree reported key")}; tr.ReplaceOrInsert(1)
 if !tr.Has(1){t.Fatal("tree unusable after empty lookup")}
}
