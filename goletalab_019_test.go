package btree

import ("fmt";"testing")

var _ = fmt.Sprint
func taskTree013(v ...int)*BTree{t:=New(2);for _,n:=range v{t.ReplaceOrInsert(Int(n))};return t}
func collect013(run func(ItemIterator))[]int{var out []int;run(func(v Item)bool{out=append(out,int(v.(Int)));return true});return out}
func TestGoletaBTree019(t *testing.T){tr:=taskTree013(2,4,6);if tr.Len()!=3{t.Fatalf("len=%d",tr.Len())}}

func TestGoletaBTree019AdjacentBoundary(t *testing.T) {
	// Exercise a distinct adjacent boundary of the same public contract.
	tr:=New(2);if tr.Len()!=0{t.Fatalf("empty len=%d",tr.Len())}
}
