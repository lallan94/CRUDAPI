package auth

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"errors"
)

func CreatePublicKey(pubKey string) (*rsa.PublicKey, error) {
	publicKey, err := base64.StdEncoding.DecodeString(pubKey)
	if err != nil {
		return nil, err
	}
	publickey, err := x509.ParsePKIXPublicKey(publicKey)
	if err != nil {
		return nil, err
	}
	switch pub := publickey.(type) {
	case *rsa.PublicKey:
		return pub, nil
	default:
		break
	}
	return nil, errors.New("key type is not RSA")
}

// func CreatePrivateKey(privatekey string) (*rsa.PrivateKey, error) { privKey, err := base64.StdEncoding.DecodeString(privatekey) if err != nil { return nil, err } privateKey, err1 := x509.ParsePKCS1PrivateKey(privKey) if err != nil { return nil, err1 } // switch priv := privateKey.(type) { // case *rsa.PrivateKey: // return priv, nil // default: // break // fall through // } if err != nil { return nil, err } return privateKey, nil //return nil, errors.New("key type is not RSA")}func decryptData(priv *rsa.PrivateKey, ciphertext []byte) ([]byte, error) { plaintext, err := decrypt(getHash(), priv, ciphertext, nil) if err != nil { return nil, err } return plaintext, err}func encryptData(pub *rsa.PublicKey, plaintext []byte) ([]byte, error) { ciphertext, err := encrypt(getHash(), pub, plaintext, nil) if err != nil { return nil, err } return ciphertext, err}// TODO:update current hashing function with sha512.new() if required// Sign with private key with sign the data with the sever private key us sha256 hashingfunc signData(privateKey *rsa.PrivateKey, data []byte) (string, error) { dataHashSum := sha512.Sum512(data) signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA512, dataHashSum[:]) if err != nil { return "", err } return base64.StdEncoding.EncodeToString(signature), err}// Verify Sign will verify the sign with device public keys and return err if there is a// validation error.func verifySign(publicKey *rsa.PublicKey, data []byte, signature string) error { sign, err := base64.StdEncoding.DecodeString(signature) if err != nil { return err } dataHashSum := sha512.Sum512(data) return rsa.VerifyPKCS1v15(publicKey, crypto.SHA512, dataHashSum[:], sign)}// Encrypts data with the availabe hash function and return encrypted byte array and errorfunc encrypt(hash hash.Hash, public *rsa.PublicKey, msg []byte, label []byte) ([]byte, error) { length := len(msg) size := public.Size() - 11 var encryptedData []byte for i := 0; i < length; i += size { finish := i + size if finish > length { finish = length } encrypted, err := rsa.EncryptPKCS1v15(rand.Reader, public, msg[i:finish]) if err != nil { return nil, err } encryptedData = append(encryptedData, encrypted...) } return encryptedData, nil}// Decrypts data with the availabe pricate key and return decrypted byte array and error.func decrypt(hash hash.Hash, private *rsa.PrivateKey, msg []byte, label []byte) ([]byte, error) { length := len(msg) size := private.PublicKey.Size() var decryptedData []byte for i := 0; i < length; i += size { finish := i + size if finish > length { finish = length } decrypted, err := rsa.DecryptPKCS1v15(rand.Reader, private, msg[i:finish]) if err != nil { return nil, err } decryptedData = append(decryptedData, decrypted...) } return decryptedData, nil}func getHash() hash.Hash { return sha256.New()}
