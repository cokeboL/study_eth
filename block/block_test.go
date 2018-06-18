package main

import (
	"bytes"
	"testing"
)

func TestBlock(t *testing.T) {
	b1 := NewBlock(nil, "block 1")
	if !b1.IsValid() {
		t.Fatal("invalid block")
		return
	}

	b2 := NewBlock(b1, "block 2")

	b3 := NewBlock(b2, "block 3")

	b := &Block{}

	t.Log("-----------------------------------------------------------")
	t.Log("b1.Index: ", b1.Index)
	t.Log("b1.Timestamp: ", b1.Timestamp)
	t.Log("b1.DataLen: ", b1.DataLen)
	t.Log("b1.Data: ", b1.Data)
	t.Log("b1.PreHash: ", b1.PreHash)
	t.Log("b1.Hash: ", b1.Hash)
	t.Log("-----------------------------------------------------------")
	t.Log("b2.Index: ", b2.Index)
	t.Log("b2.Timestamp: ", b2.Timestamp)
	t.Log("b2.DataLen: ", b2.DataLen)
	t.Log("b2.Data: ", b2.Data)
	t.Log("b2.PreHash: ", b2.PreHash)
	t.Log("b2.Hash: ", b2.Hash)
	t.Log("-----------------------------------------------------------")
	t.Log("b3.Index: ", b3.Index)
	t.Log("b3.Timestamp: ", b3.Timestamp)
	t.Log("b3.DataLen: ", b3.DataLen)
	t.Log("b3.Data: ", b3.Data)
	t.Log("b3.PreHash: ", b3.PreHash)
	t.Log("b3.Hash: ", b3.Hash)
	t.Log("-----------------------------------------------------------")

	buf := bytes.NewBuffer(nil)
	t.Log("-----------------------------------------------------------")
	b1.WriteTo(buf)
	b.ReadFrom(buf)
	if *b != *b1 {
		t.Fatal("b != b1")
	}
	t.Log("-- b1.Index: ", b.Index)
	t.Log("-- b1.Timestamp: ", b.Timestamp)
	t.Log("-- b1.DataLen: ", b.DataLen)
	t.Log("-- b1.Data: ", b.Data)
	t.Log("-- b1.PreHash: ", b.PreHash)
	t.Log("-- b1.Hash: ", b.Hash)
	t.Log("-----------------------------------------------------------")
	b2.WriteTo(buf)
	b.ReadFrom(buf)
	if *b != *b2 {
		t.Fatal("b != b2")
	}
	t.Log("-- b2.Index: ", b.Index)
	t.Log("-- b2.Timestamp: ", b.Timestamp)
	t.Log("-- b2.DataLen: ", b.DataLen)
	t.Log("-- b2.Data: ", b.Data)
	t.Log("-- b2.PreHash: ", b.PreHash)
	t.Log("-- b2.Hash: ", b.Hash)
	t.Log("-----------------------------------------------------------")
	b3.WriteTo(buf)
	b.ReadFrom(buf)
	if *b != *b3 {
		t.Fatal("b != b3")
	}
	t.Log("-- b3.Index: ", b.Index)
	t.Log("-- b3.Timestamp: ", b.Timestamp)
	t.Log("-- b3.DataLen: ", b.DataLen)
	t.Log("-- b3.Data: ", b.Data)
	t.Log("-- b3.PreHash: ", b.PreHash)
	t.Log("-- b3.Hash: ", b.Hash)
	t.Log("-----------------------------------------------------------")
	t.Log("-- equal: ", *b == *b1, *b == *b2, *b == *b3)
}
