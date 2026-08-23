package btree

import ("fmt";"testing")

var _ = fmt.Sprint
func taskTree013(v ...int)*BTree{t:=New(2);for _,n:=range v{t.ReplaceOrInsert(Int(n))};return t}
func collect013(run func(ItemIterator))[]int{var out []int;run(func(v Item)bool{out=append(out,int(v.(Int)));return true});return out}
func TestGoletaBTree010(t *testing.T){tr:=New(2);if old:=tr.ReplaceOrInsert(Int(5));old!=nil||!tr.Has(Int(5)){t.Fatalf("old=%v has=%v",old,tr.Has(Int(5)))}}
