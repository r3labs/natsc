/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package main

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/nats-io/nats"
	"github.com/r3labs/pattern"
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

	_, err := n.Subscribe(">", func(msg *nats.Msg) {
		if !pattern.Match(msg.Subject, subj) {
			return
		}

		if withreplies && isResponse(msg.Subject) {
			fmt.Println(color.MagentaString("reply: "+msg.Subject) + "\n " + string(msg.Data) + "\n")
			return
		}

		fmt.Println(color.CyanString(msg.Subject) + ": " + msg.Reply + "\n " + string(msg.Data) + "\n")
		if msg.Reply != "" {
			responses = append(responses, msg.Reply)
		}

		if maxreplies == 1 {
			done <- true
		}
		maxreplies--
	})

	// Wait for all to be done
	<-done

	return err
}
