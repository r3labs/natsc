/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package main

import "fmt"

func publish(subj, data string) error {
	if debug {
		fmt.Printf("Publishing to: %s\n", subj)
	}
	err := n.Publish(subj, []byte(data))
	if err != nil {
		return err
	}
	return n.Flush()
}
