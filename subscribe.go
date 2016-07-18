/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package main

import (
	"fmt"

	"github.com/nats-io/nats"
)

func subscribe(subj string) error {
	done := make(chan bool)

	if debug {
		fmt.Printf("Subscribing to: %s\n", subj)
	}

	_, err := n.Subscribe(subj, func(msg *nats.Msg) {
		fmt.Println(string(msg.Subject) + " " + string(msg.Data))
		if maxreplies == 1 {
			done <- true
		}
		maxreplies--
	})

	// Wait for all to be done
	<-done

	return err
}
