package hexsha2

import "crypto/sha256"
import "hex"
import "encoding/base32"


func HexDigest(input string) string {
sum := sha256.Sum256([]byte(input))
return hex.EncodeToString(sum)
}
func TruncHexDigest(len int, input string) {
 sum := sha256.Sum256([]byte(input)) 
 if (len > 64 || len<0) {
  len=64
}
 return hex.EncodeToString(sum[:len/2])
}
func Base32Digest(input string) string {
sum := sha256.Sum256([]byte(input)) 
return base32.HexEncoding.EncodeToString(sum)
}
func TruncBase32Digest(len int, input string) {
 sum := sha256.Sum256([]byte(input)) 
 if (len > 52 || len<1) {
  len= 52
 }
last = (len*5)/8
return base32.HexEncoding.EncodeToString(sum[:last])
}
 

