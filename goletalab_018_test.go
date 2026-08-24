package btree

import ("fmt";"testing")

var _ = fmt.Sprint
func TestGoletaBTree018(t *testing.T){tr:=NewOrderedG[int](2);func(){defer func(){if r:=recover();r!=nil{t.Fatalf("empty lookup panic: %v",r)}}();if _,ok:=tr.Get(7);ok{t.Fatal("unexpected hit")}}()}
