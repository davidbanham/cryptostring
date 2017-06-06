PACKAGE DOCUMENTATION

package cryptostring
    import "github.com/davidbanham/cryptostring"


FUNCTIONS

func SpecifiedLength(numbytes int) (string, error)
    SpecifiedLength returns a URL-safe, base64 encoded securely generated
    random string with the specified number of bytes.

func String() (string, error)
    String returns a URL-safe, base64 encoded securely generated random
    string with 18 bytes of entropy and 24 characters in length


