package main

import (
    "testing"
)

var ErrInvalidUTF8 = errors.New("invalid utf8")
func TestGetUTFLength(t *testing.T) {
    tests := []struct {
        input    []byte
        expected int
        err      error
    }{
        {[]byte("Hello, world!"), 13, nil}, 
        {[]byte("Привет, мир!"), 12, nil},   
        {[]byte("こんにちは"), 5, nil},     
        {[]byte{0xff, 0xfe, 0xfd}, 0, ErrInvalidUTF8},
        {[]byte{0xe0, 0xa0, 0xaf}, 1, nil}, 
        {[]byte{0xff}, 0, ErrInvalidUTF8},  
	}

    for _, test := range tests {
        result, err := GetUTFLength(test.input)
        if result != test.expected || (err != nil && err.Error() != test.err.Error()) {
            t.Errorf("GetUTFLength(%q) = %d, %v; ожидается %d, %v", test.input, result, err, test.expected, test.err)
        }
    }
}