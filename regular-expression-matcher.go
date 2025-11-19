package main
import ("fmt"
	"strings")

/* this is a regular expresion matcher for google emails (.*@gmail\.com)
   .* meaning anything
   \. meaning purely .
*/
func ExpressionMatcher (pt string, ex string) bool {

var result bool

if strings.Count(ex, "@") == 1 {

	p := []rune(pt)
        str := []rune(ex)

        rxm := make([][]bool, len(str)+1)

        for i := range rxm{
                rxm[i] = make([]bool, len(p)+1)
        }

        rxm[0][0] = true

        for j := 1; j <= len(p); j++ {
                if p[j-1] == '*' && j >= 2{
                        rxm[0][j] = rxm[0][j-2]
                }
        }

        for i := 1; i <= len(str); i++ {
                for j := 1; j <= len(p); j++ {

                        // case for only . (\.)
			if j >= 2 && p[j-1] == '.' && p[j-2] == '\\' {
                		if str[i-1] == '.' {
                    			rxm[i][j] = rxm[i-1][j-2]
                		}
                		continue
            		}

           		 //case for . and identical characters
            		if p[j-1] == '.' || p[j-1] == str[i-1] {
                		rxm[i][j] = rxm[i-1][j-1]
                	continue
            		}

            		// case for zero or more occurences (*)
            		if p[j-1] == '*' && j >= 2 {

                		rxm[i][j] = rxm[i][j-2]

                		if p[j-2] == '.' || p[j-2] == str[i-1] {
                    			rxm[i][j] = rxm[i][j] || rxm[i-1][j]
                		}
                		continue
            		}



                rxm[i][j] = false
                }
        }

	result = rxm[len(str)][len(p)]
}else {
                fmt.Println("Invalid Expression")
        }

return result
}

func main(){

	pattern := ".*@gmail\\.com"
	var expression string

	fmt.Println("pattern is: ", string(pattern))

	// Read string input for the string to match
	fmt.Printf("Enter string: \n")
	fmt.Scanln(&expression)


	if ExpressionMatcher(pattern,expression) {
		fmt.Println("pattern " + string(pattern) + " matches expression " + string(expression))
	}else {
		fmt.Println("pattern " + string(pattern) + " does not matches expression " + string(expression))
	}

}
