/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package main

import (
	"fmt"
	"time"

	"github.com/nats-io/nats"
)

func request(subj string, data []byte) error {
	var err error
	var msg *nats.Msg
	t := time.Second * time.Duration(timeout)

	if debug {
		fmt.Printf("Requesting: %s\n", subj)
	}

	for i := 0; i < retries; i++ {
		msg, err = n.Request(subj, data, t)
		if err == nil {
			fmt.Println(string(msg.Data))
			break
		}
		if debug {
			fmt.Println("Request timed out. Retrying...")
		}
	}

	return err
}
