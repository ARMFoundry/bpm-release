// Copyright (C) 2017-Present Pivotal Software, Inc. All rights reserved.
//
// This program and the accompanying materials are made available under
// the terms of the under the Apache License, Version 2.0 (the "License”);
// you may not use this file except in compliance with the License.
//
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.  See the
// License for the specific language governing permissions and limitations
// under the License.

package usertools

import (
	"errors"
	"os/user"
	"strconv"

	specs "github.com/opencontainers/runtime-spec/specs-go"
)

const VcapUser = "vcap"

//go:generate counterfeiter . UserFinder
type UserFinder interface {
	Lookup(username string) (specs.User, error)
}

type userFinder struct{}

func NewUserFinder() userFinder {
	return userFinder{}
}

func (f userFinder) Lookup(username string) (specs.User, error) {
	u, err := user.Lookup(username)
	if err != nil {
		return specs.User{}, err
	}

	uid, err := strconv.Atoi(u.Uid)
	if err != nil {
		return specs.User{}, err
	}
	if uid < 0 {
		return specs.User{}, errors.New("UID can't be negative")
	}

	gid, err := strconv.Atoi(u.Gid)
	if err != nil {
		return specs.User{}, err
	}
	if gid < 0 {
		return specs.User{}, errors.New("GID can't be negative")
	}

	return specs.User{
		UID:      uint32(uid),
		GID:      uint32(gid),
		Username: u.Username,
	}, nil
}