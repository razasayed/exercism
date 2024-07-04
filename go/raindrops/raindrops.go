package raindrops
import "strconv"

func Convert(number int) string {
	result := ""
    if (!(number%3 == 0) && !(number%5 == 0) && !(number%7 == 0)) {
        return strconv.Itoa(number)
    }
    if number%3 == 0 {
        result += "Pling"
    }
	if number%5 == 0 {
        result+="Plang"
    }
	if number%7 == 0 {
        result+="Plong"
    }
	return result
}
