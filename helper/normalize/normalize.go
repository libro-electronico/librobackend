package normalize

import (
	"regexp"
	"strings"
)

// processInput memproses input menjadi huruf kecil, menghapus spasi,
// dan hanya memperbolehkan karakter a-z, 0-9, _ dan -
func SetIntoID(input string) string {
	// Konversi ke huruf kecil
	input = strings.ToLower(input)

	// Hapus spasi
	input = strings.ReplaceAll(input, " ", "")

	// Hapus karakter khusus kecuali _ dan -
	re := regexp.MustCompile(`[^a-z0-9_-]`)
	input = re.ReplaceAllString(input, "")

	return input
}
