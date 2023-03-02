package test

import (
	"github.com/auth-server-go/utils"
	"testing"
	"hash/fnv"
)

func TestHash(t *testing.T) {
	s := "Ronit123"
	got := utils.Hash(s)
	h := fnv.New32a()
	h.Write([]byte(s))
	want := h.Sum32()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}