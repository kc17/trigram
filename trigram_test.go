package trigram_test

import (
	"testing"

	. "github.com/kkdai/trigram"
)

func TestTrigramlize(t *testing.T) {
	ret := ExtractStringToTrigram("Cod")
	if ret[0] != 4419428 {
		t.Errorf("Trigram failed, expect 4419428\n")
	}

	//string length longer than 3
	ret = ExtractStringToTrigram("Code")
	if ret[0] != 4419428 && ret[1] != 7300197 {
		t.Errorf("Trigram failed on longer string")
	}
}

func TestMapIntersect(t *testing.T) {
	mapA := make(map[int]bool)
	mapB := make(map[int]bool)

	mapA[1] = true
	mapA[2] = true
	mapB[1] = true

	ret := IntersectTwoMap(mapA, mapB)
	if len(ret) != 1 || ret[1] == false {
		t.Errorf("Map intersect error")
	}

	ret = IntersectTwoMap(mapB, mapA)
	if len(ret) != 1 || ret[1] == false {
		t.Errorf("Map intersect error")
	}

	mapA[3] = true
	mapB[3] = true
	mapA[4] = true

	ret = IntersectTwoMap(mapB, mapA)
	if len(ret) != 2 || ret[1] == false {
		t.Errorf("Map intersect error")
	}
}

func TestTrigramIndexBasicQuery(t *testing.T) {
	ti := NewTrigramIndex()
	ti.Add("Code is my life")
	ti.Add("Search")
	ti.Add("I write a lot of Codes")

	ret := ti.Query("Code")
	if ret[0] != 1 || ret[1] != 3 {
		t.Errorf("Basic query is failed.")
	}
}

func TestEmptyLessQuery(t *testing.T) {
	ti := NewTrigramIndex()
	ti.Add("Code is my life")
	ti.Add("Search")
	ti.Add("I write a lot of Codes")

	ret := ti.Query("te") //less than 3, should get all doc ID
	if len(ret) != 3 || ret[0] != 1 || ret[2] != 3 {
		t.Errorf("Error on less than 3 character query")
	}

	ret = ti.Query("")
	if len(ret) != 3 || ret[0] != 1 || ret[2] != 3 {
		t.Errorf("Error on empty character query")
	}
}

func TestDelete(t *testing.T) {
	ti := NewTrigramIndex()
	ti.Add("Code is my life")

	ti.Delete("Code", 1)
	ret := ti.Query("Code")
	if len(ret) != 0 {
		t.Error("Basic delete failed", ret)
	}

	ret = ti.Query("life")
	if len(ret) != 1 || ret[0] != 1 {
		t.Error("Basic delete failed", ret)
	}
}

func BenchmarkAdd(b *testing.B) {
	big := NewTrigramIndex()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		big.Add("1234567890")
	}
}

func BenchmarkDelete(b *testing.B) {

	big := NewTrigramIndex()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		big.Add("1234567890")
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		big.Delete("1234567890", i)
	}
}

func BenchmarkQuery(b *testing.B) {

	big := NewTrigramIndex()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		big.Add("1234567890")
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		big.Query("1234567890")
	}
}

func BenchmarkIntersect(b *testing.B) {

	DocA := make(map[int]bool)
	DocB := make(map[int]bool)
	for i := 0; i < 101; i++ {
		DocA[i] = true
		DocB[i+1] = true
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		IntersectTwoMap(DocA, DocB)
	}
}
