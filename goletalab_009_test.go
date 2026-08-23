package btree

import ("fmt";"testing")

var _ = fmt.Sprint
func taskTree013(v ...int)*BTree{t:=New(2);for _,n:=range v{t.ReplaceOrInsert(Int(n))};return t}
func collect013(run func(ItemIterator))[]int{var out []int;run(func(v Item)bool{out=append(out,int(v.(Int)));return true});return out}
func TestGoletaBTree009(t *testing.T){tr:=taskTree013(1,3);if !tr.Has(Int(3))||tr.Has(Int(2)){t.Fatalf("present=%v absent=%v",tr.Has(Int(3)),tr.Has(Int(2)))}}

func TestGoletaBTree009AdjacentBoundary(t *testing.T) {
	// Exercise a distinct adjacent boundary of the same public contract.
	tr:=taskTree013(8);if !tr.Has(Int(8))||tr.Has(Int(9)){t.Fatal("membership inverted")}
}
