/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package main

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/nats-io/nats"
)

var responses []string

func isResponse(id string) bool {
	for _, r := range responses {
		if r == id {
			return true
		}
	}
	return false
}

func subscribe(subj string) error {
	done := make(chan bool)

	if debug {
		fmt.Printf("Subscribing to: %s\n", subj)
	}

	_, err := n.Subscribe(subj, func(msg *nats.Msg) {
		fmt.Println(color.CyanString(msg.Subject) + ": " + msg.Reply + "\n " + string(msg.Data) + "\n")
		if msg.Reply != "" {
			responses = append(responses, msg.Reply)
		}

		if maxreplies == 1 {
			done <- true
		}
		maxreplies--
	})

	if err != nil {
		return err
	}

	if withreplies {
		_, err = n.Subscribe(">", func(msg *nats.Msg) {
			if isResponse(msg.Subject) {
				fmt.Println(color.MagentaString("reply: "+msg.Subject) + "\n " + string(msg.Data) + "\n")
			}
		})
	}

	// Wait for all to be done
	<-done

	return err
}
