/*
Copyright IBM Corp. 2016 All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

		 http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package mocks

import (
	"errors"
	"reflect"

	"github.com/hyperledger/fabric/bccsp"
)

type Encryptor struct {
	KeyArg       bccsp.Key
	PlaintextArg []byte
	OptsArg      bccsp.EncrypterOpts

	EncValue []byte
	EncErr   error
}

func (e *Encryptor) Encrypt(k bccsp.Key, plaintext []byte, opts bccsp.EncrypterOpts) (ciphertext []byte, err error) {
	if !reflect.DeepEqual(e.KeyArg, k) {
		return nil, errors.New("invalid key")
	}
	if !reflect.DeepEqual(e.PlaintextArg, plaintext) {
		return nil, errors.New("invalid plaintext")
	}
	if !reflect.DeepEqual(e.OptsArg, opts) {
		return nil, errors.New("invalid opts")
	}

	return e.EncValue, e.EncErr
}

type Signer struct {
	KeyArg    bccsp.Key
	DigestArg []byte
	OptsArg   bccsp.SignerOpts

	Value []byte
	Err   error
}

func (s *Signer) Sign(k bccsp.Key, digest []byte, opts bccsp.SignerOpts) (signature []byte, err error) {
	if !reflect.DeepEqual(s.KeyArg, k) {
		return nil, errors.New("invalid key")
	}
	if !reflect.DeepEqual(s.DigestArg, digest) {
		return nil, errors.New("invalid digest")
	}
	if !reflect.DeepEqual(s.OptsArg, opts) {
		return nil, errors.New("invalid opts")
	}

	return s.Value, s.Err
}