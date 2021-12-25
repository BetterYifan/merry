package cmd

import (
	"bytes"
	"github.com/spf13/cobra"
	"log"
	"strconv"
	"strings"
)

var d2bCommand = &cobra.Command{
	Use:   "d2b [IPV4Addr]",
	Short: "transfer decimal ip to binary ip",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			log.Fatal("args cannot be empty")
			return
		}
		ipv4 := args[0]
		nums := strings.Split(ipv4, ".")
		var bina bytes.Buffer
		for _, num := range nums {
			dec, _ := strconv.ParseInt(num, 10, 64)
			b := d2b(dec)
			if b == "" {
				bina.Reset()
				break
			}
			bina.WriteString(b)
			bina.WriteString(" ")
		}

		log.Println(bina.String())
		return
	},
}

func d2b(n int64) string {
	result := make([]string, 8)
	for i := range result {
		result[i] = "0"
	}

	if n > 255 {
		return strings.Join(result, "")
	}
	if n == 0 {
		return strings.Join(result, "")
	}
	index := 0
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result[7-index] = strconv.Itoa(int(lsb))
		index++
	}

	return strings.Join(result, "")
}

var b2dCommand = &cobra.Command{}
