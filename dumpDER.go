package main

import (
	"encoding/asn1"
	"fmt"
	"golang.org/x/net/html"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)



// WidthFieldNameColumn is the width of the left column used to display
// data field name. You can change it to any value you like
var WidthFieldNameColumn int = 48

// GetHtmlTitle returns the title of a DOM tree, or an empty string
func GetHtmlTitle(n *html.Node) string {

	if n.Type == html.ElementNode && n.Data == "title" {
		return n.FirstChild.Data
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		s := GetHtmlTitle(c)
		if len(s) > 0 {
			return s
		}
	}

	return ""
}

// GetOIName retrieves the name of an object identifier from the oid-info.com Web site.
// Only the last two parts of the full name are kept to avoid names to be too long and
// difficult to read.
func GetOIName(oi string) string {

	// Check if we already know this object identifier from our MapOfObjects
	on, ok := MapOfObjects[oi]
	if ok {
		return on.Name
	} else {
		MapOfObjects[oi] = ObjectName{}
	}

	// We don't know this object identifier yet, so we will retrieve it from
	// oid-info.com and add it to the MapOfObjects
	req, err := http.NewRequest("GET", fmt.Sprintf("http://oid-info.com/get/%s", oi), nil)
	if err != nil {
		log.Println("Cannot build request for oid-info failed -", err)
		return ""
	}

	// A decent user agent is required to access oid-info.com
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/60.0.3112.113 Safari/537.36")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("Get from oid-info failed -", err)
		return ""
	}
	defer res.Body.Close()

	doc, err := html.Parse(res.Body)
	if err != nil {
		log.Println("HTML parse from oid-info failed -", err)
		return ""
	}

	s := GetHtmlTitle(doc)

	// We need to split the title to keep the last 2 parts only
	f := func(c rune) bool {
		return c == ' ' || c == '{' || c == '}'
	}
	split := strings.FieldsFunc(s, f)

	switch l := len(split); l {
	case 0:
		return ("")
	case 1:
		MapOfObjects[oi] = ObjectName{fmt.Sprint("%s", split[0]), false}
		return MapOfObjects[oi].Name
	default:
		MapOfObjects[oi] = ObjectName{fmt.Sprintf("%s %s", split[l-2], split[l-1]), false}
		return MapOfObjects[oi].Name
	}

}

// PrintHex dumps a byte slice as hexadecimal values, with width and left margin
// of the dump given as parameters
func PrintHex(data []byte, prefix string, width int, margin int) {

	if len(data) == 0 {
		fmt.Printf("NUL")
		return
	}

	var i int
	var b byte
	for i, b = range data {
		switch i % width {
		case 0:
			if i == 0 {
				fmt.Printf("%02X", b)
			} else {
				// fmt.Printf(" %s%02X", strings.Repeat(" ", margin), b)
				fmt.Printf("%s|%s%02X", prefix, strings.Repeat(" ", margin-len(prefix)), b)
			}
		case 1:
			fmt.Printf(" %02X", b)
		case width - 1:
			if i == len(data)-1 {
				fmt.Printf(" %02X", b)
			} else {
				fmt.Printf(" %02X\n", b)
			}
		default:
			fmt.Printf(" %02X", b)
		}
	}

}

// GetStringFromTag returns the full name of an ASN.1 type
func GetStringFromTag(tag int) string {

	switch tag {
	case asn1.TagBoolean:
		return ("BOOLEAN")
	case asn1.TagInteger:
		return ("INTEGER")
	case asn1.TagBitString:
		return ("BIT STRING")
	case asn1.TagOctetString:
		return ("OCTET STRING")
	case asn1.TagNull:
		return ("NULL")
	case asn1.TagOID:
		return ("OBJECT IDENTIFIER")
	case asn1.TagEnum:
		return ("ENUM")
	case asn1.TagUTF8String:
		return ("UTF8 STRING")
	case asn1.TagSequence:
		return ("SEQUENCE")
	case asn1.TagSet:
		return ("SET")
	case asn1.TagNumericString:
		return ("NUMERIC STRING")
	case asn1.TagPrintableString:
		return ("PRINTABLE STRING")
	case asn1.TagT61String:
		return ("T61String")
	case asn1.TagIA5String:
		return ("IA5String")
	case asn1.TagUTCTime:
		return ("UTCTime")
	case asn1.TagGeneralizedTime:
		return ("GeneralizedTime")
	case asn1.TagGeneralString:
		return ("GENERAL STRING")
	}

	return fmt.Sprintf("%d", tag)

}

