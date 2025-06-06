// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

//go:build ca_from_tp
// +build ca_from_tp

package main

import (
	"crypto/x509"
	"flag"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/elastic/go-elasticsearch/v9"
)

func main() {
	log.SetFlags(0)

	var (
		err error

		// --> Configure the path to the certificate authority and the password
		//
		cacert   = flag.String("cacert", "certificates/ca/ca.crt", "Path to the file with certificate authority")
		password = flag.String("password", "elastic", "Elasticsearch password")
	)
	flag.Parse()

	ca, err := ioutil.ReadFile(*cacert)
	if err != nil {
		log.Fatalf("ERROR: Unable to read CA from %q: %s", *cacert, err)
	}

	// --> Clone the default HTTP transport
	//
	tp := http.DefaultTransport.(*http.Transport).Clone()

	// --> Initialize the set of root certificate authorities
	//
	if tp.TLSClientConfig.RootCAs, err = x509.SystemCertPool(); err != nil {
		log.Fatalf("ERROR: Problem adding system CA: %s", err)
	}

	// --> Add the custom certificate authority
	//
	if ok := tp.TLSClientConfig.RootCAs.AppendCertsFromPEM(ca); !ok {
		log.Fatalf("ERROR: Problem adding CA from file %q", *cacert)
	}

	es, err := elasticsearch.NewClient(
		elasticsearch.Config{
			Addresses: []string{"https://localhost:9200"},
			Username:  "elastic",
			Password:  *password,

			// --> Pass the transport to the client
			//
			Transport: tp,
		},
	)
	if err != nil {
		log.Fatalf("ERROR: Unable to create client: %s", err)
	}

	res, err := es.Info()
	if err != nil {
		log.Fatalf("ERROR: Unable to get response: %s", err)
	}

	log.Println(res)
}
