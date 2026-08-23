package btree

import ("fmt";"testing")

var _ = fmt.Sprint
func taskTree013(v ...int)*BTree{t:=New(2);for _,n:=range v{t.ReplaceOrInsert(Int(n))};return t}
func collect013(run func(ItemIterator))[]int{var out []int;run(func(v Item)bool{out=append(out,int(v.(Int)));return true});return out}
func TestGoletaBTree003(t *testing.T){tr:=taskTree013(1,2,3);got:=tr.Delete(Int(2));if got!=Int(2)||tr.Has(Int(2)){t.Fatalf("got=%v has=%v",got,tr.Has(Int(2)))}}

func TestGoletaBTree003AdjacentBoundary(t *testing.T) {
	// Exercise a distinct adjacent boundary of the same public contract.
	tr:=taskTree013(7,8);_ = tr.Delete(Int(7));if tr.Len()!=1||tr.Has(Int(7)){t.Fatalf("len=%d",tr.Len())}
}