// GetAsnValueAsString converts the raw value of parsed ASN1 data as
// something more readable depending of the ASN1 type value of the
// data read
func GetAsnValueAsString(asn *asn1.RawValue) string {

	if asn == nil {
		return ""
	}

	switch asn.Tag {
	case asn1.TagOID:
		var oi asn1.ObjectIdentifier
		_, err := asn1.Unmarshal(asn.FullBytes, &oi)
		if err != nil {
			log.Fatalln("Erreur unmarshalling -", err)
		}
		s := oi.String()
		return fmt.Sprintf("%s %s", s, GetOIName(s))
	case asn1.TagPrintableString, asn1.TagIA5String, asn1.TagNumericString, asn1.TagUTF8String:
		var asn_string string
		_, err := asn1.Unmarshal(asn.FullBytes, &asn_string)
		if err != nil {
			log.Fatalln("Erreur unmarshalling -", err)
		}
		return asn_string
	case asn1.TagUTCTime, asn1.TagGeneralizedTime:
		var t time.Time
		_, err := asn1.Unmarshal(asn.FullBytes, &t)
		if err != nil {
			log.Fatalln("Erreur unmarshalling -", err)
		}
		return t.String()
	case asn1.TagBoolean:
		var b bool
		_, err := asn1.Unmarshal(asn.FullBytes, &b)
		if err != nil {
			log.Fatalln("Erreur unmarshalling -", err)
		}
		if b {
			return ("true")
		} else {
			return ("false")
		}

	}

	return ""

}

// PrintFieldName prints on stdio the name of an ASN1 data field, making sure to
// stay in the WidthFieldNameColumn defined
func PrintFieldName(s string) {
	fmt.Printf("\n%-*.*s: ", WidthFieldNameColumn, WidthFieldNameColumn, s)
}

// Parse parses a slice of bytes as ASN1 data. It runs recursively to manage
// nested data enty
func Parse(data []byte, index int) {

	if len(data) == 0 {
		return
	}

	var asn asn1.RawValue
	var data2 []byte = data

	for {

		rest, err := asn1.Unmarshal(data2, &asn)
		if err != nil {
			log.Printf("Erreur unmarshalling - %v\n", err)
			log.Printf("\n% X\n", data)
			break
		}

		// log.Printf("[%02d] %s (%d bytes) IsCompound:%v Rest:%d", index, GetStringFromTag(asn.Tag), len(asn.Bytes), asn.IsCompound, len(rest))
		PrintFieldName(fmt.Sprintf("%s%s (%d bytes)", strings.Repeat("| ", index-1), GetStringFromTag(asn.Tag), len(asn.Bytes)))

		if asn.IsCompound {
			Parse(asn.Bytes, index+1)
		}

		if rest == nil {
			fmt.Println("-- no more rest")
			break
		}

		if !asn.IsCompound {
			s := GetAsnValueAsString(&asn)
			if len(s) == 0 {
				PrintHex(asn.Bytes, fmt.Sprintf("%s", strings.Repeat("| ", index-1)), 16, WidthFieldNameColumn+1)
			} else {
				fmt.Printf("%s", s)
			}
		}

		// log.Printf("[%02d] %s (%d bytes) IsCompound:%v Rest:%d", index, GetStringFromTag(asn.Tag), len(asn.Bytes), asn.IsCompound, len(rest))

		// Stop when no rest left, or rest is a line feed character
		if len(rest) == 0 || (len(rest) == 1 && rest[0] == '\x0a') {
			break
		} else {
			data2 = rest
		}

	}

}

// dumpDER is a Go program to read a DER file from stdin and display its structure and content
// in a readable way on stdio. Based on the Golang encoding/asn1 package to parse the DER file
func main() {

	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// Read DER file from Stdin
	der, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatalln("Erreur lecture stdin -", err)
	}

	fmt.Printf("Note: all INTEGER, OCTET STRING and BIT STRING values displayed as hexadecimal bytes")

	Parse(der, 1)

	fmt.Println()

	// log.Printf("%#v", MapOfObjects)

}
