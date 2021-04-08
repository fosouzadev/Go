package main

import "fmt"

type IPAddr [4]byte

// TODO: Add a "String() string" method to IPAddr.
func (a IPAddr) String() string {
	var res string

	for i := 0; i < len(a); i++ {
		res += fmt.Sprintf("%v", a[i])

		if i < len(a)-1 {
			res += "."
		}
	}

	// fmt.Sprintf("%v.%v.%v.%v", a[0], a[1], a[2], a[3])

	return res
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}

	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
