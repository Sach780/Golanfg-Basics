package main
import ("encoding/json"
	"fmt"
)
type SensorReading struct {
	Name 		string 	`json:"Name"`
	Capacity 	int16 	`json:"Capacity "`
	Time 		string 	`json:"Time"`
	Information	Info	`json:"Info"`
}

type Info struct {
	Descrition string  `json:"desc "`
}

func main() {
	jsonString := SensorReading{"Name": "Sachin Mishra", "Capacity": 25, "Time":  "2019-01-21T19:07:28Z", "info": {
		"desc": "a sensor reading"
	}
}
var reading SensorReading
err := json.Unmarshal([]byte(jsonString), &reading)
if err != nil {
	fmt.Printf(err)
}
fmt.Printf("%+v\n", reading)

}