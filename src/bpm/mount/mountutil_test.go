// Copyright (C) 2018-Present CloudFoundry.org Foundation, Inc. All rights reserved.
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

package mount

import (
	"io/ioutil"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Mount", func() {
	Describe("Mounts", func() {
		It("parses the contents of /proc/mounts", func() {
			mnts, err := Mounts()
			Expect(err).NotTo(HaveOccurred())
			Expect(mnts).ToNot(BeEmpty())
		})
	})

	Describe("IsMountpoint", func() {
		It("returns whether or not the given path has anything mounted on it", func() {
			mounts := []Mnt{
				{MountPoint: "/"},
				{MountPoint: "/tmp"},
				{MountPoint: "/mnt/cdrom"},
			}

			Expect(isMountpoint(mounts, "/")).To(BeTrue())
			Expect(isMountpoint(mounts, "/tmp")).To(BeTrue())
			Expect(isMountpoint(mounts, "/mnt/cdrom")).To(BeTrue())

			Expect(isMountpoint(mounts, "/home/user")).To(BeFalse())
			Expect(isMountpoint(mounts, "")).To(BeFalse())
		})
	})

	Describe("parseMountFile", func() {
		It("returns a slice of mounts", func() {
			mnts, err := ParseFstab(fileContents("testdata/mount"))
			Expect(err).NotTo(HaveOccurred())
			Expect(mnts).To(ConsistOf(
				Mnt{
					Device:     "proc",
					MountPoint: "/proc",
					Filesystem: "proc",
					Options:    []string{"rw", "nosuid", "nodev", "noexec", "relatime"},
				},
				Mnt{
					Device:     "tmpfs",
					MountPoint: "/dev",
					Filesystem: "tmpfs",
					Options:    []string{"rw", "nosuid", "size=65536k", "mode=755"},
				},
				Mnt{
					Device:     "devpts",
					MountPoint: "/dev/console",
					Filesystem: "devpts",
					Options:    []string{"rw", "nosuid", "noexec", "relatime", "gid=5", "mode=620", "ptmxmode=666"},
				},
			))
		})

		Context("when the file contains an invalid mount format", func() {
			It("returns an error", func() {
				_, err := ParseFstab(fileContents("testdata/invalid-mount"))
				Expect(err).To(HaveOccurred())
			})
		})
	})
})

func fileContents(path string) []byte {
	bs, err := ioutil.ReadFile(path)
	Expect(err).ToNot(HaveOccurred())
	return bs
}
