package btree

import ("fmt";"testing")

var _ = fmt.Sprint
func TestGoletaBTree012(t *testing.T){tr:=NewOrderedG[int](2); for i:=0;i<12;i++{tr.ReplaceOrInsert(i)}; if _,ok:=tr.Delete(5);!ok{t.Fatal("delete failed")}; for _,v:=range []int{4,6,7}{if !tr.Has(v){t.Fatalf("neighbor %d lost",v)}}}

func TestGoletaBTree012Boundary(t *testing.T) {
 tr:=NewOrderedG[int](3)
 for _,v:=range []int{20,10,30,25,35}{tr.ReplaceOrInsert(v)}
 tr.Delete(25); if tr.Len()!=4||!tr.Has(30){t.Fatalf("len=%d",tr.Len())}
}
