// Copyright 2022 Innkeeper auroralzdf auroralzdf@gmail.com. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/auroralzdf/apis.

package main

import (
	"os"

	"apis/internal/apis"
)

func main() {
	command := apis.Command()
	if err := command.Execute(); err != nil {
		os.Exit(1)
	}
}
