package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"github.com/lvlcainiao/framework/util"
	"io"
	"time"
)

const (
	BLOCK_HEAD_SIZE = 20

	BLOCK_HASH_SIZE = 64
)

var (
	ZERO_HASH_STR = string(bytes.Repeat([]byte{'0'}, BLOCK_HASH_SIZE))
)

type Block struct {
	Index     uint64
	Timestamp uint64
	DataLen   uint32
	Data      string
	PreHash   string
	Hash      string
}

func (b *Block) Calculate() string {
	str := fmt.Sprintf("%d%d%d%s%s", b.Index, b.Timestamp, b.DataLen, b.Data, b.PreHash)
	h := sha256.New()
	h.Write(util.StrToBytes(str))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

func (b *Block) IsValid() bool {
	return b.Hash == b.Calculate()
}

func (b *Block) WriteTo(w io.Writer) {
	binary.Write(w, binary.LittleEndian, &b.Index)
	binary.Write(w, binary.LittleEndian, &b.Timestamp)
	binary.Write(w, binary.LittleEndian, &b.DataLen)
	binary.Write(w, binary.LittleEndian, util.StrToBytes(b.Data))
	binary.Write(w, binary.LittleEndian, util.StrToBytes(b.PreHash))
	binary.Write(w, binary.LittleEndian, util.StrToBytes(b.Hash))
}

func (b *Block) ReadFrom(r io.Reader) {
	binary.Read(r, binary.LittleEndian, &b.Index)
	binary.Read(r, binary.LittleEndian, &b.Timestamp)
	binary.Read(r, binary.LittleEndian, &b.DataLen)

	dataLen := b.DataLen
	if b.DataLen < BLOCK_HASH_SIZE {
		dataLen = BLOCK_HASH_SIZE
	}

	buf := make([]byte, dataLen)
	binary.Read(r, binary.LittleEndian, buf[:b.DataLen])
	b.Data = string(buf[:b.DataLen])

	binary.Read(r, binary.LittleEndian, buf[:BLOCK_HASH_SIZE])
	b.PreHash = string(buf[:BLOCK_HASH_SIZE])

	binary.Read(r, binary.LittleEndian, buf[:BLOCK_HASH_SIZE])
	b.Hash = string(buf[:BLOCK_HASH_SIZE])
}

// func (b *Block) Equal(other *Block) bool {
// 	return b.Index == other.Index && b.Timestamp == other.Timestamp && b.DataLen == other.DataLen &&
// 		b.Data == other.Data && b.PreHash == other.PreHash && b.Hash == other.Hash
// }

func NewBlock(old *Block, data string) *Block {
	b := &Block{
		Timestamp: uint64(time.Now().UnixNano()),
		DataLen:   uint32(len(data)),
		Data:      data,
	}

	if old != nil {
		b.Index = old.Index + 1
		b.PreHash = old.Hash
	} else {
		b.PreHash = ZERO_HASH_STR
	}

	b.Hash = b.Calculate()

	return b
}

func main() {
	b1 := NewBlock(nil, "block 1")
	if !b1.IsValid() {
		fmt.Println("invalid block")
		return
	}

	b2 := NewBlock(b1, "block 2")

	buf := bytes.NewBuffer(nil)

	b3 := NewBlock(b2, "block 3")

	b := &Block{}

	fmt.Println("-----------------------------------------------------------")
	fmt.Println("b1.Index: ", b1.Index)
	fmt.Println("b1.Timestamp: ", b1.Timestamp)
	fmt.Println("b1.DataLen: ", b1.DataLen)
	fmt.Println("b1.Data: ", b1.Data)
	fmt.Println("b1.PreHash: ", b1.PreHash)
	fmt.Println("b1.Hash: ", b1.Hash)
	fmt.Println("-----------------------------------------------------------")
	fmt.Println("b2.Index: ", b2.Index)
	fmt.Println("b2.Timestamp: ", b2.Timestamp)
	fmt.Println("b2.DataLen: ", b2.DataLen)
	fmt.Println("b2.Data: ", b2.Data)
	fmt.Println("b2.PreHash: ", b2.PreHash)
	fmt.Println("b2.Hash: ", b2.Hash)
	fmt.Println("-----------------------------------------------------------")
	fmt.Println("b3.Index: ", b3.Index)
	fmt.Println("b3.Timestamp: ", b3.Timestamp)
	fmt.Println("b3.DataLen: ", b3.DataLen)
	fmt.Println("b3.Data: ", b3.Data)
	fmt.Println("b3.PreHash: ", b3.PreHash)
	fmt.Println("b3.Hash: ", b3.Hash)
	fmt.Println("-----------------------------------------------------------")

	fmt.Println("-----------------------------------------------------------")
	b1.WriteTo(buf)
	b.ReadFrom(buf)
	fmt.Println("-- b1.Index: ", b.Index)
	fmt.Println("-- b1.Timestamp: ", b.Timestamp)
	fmt.Println("-- b1.DataLen: ", b.DataLen)
	fmt.Println("-- b1.Data: ", b.Data)
	fmt.Println("-- b1.PreHash: ", b.PreHash)
	fmt.Println("-- b1.Hash: ", b.Hash)
	fmt.Println("-----------------------------------------------------------")
	b2.WriteTo(buf)
	b.ReadFrom(buf)
	fmt.Println("-- b2.Index: ", b.Index)
	fmt.Println("-- b2.Timestamp: ", b.Timestamp)
	fmt.Println("-- b2.DataLen: ", b.DataLen)
	fmt.Println("-- b2.Data: ", b.Data)
	fmt.Println("-- b2.PreHash: ", b.PreHash)
	fmt.Println("-- b2.Hash: ", b.Hash)
	fmt.Println("-----------------------------------------------------------")
	b3.WriteTo(buf)
	b.ReadFrom(buf)
	fmt.Println("-- b3.Index: ", b.Index)
	fmt.Println("-- b3.Timestamp: ", b.Timestamp)
	fmt.Println("-- b3.DataLen: ", b.DataLen)
	fmt.Println("-- b3.Data: ", b.Data)
	fmt.Println("-- b3.PreHash: ", b.PreHash)
	fmt.Println("-- b3.Hash: ", b.Hash)
	fmt.Println("-----------------------------------------------------------")
}
