import (
	"fmt"
)
func Test(){
	m := make(map[string]interface{})
	m["name"] = "好家伙"
	m["age"] = 18
	m["isMale"] = true
	fmt.Println(m)
}