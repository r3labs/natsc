/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/nats-io/nats"
)

var n *nats.Conn
var url string
var debug bool
var retries int
var timeout uint
var maxreplies int

func exit(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	os.Exit(0)
}

func options() (string, string, string) {
	mode := os.Args[1]

	if len(os.Args) > 2 && mode != "--help" {
		os.Args = append(os.Args[:1], os.Args[1+1:]...)
	}

	flag.StringVar(&url, "s", nats.DefaultURL, "nats url")
	flag.BoolVar(&debug, "v", false, "verbose")
	flag.IntVar(&retries, "r", 1, "retries")
	flag.UintVar(&timeout, "t", 1, "timeout")
	flag.IntVar(&maxreplies, "n", 0, "maximum replies")

	flag.Parse()

	return mode, flag.Arg(0), flag.Arg(1)
}

func connect() {
	var err error
	n, err = nats.Connect(url)
	if err != nil {
		exit(err)
	}
}

func main() {
	var err error
	mode, subject, data := options()

	connect()

	switch mode {
	case "pub", "publish":
		err = publish(subject, data)
	case "sub", "subscribe":
		err = subscribe(subject)
	case "req", "request":
		err = request(subject, data)
	case "rep", "reply":
		//err = reply(subject, data)
	case "pubsub":
		err = pubsub(subject, data)
	default:
		flag.Usage()
	}

	exit(err)

}
