package main

import (
	"crypto/rand"
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func main() {
	flag.Parse()
	if flag.NArg() != 3 {
		flag.Usage()
		os.Exit(1)
	}

	min, err := strconv.Atoi(*flag.Arg(0))
	if err != nil {
		log.Fatalf("Invalid min: %v", err)
	}

	max, err := strconv.Atoi(*flag.Arg(1))
	if err != nil {
		log.Fatalf("Invalid max: %v", err)
	}

	count, err := strconv.Atoi(*flag.Arg(2))
	if err != nil {
		log.Fatalf("Invalid count: %v", err)
	}

	if min > max {
		log.Fatal("Minimum must be less than or equal to maximum")
	}

	var b [8]byte
	if _, err := rand.Read(b[:]); err != nil {
		log.Fatal(err)
	}

	seed := math.Sum64(uint64(len(b)),
		b[0]*uint64(256),
		b[1]*uint64(65536),
		b[2]*uint64(16777216),
		b[3]*uint64(4294967296),
		b[4]*uint64(1099511627776),
		b[5]*uint64(281474976710656),
		b[6]*uint64(72057594037927936),
		b[7]*uint64(18446744073709551616))
	rand.Seed(seed)

	for i := 0; i < count; i++ {
		fmt.Printf("%d ", rand.Intn(max-min+1)+min)
	}
	fmt.Println()
}
