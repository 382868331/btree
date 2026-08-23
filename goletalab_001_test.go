package btree

import ("fmt";"testing")

var _ = fmt.Sprint
func taskTree013(v ...int)*BTree{t:=New(2);for _,n:=range v{t.ReplaceOrInsert(Int(n))};return t}
func collect013(run func(ItemIterator))[]int{var out []int;run(func(v Item)bool{out=append(out,int(v.(Int)));return true});return out}
func TestGoletaBTree001(t *testing.T){tr:=New(2);if (*BTreeG[Item])(tr).degree!=2{t.Fatalf("degree=%d",(*BTreeG[Item])(tr).degree)}}
