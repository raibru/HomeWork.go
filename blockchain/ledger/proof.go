package ledger

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"log"
	"math"
	"math/big"
	"regexp"
	"strings"

	"github.com/raibru/blockchain/config"
)

// ProofOfWork data structure
type ProofOfWork struct {
	Block  *Block
	Target *big.Int
}

// CreateProof create a new proof of work
func CreateProof(b *Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-(config.GetConfig().Difficulty)))

	pow := &ProofOfWork{b, target}

	return pow
}

// InitData initialize a proof-of-work with nonce
func (pow *ProofOfWork) InitData(nonce int) []byte {
	data := bytes.Join(
		[][]byte{
			pow.Block.HashPrev,
			pow.Block.Data,
			ToHex(int64(nonce)),
			ToHex(int64(config.GetConfig().Difficulty)),
		},
		[]byte{},
	)

	return data
}

// Run execute the rule of the proof-of-work
func (pow *ProofOfWork) Run() (int, []byte) {
	var intHash big.Int
	var hash [32]byte

	nonce := 0

	for nonce < math.MaxInt64 {
		data := pow.InitData(nonce)
		hash = sha256.Sum256(data)

		fmt.Printf("\r%x", hash)
		intHash.SetBytes(hash[:])

		if intHash.Cmp(pow.Target) == -1 {
			printFoundHash(pow.Target, &intHash)
			break
		} else {
			nonce++
		}

	}
	fmt.Println()

	return nonce, hash[:]
}

// Validate the proof-of-work
func (pow *ProofOfWork) Validate() bool {
	var intHash big.Int

	data := pow.InitData(pow.Block.Nonce)

	hash := sha256.Sum256(data)
	intHash.SetBytes(hash[:])

	return intHash.Cmp(pow.Target) == -1
}

// ToHex convert an int64 value to a byte slice
func ToHex(num int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)

	}

	return buff.Bytes()
}

func printFoundHash(tg *big.Int, ih *big.Int) {
	ts := fmt.Sprintf("%v", *tg)
	is := fmt.Sprintf("%v", *ih)
	r, _ := regexp.Compile("[0-9]+ [0-9]+ [0-9]+ [0-9]+")
	ts = r.FindString(ts)
	is = r.FindString(is)
	tss := strings.Split(ts, " ")
	iss := strings.Split(is, " ")
	fmt.Printf("\n\nTarget    (big.Int): ")
	for _, e := range tss {
		fmt.Printf("%24s ", e)
	}
	fmt.Printf("\nData Hash (big.Int): ")
	for _, e := range iss {
		fmt.Printf("%24s ", e)
	}
	fmt.Printf("\n")
}
