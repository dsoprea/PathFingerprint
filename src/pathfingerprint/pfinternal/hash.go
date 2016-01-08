package pfinternal

import (
    "fmt"
    "errors"
    "hash"
    "crypto/sha1"
    "crypto/sha256"
)

const (
    Sha1Algorithm = "sha1"
    Sha256Algorithm = "sha256"
)

func getHashObject(algorithmName *string) (hash.Hash, error) {
    if *algorithmName == Sha1Algorithm {
        return sha1.New(), nil
    } else if *algorithmName == Sha256Algorithm {
        return sha256.New(), nil
    } else {
        message := fmt.Sprintf("Hash algorithm [%s] is not valid/supported", algorithmName)
        err := errors.New(message)
        return nil, err
    }
}
