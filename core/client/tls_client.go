package client

import (
	"crypto/tls"
	"crypto/x509"
	"github.com/netm4ul/netm4ul/cmd/colors"
	"io/ioutil"
	"log"
)

// Build the TLS configuration for server
func TLSBuildClientConf() *tls.Config {

	caCert := "certificates/NetM4ul_CA.crt"
	publKey := "certificates/NetM4ul_Client.crt"
	privKey := "certificates/NetM4ul_Client.pem"

	// Read CA file and initialise
	certBytes, err := ioutil.ReadFile(caCert)
	if err != nil {
		log.Println(colors.Red("Unable to read %s : %s"), caCert, err.Error())
	}

	caCertPool := x509.NewCertPool()
	if ok := caCertPool.AppendCertsFromPEM(certBytes); !ok {
		log.Println(colors.Red("Unable to add CA certificate to certificate pool"))
	}

	// Read own KeyPair
	cert, err := tls.LoadX509KeyPair(publKey, privKey)
	if err != nil {
		log.Println(colors.Red("Unable to read Client X509KeyPair : %s"), err.Error())
	}

	tlsConfig := &tls.Config{
		// Ensure that we only use our "CA" to validate certificates
		RootCAs: caCertPool,
		// Server Key Material
		Certificates: []tls.Certificate{cert},
	}

	tlsConfig.BuildNameToCertificate()

	return tlsConfig
}
