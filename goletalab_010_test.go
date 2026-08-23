package btree

import ("fmt";"testing")

var _ = fmt.Sprint
func taskTree013(v ...int)*BTree{t:=New(2);for _,n:=range v{t.ReplaceOrInsert(Int(n))};return t}
func collect013(run func(ItemIterator))[]int{var out []int;run(func(v Item)bool{out=append(out,int(v.(Int)));return true});return out}
func TestGoletaBTree010(t *testing.T){tr:=New(2);if old:=tr.ReplaceOrInsert(Int(5));old!=nil||!tr.Has(Int(5)){t.Fatalf("old=%v has=%v",old,tr.Has(Int(5)))}}

func TestGoletaBTree010AdjacentBoundary(t *testing.T) {
	// Exercise a distinct adjacent boundary of the same public contract.
	tr:=taskTree013(2);old:=tr.ReplaceOrInsert(Int(2));if old!=Int(2)||tr.Len()!=1{t.Fatalf("old=%v len=%d",old,tr.Len())}
}
