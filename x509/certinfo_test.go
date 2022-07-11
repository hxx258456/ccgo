package x509_test

import (
	"fmt"
	"testing"

	"github.com/hxx258456/ccgo/x509"
)

func TestCertinfo(t *testing.T) {
	certPath := "testdata/sm2_sign_cert.cer"

	cert, _ := x509.ReadCertificateFromPemFile(certPath)

	certText, err := x509.CertificateText(cert)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(certText)
}
