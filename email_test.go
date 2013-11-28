// Copyright 2013 Dmitry Chestnykh. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package validator

import (
	"testing"
)

// Validation tests.

var validEmails = []string{
	"a@example.com",
	"postmaster@example.com",
	"president@kremlin.gov.ru",
	"example@example.co.uk",
	"test@example.com.",
}

var invalidEmails = []string{
	"",
	"example",
	"example.com",
	".com",
	"адрес@пример.рф",
}

func TestIsValidEmail(t *testing.T) {
	for i, v := range validEmails {
		if !IsValidEmail(v) {
			t.Errorf("%d: didn't accept valid email: %s", i, v)
		}
	}
	for i, v := range invalidEmails {
		if IsValidEmail(v) {
			t.Errorf("%d: accepted invalid email: %s", i, v)
		}
	}
}

func TestValidateEmailByResolvingDomain(t *testing.T) {
	err := ValidateEmailByResolvingDomain("abuse@gmail.com")
	if err != nil {
		t.Errorf("%s", err)
	}
	err = ValidateEmailByResolvingDomain("nomx@example.com")
	if err == nil {
		t.Errorf("invalid email address validated")
	}
}

// Normalization tests.

var sameEmails = []string{
	"test@example.com",
	"test@example.com.",
	"test@EXAMPLE.COM.",
	"test@EXAMPLE.COM",
	"test@ExAmpLE.com",
	" test@example.com \n",
}

var differentEmails = []string{
	"test@example.com",
	"TEST@example.com",
	"president@whitehouse.gov",
}

func TestNormalizeEmail(t *testing.T) {
	for i, v0 := range sameEmails {
		for j, v1 := range sameEmails {
			if i == j {
				continue
			}
			nv0 := NormalizeEmail(v0)
			nv1 := NormalizeEmail(v1)
			if nv0 == "" {
				t.Errorf("%d: email invalid: %q", i, nv0)
			}
			if nv0 != nv1 {
				t.Errorf("%d-%d: normalized emails differ: %q and %q", i, j, nv0, nv1)
			}
		}
	}
	for i, v0 := range differentEmails {
		for j, v1 := range differentEmails {
			if i == j {
				continue
			}
			nv0 := NormalizeEmail(v0)
			nv1 := NormalizeEmail(v1)
			if nv0 == "" {
				t.Errorf("%d: email invalid: %q", i, nv0)
			}
			if nv0 == nv1 {
				t.Errorf("%d-%d: normalized emails are the same: %q and %q", i, j, nv0, nv1)
			}
		}
	}
}
