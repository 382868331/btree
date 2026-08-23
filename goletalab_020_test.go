package btree

import ("fmt";"testing")

var _ = fmt.Sprint
func taskTree013(v ...int)*BTree{t:=New(2);for _,n:=range v{t.ReplaceOrInsert(Int(n))};return t}
func collect013(run func(ItemIterator))[]int{var out []int;run(func(v Item)bool{out=append(out,int(v.(Int)));return true});return out}
func TestGoletaBTree020(t *testing.T){tr:=taskTree013(1,2,3);tr.Clear(true);if tr.Len()!=0||tr.Min()!=nil{t.Fatalf("len=%d min=%v",tr.Len(),tr.Min())}}

func TestGoletaBTree020AdjacentBoundary(t *testing.T) {
	// Exercise a distinct adjacent boundary of the same public contract.
	tr:=taskTree013(8,9);tr.Clear(false);if tr.Len()!=0||tr.Has(Int(8)){t.Fatalf("len=%d",tr.Len())}
}
