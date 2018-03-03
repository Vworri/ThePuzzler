package puzzler

import (
	"testing"
)

func Test_md5Hash(t *testing.T) {
	t.Parallel()
	tst := Md5Hash("")
	if tst != "d41d8cd98f00b204e9800998ecf8427e" {

		t.Fail()
	}

}

func Test_Sha1Hash(t *testing.T) {
	t.Parallel()
	tst := Sha1Hash("")
	if tst == "a7ffc6f8bf1ed76651c14756a061d662f580ff4de43b49fa82d80a4b80f8434a" {
		t.Fail()
	}

}

func Test_Sha256Hash(t *testing.T) {
	t.Parallel()
	tst := Sha256Hash("")
	if tst != "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855" {
		t.Fail()
	}

}

func Test_Sha3256(t *testing.T) {
	//t.Parallel()
	tst := Sha3256("")
	if tst != "a7ffc6f8bf1ed76651c14756a061d662f580ff4de43b49fa82d80a4b80f8434a" {
		t.Fail()
	}

}

func Test_Sha512Hash(t *testing.T) {
	//t.Parallel()
	tst := Sha512Hash("")
	if tst != "cf83e1357eefb8bdf1542850d66d8007d620e4050b5715dc83f4a921d36ce9ce47d0d13c5d85f2b0ff8318d2877eec2f63b931bd47417a81a538327af927da3e" {
		t.Fail()
	}

}
