// Copyright 2013 tsuru authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package testing

import (
	"bytes"
	"github.com/globocom/tsuru/exec"
	"launchpad.net/gocheck"
	"testing"
)

type S struct{}

var _ = gocheck.Suite(&S{})

func Test(t *testing.T) { gocheck.TestingT(t) }

func (s *S) TestFakeExecutorImplementsExecutor(c *gocheck.C) {
	var _ exec.Executor = &FakeExecutor{}
}

func (s *S) TestExecute(c *gocheck.C) {
	var e FakeExecutor
	var b bytes.Buffer
	cmd := "ls"
	args := []string{"-lsa"}
	err := e.Execute(cmd, args, nil, &b, &b)
	c.Assert(err, gocheck.IsNil)
	cmd = "ps"
	args = []string{"aux"}
	err = e.Execute(cmd, args, nil, &b, &b)
	c.Assert(err, gocheck.IsNil)
	cmd = "ps"
	args = []string{"-ef"}
	err = e.Execute(cmd, args, nil, &b, &b)
	c.Assert(err, gocheck.IsNil)
	c.Assert(e.ExecutedCmd("ls", []string{"-lsa"}), gocheck.Equals, true)
	c.Assert(e.ExecutedCmd("ps", []string{"aux"}), gocheck.Equals, true)
	c.Assert(e.ExecutedCmd("ps", []string{"-ef"}), gocheck.Equals, true)
}

func (s *S) TestFakeExecutorOutput(c *gocheck.C) {
	e := FakeExecutor{output: []byte("ble")}
	var b bytes.Buffer
	cmd := "ls"
	args := []string{"-lsa"}
	err := e.Execute(cmd, args, nil, &b, &b)
	c.Assert(err, gocheck.IsNil)
	c.Assert(e.ExecutedCmd("ls", []string{"-lsa"}), gocheck.Equals, true)
	c.Assert(b.String(), gocheck.Equals, "ble")
}

func (s *S) TestErrorExecutorImplementsExecutor(c *gocheck.C) {
	var _ exec.Executor = &ErrorExecutor{}
}

func (s *S) TestErrorExecute(c *gocheck.C) {
	var e ErrorExecutor
	var b bytes.Buffer
	cmd := "ls"
	args := []string{"-lsa"}
	err := e.Execute(cmd, args, nil, &b, &b)
	c.Assert(err, gocheck.NotNil)
	cmd = "ps"
	args = []string{"aux"}
	err = e.Execute(cmd, args, nil, &b, &b)
	c.Assert(err, gocheck.NotNil)
	cmd = "ps"
	args = []string{"-ef"}
	err = e.Execute(cmd, args, nil, &b, &b)
	c.Assert(err, gocheck.NotNil)
	c.Assert(e.ExecutedCmd("ls", []string{"-lsa"}), gocheck.Equals, true)
	c.Assert(e.ExecutedCmd("ps", []string{"aux"}), gocheck.Equals, true)
	c.Assert(e.ExecutedCmd("ps", []string{"-ef"}), gocheck.Equals, true)
}

func (s *S) TestErrorExecutorOutput(c *gocheck.C) {
	e := ErrorExecutor{output: []byte("ble")}
	var b bytes.Buffer
	cmd := "ls"
	args := []string{"-lsa"}
	err := e.Execute(cmd, args, nil, &b, &b)
	c.Assert(err, gocheck.NotNil)
	c.Assert(e.ExecutedCmd("ls", []string{"-lsa"}), gocheck.Equals, true)
	c.Assert(b.String(), gocheck.Equals, "ble")
}

func (s *S) TestGetCommands(c *gocheck.C) {
	e := FakeExecutor{}
	b := &bytes.Buffer{}
	err := e.Execute("sudo", []string{"ifconfig", "-a"}, nil, b, b)
	c.Assert(err, gocheck.IsNil)
	cmds := e.GetCommands("sudo")
	expected := []command{{name: "sudo", args: []string{"ifconfig", "-a"}}}
	c.Assert(cmds, gocheck.DeepEquals, expected)
}
