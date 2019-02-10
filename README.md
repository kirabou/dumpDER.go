# dumpDER.go - dump a DER file, displaying its structure and content

## Description
`dumpDER` is an example of using the `encoding/asn1` package to read and parse a DER file from `stdin`. It will display the DER file structure on `stdout` as well as the data it included, trying to improve the information displayed by converting the most current types (time, string, etc.) in a readable way and retrieving the name of object identifiers from the object identifier repository.

## Build
Make sure both files `dumpDER.go` and `object_name.go` are in the same empty directory, and use the `go build` command to build the `dumpDER` executable.

## Usage

Use the `dumpDER` command providing a DER file as its standard input. Example :

```
dumpDER < file.der
```
I tried it with smime.p7s files, as well as cert, request and keypair files. I took my test files from [X509 certificate examples for testing and verification](http://fm4dd.com/openssl/certexamples.htm)

```
$ ./dumpDER </512b-rsa-example-request.der
Note: all INTEGER, OCTET STRING and BIT STRING values displayed as hexadecimal bytes
SEQUENCE (260 bytes)                            :
| SEQUENCE (175 bytes)                          :
| | INTEGER (1 bytes)                           : 00
| | SEQUENCE (74 bytes)                         :
| | | SET (11 bytes)                            :
| | | | SEQUENCE (9 bytes)                      :
| | | | | OBJECT IDENTIFIER (3 bytes)           : 2.5.4.6 attributeType(4) countryName(6)
| | | | | PRINTABLE STRING (2 bytes)            : JP
| | | SET (14 bytes)                            :
| | | | SEQUENCE (12 bytes)                     :
| | | | | OBJECT IDENTIFIER (3 bytes)           : 2.5.4.8 attributeType(4) stateOrProvinceName(8)
| | | | | UTF8 STRING (5 bytes)                 : Tokyo
| | | SET (17 bytes)                            :
| | | | SEQUENCE (15 bytes)                     :
| | | | | OBJECT IDENTIFIER (3 bytes)           : 2.5.4.10 attributeType(4) organizationName(10)
| | | | | UTF8 STRING (8 bytes)                 : Frank4DD
| | | SET (24 bytes)                            :
| | | | SEQUENCE (22 bytes)                     :
| | | | | OBJECT IDENTIFIER (3 bytes)           : 2.5.4.3 attributeType(4) commonName(3)
| | | | | UTF8 STRING (15 bytes)                : www.example.com
| | SEQUENCE (92 bytes)                         :
| | | SEQUENCE (13 bytes)                       :
| | | | OBJECT IDENTIFIER (9 bytes)             : 1.2.840.113549.1.1.1 pkcs-1(1) rsaEncryption(1)
| | | | NULL (0 bytes)                          : NUL
| | | BIT STRING (75 bytes)                     : 00 30 48 02 41 00 9B FC 66 90 79 84 42 BB AB 13
| | | |                                           FD 2B 7B F8 DE 15 12 E5 F1 93 E3 06 8A 7B B8 B1
| | | |                                           E1 9E 26 BB 95 01 BF E7 30 ED 64 85 02 DD 15 69
| | | |                                           A8 34 B0 06 EC 3F 35 3C 1E 1B 2B 8F FA 8F 00 1B
| | | |                                           DF 07 C6 AC 53 07 02 03 01 00 01
| | 0 (0 bytes)                                 :
| SEQUENCE (13 bytes)                           :
| | OBJECT IDENTIFIER (9 bytes)                 : 1.2.840.113549.1.1.5 pkcs-1(1) sha1-with-rsa-signature(5)
| | NULL (0 bytes)                              : NUL
| BIT STRING (65 bytes)                         : 00 72 39 5E 76 63 5E F2 F3 1C 35 57 FC 6F AE ED
| |                                               EB 2C FA D5 C5 80 17 4B 94 A0 BC DA 5F 06 C8 F7
| |                                               F2 53 55 B5 3B EE 1F F3 20 AE 80 60 9A 34 A9 9E
| |                                               A2 AA 06 20 43 92 86 36 61 41 13 DA A9 86 8C 0B
| |                                               BD
```


## Known limitations

1. The `encoding/asn1`package only implements a subset of ASN.1. Not all ASN.1 types are recognized. The list of the ASN.1 types than can be recognized by the `encoding/asn1` package are described in the [Unmarshal function documentation](https://golang.org/pkg/encoding/asn1/#Unmarshal)
2. To retrieve the name of an ASN.1 object identifier, I am using a GET from the [OID Repository Web site](http://oid-info.com/). It's not very efficient, especially for DER file with a lot of objects, and should be optimized &#10132; this is now fixed: `dumpDER.gp` manages a map of object identifiers names. It now knows the most often used objects identifiers and their names. When a new objectif identifier is found, its name is added to the map to avoid multiple requests to oid-info.com
3. Some objects use OCTET STRING or BIT STRING as extension for other ASN.1 data, as described [here](https://stackoverflow.com/questions/15299201/asn-1-octet-strings). These extensions are currently not parsed as ASN.1 data and only displayed as hexadecimal bytes  &#10132; as a partial fix, `dumpDER` now tries to parse OCTET STRING and BIT STRING that start like a SEQUENCE


