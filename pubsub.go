/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package main

import (
	"bufio"
	"os"
	"strings"
)

func pubsub(subj, data string) error {
	go subscribe(subj)
	reader := bufio.NewReader(os.Stdin)
	for {
		data, _ := reader.ReadString('\n')
		publish(subj, strings.Trim(data, "\n"))
	}
}
