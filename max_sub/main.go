package main

import "fmt"

func main() {
	msg := "123451234"

	start_pos := 0
	max_len := 0
	//var note_last_pos map[byte]int
	note_last_pos := make(map[byte]int)

	for pos, i := range []byte(msg) {
		fmt.Printf("%d , %c\n", pos, i)
		last_pos, OK := note_last_pos[i]
		if OK {
			if start_pos <= last_pos {
				cur_len := pos - start_pos
				if cur_len > max_len {
					max_len = cur_len
				}
				start_pos = last_pos + 1
			}
		}
		note_last_pos[i] = pos
	}

	fmt.Println(max_len)


	var nameinfo map[string]string
	nameinfo = map[string]string{}
	nameinfo["a"] = "b"
	fmt.Println(nameinfo)
}