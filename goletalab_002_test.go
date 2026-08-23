package btree

import ("fmt";"testing")

var _ = fmt.Sprint
func taskTree013(v ...int)*BTree{t:=New(2);for _,n:=range v{t.ReplaceOrInsert(Int(n))};return t}
func collect013(run func(ItemIterator))[]int{var out []int;run(func(v Item)bool{out=append(out,int(v.(Int)));return true});return out}
func TestGoletaBTree002(t *testing.T){tr:=taskTree013(1,2,3);clone:=tr.Clone();clone.ReplaceOrInsert(Int(4));if tr.Has(Int(4)){t.Fatal("clone write leaked to original")}}
