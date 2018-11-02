package main


type ObjectName struct {
	Name          string
	HaveExtension bool
}

var MapOfObjects = map[string]ObjectName{
	"2.5.4.11": ObjectName{
		Name:          "attributeType(4) organizationalUnitName(11)",
		HaveExtension: false,
	},
	"2.16.840.1.101.3.4.1.22": ObjectName{
		Name:          "aes(1) aes192-CBC(22)",
		HaveExtension: false,
	},
	"2.5.4.3": ObjectName{
		Name:          "attributeType(4) commonName(3)",
		HaveExtension: false,
	},
	"2.5.29.14": ObjectName{
		Name:          "certificateExtension(29) subjectKeyIdentifier(14)",
		HaveExtension: false,
	},
	"1.2.840.113549.1.7.1": ObjectName{
		Name:          "pkcs-7(7) data(1)",
		HaveExtension: false,
	},
	"1.2.840.113549.1.9.1": ObjectName{
		Name:          "pkcs-9(9) emailAddress(1)",
		HaveExtension: false,
	},
	"2.16.840.1.101.3.4.1.42": ObjectName{
		Name:          "aes(1) aes256-CBC(42)",
		HaveExtension: false,
	},
	"1.3.14.3.2.7": ObjectName{
		Name:          "algorithms(2) desCBC(7)",
		HaveExtension: false,
	},
	"2.5.29.35": ObjectName{
		Name:          "certificateExtension(29) authorityKeyIdentifier(35)",
		HaveExtension: false,
	},
	"2.5.29.17": ObjectName{
		Name:          "certificateExtension(29) subjectAltName(17)",
		HaveExtension: false,
	},
	"1.2.840.113549.1.9.5": ObjectName{
		Name:          "pkcs-9(9) signing-time(5)",
		HaveExtension: false,
	},
	"1.2.840.113549.1.1.11": ObjectName{
		Name:          "pkcs-1(1) sha256WithRSAEncryption(11)",
		HaveExtension: false,
	},
	"2.5.29.32": ObjectName{
		Name:          "certificateExtension(29) certificatePolicies(32)",
		HaveExtension: false,
	},
	"1.2.840.113549.3.2": ObjectName{
		Name:          "encryptionalgorithm(3) rc2-cbc(2)",
		HaveExtension: false,
	},
	"1.2.840.113549.1.7.2": ObjectName{
		Name:          "pkcs-7(7) signedData(2)",
		HaveExtension: false,
	},
	"2.5.4.7": ObjectName{
		Name:          "attributeType(4) localityName(7)",
		HaveExtension: false,
	},
	"1.2.840.113549.1.1.5": ObjectName{
		Name:          "pkcs-1(1) sha1-with-rsa-signature(5)",
		HaveExtension: false,
	},
	"2.16.840.1.101.3.4.1.2": ObjectName{
		Name:          "aes(1) aes128-CBC(2)",
		HaveExtension: false,
	},
	"2.5.29.15": ObjectName{
		Name:          "certificateExtension(29) keyUsage(15)",
		HaveExtension: false,
	},
	"1.2.840.113549.1.1.1": ObjectName{
		Name:          "pkcs-1(1) rsaEncryption(1)",
		HaveExtension: false,
	},
	"2.5.29.19": ObjectName{
		Name:          "certificateExtension(29) basicConstraints(19)",
		HaveExtension: false,
	},
	"1.2.840.113549.1.9.3": ObjectName{
		Name:          "pkcs-9(9) contentType(3)",
		HaveExtension: false,
	},
	"2.16.840.1.101.3.4.2.1": ObjectName{
		Name:          "hashAlgs(2) sha256(1)",
		HaveExtension: false,
	},
	"2.5.4.8": ObjectName{
		Name:          "attributeType(4) stateOrProvinceName(8)",
		HaveExtension: false,
	},
	"2.5.29.37": ObjectName{
		Name:          "certificateExtension(29) extKeyUsage(37)",
		HaveExtension: false,
	},
	"2.5.4.6": ObjectName{
		Name:          "attributeType(4) countryName(6)",
		HaveExtension: false,
	},
	"2.5.4.10": ObjectName{
		Name:          "attributeType(4) organizationName(10)",
		HaveExtension: false,
	},
	"1.2.840.113549.3.7": ObjectName{
		Name:          "encryptionalgorithm(3) des-ede3-cbc(7)",
		HaveExtension: false,
	},
	"2.5.29.31": ObjectName{
		Name:          "certificateExtension(29) cRLDistributionPoints(31)",
		HaveExtension: false,
	},
	"1.3.6.1.5.5.7.1.1": ObjectName{
		Name:          "pe(1) authorityInfoAccess(1)",
		HaveExtension: false,
	},
	"1.2.840.113549.1.9.4": ObjectName{
		Name:          "pkcs-9(9) messageDigest(4)",
		HaveExtension: false,
	},
	"1.2.840.113549.1.9.15": ObjectName{
		Name:          "pkcs-9(9) smimeCapabilities(15)",
		HaveExtension: false,
	},
}