/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package main

import (
	"errors"
	"os/exec"
	"testing"
	"time"

	"github.com/nats-io/nats"
	. "github.com/smartystreets/goconvey/convey"
)

func wait(ch chan bool) error {
	return waitTime(ch, 500*time.Millisecond)
}

func waitTime(ch chan bool, timeout time.Duration) error {
	select {
	case <-ch:
		return nil
	case <-time.After(timeout):
	}
	return errors.New("timeout")
}

var body string

func TestBasicRedirections(t *testing.T) {
	Convey("Scenario: publis", t, func() {
		Convey("Given nats is up and running", func() {
			body = ""
			n, _ := nats.Connect(nats.DefaultURL)
			chtest := make(chan bool)

			n.Subscribe("test.publish", func(msg *nats.Msg) {
				chtest <- true
			})

			Convey("When I call an existing endpoint", func() {
				cmd := exec.Command("natsc", "publish", "test.publish", "something")
				output, err := cmd.CombinedOutput()
				Convey("Then the message should be received", func() {
					So(err, ShouldBeNil)
					So(string(output), ShouldEqual, "")
					ch := wait(chtest)
					So(ch, ShouldBeNil)
				})
			})
		})
	})
}
