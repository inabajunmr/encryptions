package caesar

import "testing"

func TestEncrypt_KeyIs0(t *testing.T) {
	result, err := Encrypt("ABC def", 0)
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}

	if result != "ABC def" {
		t.Fatalf("failed test. result is %v", result)
	}
}

func TestEncrypt_KeyIsNot0(t *testing.T) {
	// key is 1
	result, err := Encrypt("ABC def", 1)
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}
	if result != "BCD efg" {
		t.Fatalf("failed test. result is %v", result)
	}

	// key is 10
	result, err = Encrypt("vwx XYZ", 10)
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}
	if result != "fgh HIJ" {
		t.Fatalf("failed test. result is %v", result)
	}
}

func TestEncrypt_InvalidValue(t *testing.T) {
	// key is 1
	_, err := Encrypt("ABC def 123", 1)
	if err == nil {
		t.Fatalf("failed test")
	}
}

// exception case

func TestDecrypt_KeyIs0(t *testing.T) {
	c, _ := Encrypt("ABC def", 0)
	result, err := Decrypt(c, 0)
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}
	if result != "ABC def" {
		t.Fatalf("failed test. result is %v", result)
	}
}

func TestDecrypt_KeyIsNot0(t *testing.T) {
	// key is 1
	c, _ := Encrypt("ABC def", 1)
	result, err := Decrypt(c, 1)
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}
	if result != "ABC def" {
		t.Fatalf("failed test. result is %v", result)
	}

	// key is 10
	c, _ = Encrypt("vwx XYZ", 10)
	result, err = Decrypt(c, 10)

	if err != nil {
		t.Fatalf("failed test %#v", err)
	}
	if result != "vwx XYZ" {
		t.Fatalf("failed test. result is %v", result)
	}
}

func TestDecrypt_InvalidValue(t *testing.T) {
	// key is 1
	_, err := Decrypt("ABC def 123", 1)
	if err == nil {
		t.Fatalf("failed test")
	}
}
