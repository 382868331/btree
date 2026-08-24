package btree
import("fmt";"testing")
var _=fmt.Sprint
func TestGoletaBTree008(t *testing.T){fl:=NewFreeListG[int](1);if !fl.freeNode(new(node[int])){t.Fatal("first node rejected")};if fl.freeNode(new(node[int])){t.Fatalf("accepted beyond capacity len=%d",len(fl.freelist))}}
