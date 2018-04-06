// todo - make it a library!!
// todo - use return-by-ref/pointer (not by-value) for block



package main

import (
  "fmt"
  "crypto/sha256"
  "encoding/hex"
  "time"
  "strings"
  "encoding/binary"
)


type Block struct {
  Time          int64       // seconds since (unix) epoch (1970-01-01)
  Data          string      // use []byte too why? why not??
  Prev          []byte      // use []byte/int256/uint256 ??
  Difficulty    string
  Hash          []byte      // use []byteint256/uint256 ??
  Nonce         int64       // number used once - lucky (mining) lottery number
}


// bin(ary) bytes and integer number to (conversion) string helpers
func binToStr( bytes []byte ) string {
  return hex.EncodeToString( bytes )
}

func int64ToBin( num int64 ) []byte {
  b := make( []byte, 8 )
  binary.LittleEndian.PutUint64(b, uint64(num))
  return b
}





func calcHash( data []byte ) []byte {
  hashed := sha256.Sum256( data )
  return hashed[:]    // note: use [:] to convert [32]byte to slice ([]byte)
}



///
// fix: change data to []byte - why? why not?

func computeHashWithProofOfWork( data []byte, difficulty string ) (int64,[]byte) {
  nonce := int64( 0 )
  for {
    hash := calcHash( append( int64ToBin(nonce), data... ) )
    if strings.HasPrefix( binToStr(hash), difficulty ) {
        return nonce,hash    // bingo! proof of work if hash starts with leading zeros (00)
    } else {
        nonce += 1           // keep trying (and trying and trying)
    }
  }
}



func NewBlock( data string, prev []byte ) Block {
  t           := time.Now().Unix()
  difficulty  := "0000"

  // todo/check:
  //   "best" way to append bytes together ???
  nonce, hash := computeHashWithProofOfWork(
                   append(
                     append(
                       append(int64ToBin(t), []byte(difficulty)...),
                       prev...),
                     []byte(data)...),
                   difficulty )

  return Block { Time:       t,
                 Data:       data,
                 Prev:       prev,
                 Difficulty: difficulty,
                 Hash:       hash,
                 Nonce:      nonce }
}


func (b Block) Dump() {
  fmt.Println( b )
  fmt.Println( binToStr(b.Hash) )
  // use Printf( "%x" )
}



func main() {
  b0 := NewBlock( "Hello, Cryptos!", []byte("0000000000000000000000000000000000000000000000000000000000000000") )
  b1 := NewBlock( "Hello, Cryptos! - Hello, Cryptos!", b0.Hash )

  fmt.Println( b0 )
  // {1522691756 Hello, Cryptos!
  //    0000000000000000000000000000000000000000000000000000000000000000
  //    00009f597a8e28fc42a450c0ed2eff1b6507f76f6a7d1e112686700ce37e3676
  //    42278}
  fmt.Println( len( b0.Hash ))
  // => 32
  fmt.Println( len( b0.Prev ))
  // => 32

  b0.Dump()
  b1.Dump()

  fmt.Println( b1 )
  // {1522691756 Hello, Cryptos! - Hello, Cryptos!
  //     00009f597a8e28fc42a450c0ed2eff1b6507f76f6a7d1e112686700ce37e3676
  //     00009ef5ea432f840c3fb23dbedb5cce4c72e2951a140c1289dda1fedbcd6e99
  //     105106}

  fmt.Println( len( b1.Hash ))
  // => 32
  fmt.Println( len( b1.Prev ))
  // => 32

  blockchain := []Block {b0, b1}
  fmt.Println( blockchain )
  // => [{1522691756 Hello, Cryptos!
  //        0000000000000000000000000000000000000000000000000000000000000000
  //        00009f597a8e28fc42a450c0ed2eff1b6507f76f6a7d1e112686700ce37e3676
  //        42278}
  //     {1522691756 Hello, Cryptos! - Hello, Cryptos!
  //        00009f597a8e28fc42a450c0ed2eff1b6507f76f6a7d1e112686700ce37e3676
  //        00009ef5ea432f840c3fb23dbedb5cce4c72e2951a140c1289dda1fedbcd6e99
  //        105106}]
}
