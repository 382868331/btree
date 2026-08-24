package btree

import ("fmt";"testing")

var _ = fmt.Sprint
func TestGoletaBTree012(t *testing.T){tr:=NewOrderedG[int](2); for i:=0;i<12;i++{tr.ReplaceOrInsert(i)}; if _,ok:=tr.Delete(5);!ok{t.Fatal("delete failed")}; for _,v:=range []int{4,6,7}{if !tr.Has(v){t.Fatalf("neighbor %d lost",v)}}}
