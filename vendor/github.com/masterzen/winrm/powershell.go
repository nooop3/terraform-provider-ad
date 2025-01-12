package winrm

import (
	"encoding/base64"
	"golang.org/x/text/encoding/unicode"
	"fmt"
)

// Powershell wraps a PowerShell script
// and prepares it for execution by the winrm client
func Powershell(psCmd string) string {
	// 2 byte chars to make PowerShell happy
	// wideCmd := ""
	// for _, b := range []byte(psCmd) {
		// wideCmd += string(b) + "\x00"
	// }

	// // Base64 encode the command
	// input := []uint8(wideCmd)
	// encodedCmd := base64.StdEncoding.EncodeToString(input)

	encoder := unicode.UTF16(unicode.LittleEndian, unicode.IgnoreBOM).NewEncoder()
	encoded, err := encoder.String(psCmd)
	if err != nil {
		return ""
	}
	// Finally make it base64 encoded which is required for powershell.
	encodedCmd := base64.StdEncoding.EncodeToString([]byte(encoded))

	// Create the powershell.exe command line to execute the script
	return fmt.Sprintf("powershell.exe -EncodedCommand %s", encodedCmd)
}
