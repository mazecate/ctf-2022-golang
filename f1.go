package main

import (
	"fmt"
	"math/rand"
	"os"
	_ "encoding/binary"
	_ "math"
)

type tbl struct {
    jx string
	j int
	i uint64
	flagJ uint8
	x uint8
	uintj uint64
}

func printTbl(a tbl) {
	fmt.Println(a.jx, "\t\t", a.j, "\t\t", a.i, "\t\t", a.flagJ, "\t\t", a.x,  "\t\t\t", a.uintj)
	// fmt.Printf("type >>>>> %T\n", string(flag[j] ^ x))
	// fmt.Printf("type >>>>> %T\n", j)
	// fmt.Printf("type >>>>> %T\n", i)
	// fmt.Printf("type >>>>> %T\n", flag[j])
	// fmt.Printf("type >>>>> %T\n", x)
	// fmt.Println(string(flag[j] ^ x), "\t", j, "\t", i, "\t", flag[j], "\t", x)
}

func main() {
	rand.Seed(1337)

	flag, err := os.ReadFile("flag.enc")
	if err != nil {
		fmt.Println("cannot open flag.enc")
		os.Exit(1)
	}

	fmt.Println()
	fmt.Println("flag >>>>>>> ", flag)
	fmt.Printf("type >>>>> %T\n", flag)
	fmt.Println("length >>>>>> ", len(flag))
	fmt.Println()

	
	fmt.Println("FlagJ^X\t\tj\t\ti\t\tflagJ\t\tx\t\t\tleftshit")
	for i, j := uint64(0), 0; j < len(flag); i++ {
		rand.Uint64()
		if i == uint64(1)<<j {
			x := byte(rand.Uint64())
			printTbl(tbl{jx:string(flag[j] ^ x), j: j, i: i, flagJ: flag[j], x: x, uintj: uint64(1)<<j})
			// fmt.Print(string(flag[j] ^ x))
			j += 1
			// hkcert22{y0u_c4n_n4vig4t3_r4nd0mn3s5_f}
		}
	}
	

	// var powX float64 = 2
	// for i, j := uint64(0), 0; j < len(flag); i++ {
	// 	rand.Uint64()
	// 	if i == uint64(1)<<float64(math.Pow(2, j)) {
	// 		fmt.Println("i >>> ", i , "  ", "     ", uint64(1), "    ",uint64(1)<<j)
	// 		// x := byte(rand.Uint64())
	// 		// fmt.Println("powX >>> ", powX)
	// 		// x:=byte(math.Pow(2, powX))
	// 		// printTbl(tbl{jx:string(flag[j] ^ x), j: j, i: i, flagJ: flag[j], x: x})
	// 		// fmt.Print(string(flag[j] ^ x))
	// 		j += 1
	// 		powX += 1
	// 		// hkcert22{y0u_c4n_n4vig4t3_r4nd0mn0mn3s5_f}
	// 	}
	// }
	fmt.Println()
}
