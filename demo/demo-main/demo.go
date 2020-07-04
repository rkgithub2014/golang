package main

import 
	(
	"fmt"
	"os/exec"
	"os"
	"encoding/json"
	)

func main() {
	jsonFunc()
}
func jsonFunc(){
type ColorGroup struct {
	ID int
	Name string
	Colors []string
}
group := ColorGroup {
	ID: 1,
	Name: "Reds",
	Colors: [] string {"Red","Blue","Violet"},
}
b,err :=json.Marshal(group)
if err != nil {
	fmt.Println("error:", err)
}
os.Stdout.Write(b)
}
func psqlFunc() {
	fmt.Println("Hello on Mac!")
	cmd := exec.Command("psql","-h", "localhost","-p", "5432","-d", "testdb","-U", "ramesh","-f", "install.sql")
	stdOutstdErr,err := cmd.CombinedOutput()
	if(err!=nil){
		fmt.Println(err)
	}
	fmt.Printf("%s\n",stdOutstdErr)
}
