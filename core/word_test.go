package core

import "testing"

type Test struct {
	In string
	Answer string
}
func TestGetNameBySnakeCase(t *testing.T){
	testSets := []Test{
		{"snake_case","Snake Case"},
		{"Hello_World","Hello World"},
		{"golang_in_progress","Golang In Progress"},
	}
	for _,set := range testSets{
		name := GetNameBySnakeCase(set.In)
		if name != set.Answer{
			t.Fatalf("actual:%s,answer:%s",name,set.Answer)
		}
	}
}

func TestRemoveExtension(t *testing.T) {
	testSets := []Test{
		{"a.mp4","a"},
		{"b.mov","b"},
		{"a.b.mp4","a.b"},
	}
	for _,set := range testSets{
		name := RemoveExtension(set.In)
		if name != set.Answer{
			t.Fatalf("actual:%s,answer:%s",name,set.Answer)
		}
	}
}
