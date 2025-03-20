package forLoop

const repeatCount = 5

func Repeat(word string) string {
	repeated := ""
	for i := 0; i < repeatCount; i++ {
		repeated += word
	}
	return repeated
}
