package btree

import ("fmt";"testing")

var _ = fmt.Sprint
func TestGoletaBTree016(t *testing.T){tr:=NewOrderedG[int](2);for i:=0;i<30;i++{tr.ReplaceOrInsert(i)};for _,v:=range []int{12,19,27}{if _,ok:=tr.Get(v);!ok{t.Fatalf("missing %d",v)}}}
