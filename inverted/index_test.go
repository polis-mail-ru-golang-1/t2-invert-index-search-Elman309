package inverted

import (
	"reflect"
	"testing"
)

func TestUpdate(t *testing.T) {
	expected := Index{
		"word": {
			"doc1": 1,
			"doc2": 2,
		},
	}
	got := NewIndex()
	got.Update("doc1", "word")
	got.Update("doc2", "word")
	got.Update("doc2", "word")
	if !reflect.DeepEqual(expected, got) {
		t.Error("Index.Update invalid result")
	}
}

func TestProcessQuery(t *testing.T) {
	InitTokenize()

	idx := Index{
		"word1": {
			"doc1": 1,
			"doc2": 2,
		},
		"word2": {
			"doc2": 3,
		},
	}
	expected := Documents{
		"doc2": 5,
		"doc1": 1,
	}
	got := idx.ProcessQuery("word1 word2")

	if !reflect.DeepEqual(expected, got) {
		t.Error("expected", expected, "got", got)
	}
}

func TestMerge(t *testing.T) {
	idx1 := Index{
		"word1": {
			"doc1": 1,
			"doc2": 2,
		},
		"word2": {
			"doc2": 3,
		},
	}
	idx2 := Index{
		"word1": {
			"doc3": 3,
			"doc4": 4,
		},
		"word2": {
			"doc1": 32,
		},
	}
	got := NewIndex()
	got.Merge(idx1)
	got.Merge(idx2)
	expected := Index{
		"word1": {
			"doc1": 1,
			"doc2": 2,
			"doc3": 3,
			"doc4": 4,
		},
		"word2": {
			"doc1": 32,
			"doc2": 3,
		},
	}

	if !reflect.DeepEqual(expected, got) {
		t.Error("expected", expected, "got", got)
	}
}
