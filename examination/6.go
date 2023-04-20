package examination

import (
	"encoding/json"
	"fmt"
	"time"
)

func main6() {
	t := struct {
		time.Time
		N int
	}{
		time.Date(2021, 3, 30, 0, 0, 0, 0, time.UTC),
		5,
	}

	m, _ := json.Marshal(t)
	fmt.Printf("%s", m)
}

/**
A. {"Time": "2021-03-30T00:00:00Z", "N": 5}
B. "2021-03-30T00:00:00Z"
C. {"N": 5}
D. <nil>
E. panic
**/
