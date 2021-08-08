# Problem statement

We have a program that processes 1000 jobs asynchronously. We believe that there is a better way to do it. So we ask to answer the following questions:


```
package main


import (
	"fmt"
	"time"
	"context"
)


func do(ctx context.Context, id int) {
	time.Sleep(100 * time.Millisecond)
	fmt.Printf("%d done\n", id)
}


func main() {
	for id := 0; id < 1000; id++ {
		do(context.Background(), id)
	}
}
```

1. Could you describe the problems that you see in the code snipped below
1. Could you provide us with a better solution? (You can submit your version of code in you answer)
1. Could you describe the pros ands cons of your solution
