package holocron

import (
	"os"
	"fmt"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"path/filepath"
	"strings"
	b64 "encoding/base64"
	"github.com/skip2/go-qrcode"
	"github.com/nogoegst/balloon"

)

type Holocron struct {
	Name string `json:"name"`
	Gatekeeper string `json:"gatekeeper"`
	Ascertainment string `json:"ascertainment"`
	Treasure string `json:"treasure"`
	Salt string `json:"salt"`
}

func NewHolocron(name, gatekeeper, ascertainment, treasure, salt string) Holocron{
	return Holocron{
		Name: name,
		Gatekeeper: gatekeeper,
		Ascertainment: ascertainment,
		Salt: salt,
		Treasure: treasure,
	}
}

func (h Holocron) Forge(outdir string) error {
	cipher := Encrypt(h.Ascertainment, h.Treasure, h.Salt)
	message := strings.ReplaceAll(fmt.Sprintf(
		`%s

        Decode the following cipher, using the answer to the above prompt, to reveal the treasure:

        %s
        - [salt]-[nonce]-[encrypted]
        - Sha256 Balloon Hash (4096, 32) Answer to obtain the AES-GCM decoding key https://web.archive.org/details/https://en.wikipedia.org/wiki/Balloon_hashing
        `,
		h.Gatekeeper,
		cipher,
	), "  ", " ")
	encoded := generateBase64Encoding(h.Name, message)
	return toQrCodeFile(filepath.Join(outdir, fmt.Sprintf("%s.png", h.Name)), encoded)
}

func (h Holocron) ForgeForWeb() (string, error) {
	cipher := Encrypt(h.Ascertainment, h.Treasure, h.Salt)
	message := strings.ReplaceAll(fmt.Sprintf(
		`%s

        Decode the following cipher, using the answer to the above prompt, to reveal the treasure:

        %s

        Decode using https://holocron.algo.xyz or the algorithm below:
        - [salt]-[nonce]-[encrypted]
        - Sha256 Balloon Hash (4096, 32) Answer to obtain the AES-GCM decoding key https://web.archive.org/details/https://en.wikipedia.org/wiki/Balloon_hashing
        `,
		h.Gatekeeper,
		cipher,
	), "  ", " ")
	encoded := generateBase64Encoding(h.Name, message)
	return toQrCodeBase64(encoded)
}

func (h Holocron) ForgeLightForWeb() (string, error) {
	cipher := Encrypt(h.Ascertainment, h.Treasure, h.Salt)
	message := strings.ReplaceAll(fmt.Sprintf(
		`%s :
		
        %s  (Decode: https://holocron.algo.xyz)
        `,
		h.Gatekeeper,
		cipher,
	), "  ", " ")
	encoded := generateBase64Encoding(h.Name, message)
	return toQrCodeBase64(encoded)
}

func deriveKey(passphrase string, salt []byte) ([]byte, []byte) {
	if salt == nil {
		salt = make([]byte, 8)
		// http://www.ietf.org/rfc/rfc2898.txt
		// Salt.
		rand.Read(salt)
	}
	// return pbkdf2.Key([]byte(passphrase), salt, 1000, 32, sha256.New), salt
	return balloon.Balloon(sha256.New(), []byte(passphrase), salt, 4096, 32), salt
}

func Encrypt(passphrase, plaintext, saltString string) string {
	var key []byte 
	var salt []byte
	if len(saltString) > 0 {
		key, salt = deriveKey(passphrase, []byte(saltString))
	} else {
		key, salt = deriveKey(passphrase, nil)
	}
	iv := make([]byte, 12)
	// http://nvlpubs.nist.gov/nistpubs/Legacy/SP/nistspecialpublication800-38d.pdf
	// Section 8.2
	rand.Read(iv)
	b, _ := aes.NewCipher(key)
	aesgcm, _ := cipher.NewGCM(b)
	data := aesgcm.Seal(nil, iv, []byte(plaintext), nil)
	return hex.EncodeToString(salt) + "-" + hex.EncodeToString(iv) + "-" + hex.EncodeToString(data)
}

func Decrypt(passphrase, ciphertext string) string {
	arr := strings.Split(ciphertext, "-")
	salt, _ := hex.DecodeString(arr[0])
	iv, _ := hex.DecodeString(arr[1])
	data, _ := hex.DecodeString(arr[2])
	key, _ := deriveKey(passphrase, salt)
	b, _ := aes.NewCipher(key)
	aesgcm, _ := cipher.NewGCM(b)
	data, _ = aesgcm.Open(nil, iv, data, nil)
	return string(data)
}

func generateBase64Encoding(name, cipher string) string {
	sEnc := b64.StdEncoding.EncodeToString([]byte(cipher))
	return fmt.Sprintf(
		"data:text/plain;content-disposition=attachment;filename=%s.txt;base64,%s",
		name,
		sEnc,
	)
}

func toQrCodeFile(filename, message string) error {
	var png []byte
  	png, err := qrcode.Encode(message, qrcode.Medium, 256)
  	if err != nil {
  		return err
  	}
  	return os.WriteFile(filename, png, 0644)
}

func toQrCodeBase64(message string) (string, error) {
	var png []byte
  	png, err := qrcode.Encode(message, qrcode.Medium, 256)
  	if err != nil {
  		return "", err
  	}
  	return "data:image/png;base64,"+ b64.StdEncoding.EncodeToString(png), nil
}