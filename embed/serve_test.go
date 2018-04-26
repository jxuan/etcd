// Copyright 2017 The etcd Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package embed

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/coreos/etcd/auth"
	"golang.org/x/crypto/bcrypt"
)

// TestStartEtcdWrongToken ensures that StartEtcd with wrong configs returns with error.
func TestStartEtcdWrongToken(t *testing.T) {
	tdir, err := ioutil.TempDir(os.TempDir(), "token-test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tdir)
	cfg := NewConfig()
	cfg.Dir = tdir
	cfg.AuthToken = "wrong-token"
	if _, err = StartEtcd(cfg); err != auth.ErrInvalidAuthOpts {
		t.Fatalf("expected %v, got %v", auth.ErrInvalidAuthOpts, err)
	}
}

// TestStartEtcdLargeBcryptCost ensures that StartEtcd with invalid large bcrypt-cost returns with error.
func TestStartEtcdLargeBcryptCost(t *testing.T) {
	tdir, err := ioutil.TempDir(os.TempDir(), "large-bcrypt-cost-test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tdir)
	cfg := NewConfig()
	cfg.Dir = tdir
	cfg.BcryptCost = uint(bcrypt.MaxCost) + 1 // Greater than bcrypt.MaxCost
	if _, err = StartEtcd(cfg); err != auth.ErrInvalidAuthOpts {
		t.Fatalf("expected %v, got %v", auth.ErrInvalidAuthOpts, err)
	}
}

// TestStartEtcdSmallBcryptCost ensures that StartEtcd with invalid small bcrypt-cost returns with error.
func TestStartEtcdSmallBcryptCost(t *testing.T) {
	tdir, err := ioutil.TempDir(os.TempDir(), "small-bcrypt-cost-test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tdir)
	cfg := NewConfig()
	cfg.Dir = tdir
	cfg.BcryptCost = uint(bcrypt.MinCost) - 1 // Smaller than bcrypt.MinCost
	if _, err = StartEtcd(cfg); err != auth.ErrInvalidAuthOpts {
		t.Fatalf("expected %v, got %v", auth.ErrInvalidAuthOpts, err)
	}
}
