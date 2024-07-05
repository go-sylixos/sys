// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build aix || darwin || dragonfly || freebsd || linux || netbsd || openbsd || solaris || sylixos

package unix_test

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"testing"

	"golang.org/x/sys/unix"
)

func TestMmap(t *testing.T) {
	mmapProt := unix.PROT_NONE
	mprotectProt := unix.PROT_READ | unix.PROT_WRITE
	// On NetBSD PAX mprotect prohibits setting protection bits
	// missing from the original mmap call unless explicitly
	// requested with PROT_MPROTECT.
	if runtime.GOOS == "netbsd" {
		// PROT_MPROTECT(x) is defined as ((x) << 3):
		// https://github.com/NetBSD/src/blob/aba449a55bf91b44bc68f542edd9afa341962b89/sys/sys/mman.h#L73
		mmapProt = mprotectProt << 3
	}
	b, err := unix.Mmap(-1, 0, unix.Getpagesize(), mmapProt, unix.MAP_ANON|unix.MAP_PRIVATE)
	if err != nil {
		t.Fatalf("Mmap: %v", err)
	}
	if err := unix.Mprotect(b, mprotectProt); err != nil {
		t.Fatalf("Mprotect: %v", err)
	}

	b[0] = 42

	if runtime.GOOS == "aix" {
		t.Skip("msync returns invalid argument for AIX, skipping msync test")
	} else {
		if err := unix.Msync(b, unix.MS_SYNC); err != nil {
			t.Fatalf("Msync: %v", err)
		}
	}

	if err := unix.Madvise(b, unix.MADV_DONTNEED); err != nil {
		t.Fatalf("Madvise: %v", err)
	}
	if err := unix.Munmap(b); err != nil {
		t.Fatalf("Munmap: %v", err)
	}
}

func TestMmapOnFile(t *testing.T) {
	tmpdir := filepath.Join(t.TempDir(), "testdata")
	if err := os.Mkdir(tmpdir, 0700); err != nil {
		t.Fatal(err)
	}
	filename := filepath.Join(tmpdir, "memmapped_file")
	destination, err := os.Create(filename)
	if err != nil {
		t.Fatal("os.Create:", err)
		return
	}

	fmt.Println(filename)

	fmt.Fprintf(destination, "%s\n", "0 <- Flipped between 0 and 1 when test runs successfully")
	fmt.Fprintf(destination, "%s\n", "//Do not change contents - mmap test relies on this")
	destination.Close()
	fd, err := unix.Open(filename, unix.O_RDWR, 0777)
	if err != nil {
		t.Fatalf("Open: %v", err)
	}

	b, err := unix.Mmap(fd, 0, 8, unix.PROT_READ, unix.MAP_SHARED)
	if err != nil {
		t.Fatalf("Mmap: %v", err)
	}

	if err := unix.Mprotect(b, unix.PROT_READ|unix.PROT_WRITE); err != nil {
		t.Fatalf("Mprotect: %v", err)
	}

	// Flip flag in test file via mapped memory
	flagWasZero := true
	if b[0] == '0' {
		b[0] = '1'
	} else if b[0] == '1' {
		b[0] = '0'
		flagWasZero = false
	}

	if err := unix.Msync(b, unix.MS_SYNC); err != nil {
		t.Fatalf("Msync: %v", err)
	}

	// Read file from FS to ensure flag flipped after msync
	buf, err := os.ReadFile(filename)
	if err != nil {
		t.Fatalf("Could not read mmapped file from disc for test: %v", err)
	}
	if flagWasZero && buf[0] != '1' || !flagWasZero && buf[0] != '0' {
		t.Error("Flag flip in MAP_SHARED mmapped file not visible")
	}

	if err := unix.Munmap(b); err != nil {
		t.Fatalf("Munmap: %v", err)
	}
}

func TestMmapPtr(t *testing.T) {
	p, err := unix.MmapPtr(-1, 0, nil, uintptr(2*unix.Getpagesize()),
		unix.PROT_NONE, unix.MAP_ANON|unix.MAP_PRIVATE)
	if err != nil {
		t.Fatalf("MmapPtr: %v", err)
	}

	if _, err := unix.MmapPtr(-1, 0, p, uintptr(unix.Getpagesize()),
		unix.PROT_READ|unix.PROT_WRITE, unix.MAP_ANON|unix.MAP_PRIVATE|unix.MAP_FIXED); err != nil {
		t.Fatalf("MmapPtr: %v", err)
	}

	*(*byte)(p) = 42

	if err := unix.MunmapPtr(p, uintptr(2*unix.Getpagesize())); err != nil {
		t.Fatalf("MunmapPtr: %v", err)
	}
}
