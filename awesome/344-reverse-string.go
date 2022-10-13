package awesome


func reverseString(s []byte)  {
	length := 0
	for ;; {

		if length >= len(s) / 2 {
			break
		}
		temp := s[0 + length]
		s[0 + length] = s[len(s) - 1 - length]
		s[len(s) - 1 - length] = temp

		length ++
	}
}