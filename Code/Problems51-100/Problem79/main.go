package main

import (
	"fmt"
	"os"
	"strings"
)

type requirements struct {
	predecessor []bool
	exists      bool
}

func solve(req []requirements) int {
	sol := 0
	numsLeft := 0

	for i := range req {
		if req[i].exists {
			numsLeft++
		}
	}

	for numsLeft > 0 {
		leastReqs := 99
		var leastInd int
		for i, r := range req {
			if r.exists {
				sum := 0
				for _, p := range r.predecessor {
					if p {
						sum++
					}
				}
				if sum < leastReqs {
					leastReqs = sum
					leastInd = i
				}
			}
		}
		req[leastInd].exists = false
		for i := range req {
			req[i].predecessor[leastInd] = false
		}
		numsLeft--
		sol *= 10
		sol += leastInd
	}
	return sol
}

func main() {
	file, err := os.Open("p079_keylog.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	buff := make([]byte, 10000)
	file.Read(buff)
	attempts := strings.Split(string(buff), "\n")
	req := make([]requirements, 10)
	for i := range req {
		req[i].predecessor = make([]bool, 10)
	}
	for _, at := range attempts[:len(attempts)-1] {
		req[int(at[2])-48].predecessor[int(at[1])-48] = true
		req[int(at[2])-48].predecessor[int(at[0])-48] = true
		req[int(at[1])-48].predecessor[int(at[0])-48] = true
		req[int(at[0])-48].exists = true
		req[int(at[1])-48].exists = true
		req[int(at[2])-48].exists = true
	}
	fmt.Println(solve(req))

}
