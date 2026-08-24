package btree

import ("fmt";"testing")

var _ = fmt.Sprint
func TestGoletaBTree015(t *testing.T){tr:=NewOrderedG[int](2);for _,v:=range []int{0,-8,8,-4,4,-6,-2,2,6,-7,-5,-3,-1,1,3,5,7}{tr.ReplaceOrInsert(v)};for v:=-8;v<=8;v++{if !tr.Has(v){t.Fatalf("missing %d",v)}}}
