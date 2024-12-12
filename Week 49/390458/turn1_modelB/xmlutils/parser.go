package xmlutils

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"strings"
)

// PrettyPrint prints XML data in a pretty-printed format
func PrettyPrint(v interface{}) {
	var out bytes.Buffer
	enc := xml.NewEncoder(&out)
	enc.Indent("", "  ")
	err := enc.Encode(v)
	if err != nil {
		fmt.Println("Error pretty-printing XML:", err)
		return
	}
	fmt.Println(strings.TrimSpace(out.String()))
}

// ParseXML parses XML from a reader
func ParseXML(reader io.Reader, v interface{}) error {
	return xml.NewDecoder(reader).Decode(v)
}
