package main 

import (
  "github.com/joonakannisto/hexsha2"
  "fmt"
)

func main() {
  fmt.Println(hexsha2.HexDigest("testing"))
  fmt.Println(hexsha2.Base32Digest("testing"))
  fmt.Println(hexsha2.TruncHexDigest(8,"testing"))
  fmt.Println(hexsha2.TruncBase32Digest(8,"testing"))
}
