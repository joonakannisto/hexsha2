package hexsha2_test 
import (
  "testing" 
  "github.com/joonakannisto/hexsha2"
)
func TestOut (t *testing.T) {
  if hexsha2.HexDigest("testing... 1 2 3") == "" {
    return t.Error("expected output")
  }
  if len(hexsha2.TruncHexDigest(5, "testing... 1 2 3") != 5 {
    return t.Error("length fails for truncated hex")
  }
  if len(hexsha2.TruncBase32Digest(5,"testing... 1 2 3") !=5 {
    return t.Error ("length fails for trunc base32")
  }
}
