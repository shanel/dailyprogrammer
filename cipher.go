package main

import (
	"bytes"
    "fmt"
)

func Encode(text,key string, mapper map[string]int) string {
	rev := []string{"A","B","C","D","E","F","G","H","I","J","K","L","M","N","O","P","Q","R","S","T","U","V","W","X","Y","Z"}
	var key_buffer bytes.Buffer
	for i := 0; i < len(text); i++ {
		key_buffer.WriteString(string(key[i % len(key)]))
	}
	var out_buffer bytes.Buffer
	for i := 0; i < len(text); i++ {
		text_key := text[i]
		b_s := key_buffer.String()
		key_key := b_s[i]
		total := (mapper[string(text_key)] + mapper[string(key_key)]) % 26
		out_buffer.WriteString(rev[total])
	}
	return out_buffer.String()
}

func Decode(text,key string, mapper map[string]int) string {
	rev := []string{"A","B","C","D","E","F","G","H","I","J","K","L","M","N","O","P","Q","R","S","T","U","V","W","X","Y","Z"}
	var key_buffer bytes.Buffer
	for i := 0; i < len(text); i++ {
		key_buffer.WriteString(string(key[i % len(key)]))
	}
	var out_buffer bytes.Buffer
	for i := 0; i < len(text); i++ {
		text_key := text[i]
		b_s := key_buffer.String()
		key_key := b_s[i]
		total := (mapper[string(text_key)] - mapper[string(key_key)] + 26) % 26
		out_buffer.WriteString(rev[total])
	}
	return out_buffer.String()
}


func main() {
	the_map := map[string]int{
		"A": 0,
		"B": 1,
		"C": 2,
		"D": 3,
		"E": 4,
		"F": 5,
		"G": 6,
		"H": 7,
		"I": 8,
		"J": 9,
		"K": 10,
		"L": 11,
		"M": 12,
		"N": 13,
		"O": 14,
		"P": 15,
		"Q": 16,
		"R": 17,
		"S": 18,
		"T": 19,
		"U": 20,
		"V": 21,
		"W": 22,
		"X": 23,
		"Y": 24,
		"Z": 25}
	fmt.Println(Encode("THECAKEISALIE", "GLADOS", the_map))
	fmt.Println(Decode("ZSEFOCKTSDZAK", "GLADOS", the_map))
}
