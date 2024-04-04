package pdfb

import (
	"fmt"
	"log"
	"os"
	"strings"
)

// Used to decode base64 encoded string
// func decode(b64Str string) (b []byte) {
// 	b, err := base64.StdEncoding.DecodeString(b64Str)
// 	if err != nil {
// 		log.ErrorFatal("%s", err)
// 	}
// 	return
// }

func fileExists(f string) bool {
	if _, err := os.Stat(f); os.IsNotExist(err) {
		return false
	}
	return true
}

// Used to report any errors or display a success message
func (p *Pdfb) checkpoint(str string) {
	if p.pdf.Err() {
		log.Fatalln(p.pdf.Error().Error())
	} else {
		fmt.Println("-- Checkpoint:", str)
	}
}

// used to generate an align string
func (p *Pdfb) makeAlignStr(alignInput string) (alignStr string) {
	alignInput = strings.ToLower(alignInput)

	switch {
	case alignInput == "l" || alignInput == "left":
		alignStr = "L"
	case alignInput == "c" || alignInput == "centre":
		alignStr = "C"
	case alignInput == "r" || alignInput == "right":
		alignStr = "R"
	default:
		log.Fatalf("Invalid align input (%s)\n", alignInput)
	}

	return
}
