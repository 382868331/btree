package btree
import("fmt";"testing")
var _=fmt.Sprint
func TestGoletaBTree008(t *testing.T){fl:=NewFreeListG[int](1);if !fl.freeNode(new(node[int])){t.Fatal("first node rejected")};if fl.freeNode(new(node[int])){t.Fatalf("accepted beyond capacity len=%d",len(fl.freelist))}}

func TestGoletaBTree008Boundary(t *testing.T){
 fl:=NewFreeListG[int](0)
 if fl.freeNode(new(node[int])){t.Fatal("zero-capacity freelist retained node")}
 if len(fl.freelist)!=0{t.Fatalf("len=%d",len(fl.freelist))}
}
