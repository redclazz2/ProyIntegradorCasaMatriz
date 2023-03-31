package utils

import (
	"crypto/sha512"
	"encoding/hex"
)

func Sha512hash(texto string) string {
	// Crear un nuevo hash SHA-512
	hash := sha512.New()

	// Escribir los datos en el hash
	hash.Write([]byte(texto))

	// Calcular el hash resultante y convertirlo a una cadena
	hashResult := hex.EncodeToString(hash.Sum(nil))

	return hashResult
}
