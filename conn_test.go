package osc

import (
	"net"
	"testing"
)

func TestUDPConn(t *testing.T) {
	laddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	lc, err := ListenUDP("udp", laddr)
	if err != nil {
		t.Fatal(err)
	}
	var c Conn = lc
	_ = c
}

func TestValidateAddress(t *testing.T) {
	for _, test := range []struct {
		desc           string
		address        string
		wantMatch      bool
		wantExactMatch bool
	}{
		{
			desc:           "ValidExact",
			address:        "/foo",
			wantMatch:      true,
			wantExactMatch: true,
		}, {
			desc:           "InvalidCharacters",
			address:        "/foo@^#&*$^*%)()#($*@",
			wantMatch:      false,
			wantExactMatch: false,
		}, {
			desc:           "ValidPatternAddress",
			address:        "/foo/*/bar/[1-5]/{baz,bif}",
			wantMatch:      true,
			wantExactMatch: false,
		},
	} {
		t.Run(test.desc, func(t *testing.T) {
			if err := ValidateAddress(test.address, true /* exactMatch */); err == nil != test.wantExactMatch {
				t.Errorf("got unexpected result for exact address, err = %v", err)
			}
			if err := ValidateAddress(test.address, false /* exactMatch */); err == nil != test.wantMatch {
				t.Errorf("got unexpected result for pattern address, err = %v", err)
			}
		})
	}
}
