package main

import (
	"encoding/csv"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/netip"
	"os"
	"strings"

	"github.com/psanford/cloudip"
)

var format = flag.String("format", "text", "Output format (text|csv|json)")

func main() {
	flag.Parse()
	if *format != "text" && *format != "csv" && *format != "json" {
		log.Fatalf("invalid -format. Valid values are: text,csv,json")
	}

	if len(flag.Args()) < 1 {
		log.Fatalf("Usage: %s <ip>\n", os.Args[0])
	}

	csvOut := csv.NewWriter(os.Stdout)
	jsonOut := json.NewEncoder(os.Stdout)
	jsonOut.SetIndent("", "  ")

	if *format == "csv" {
		csvOut.Write([]string{"prefix", "provider", "region", "services"})
	}

	exitCode := 0

	for _, ipStr := range flag.Args() {

		addr, err := netip.ParseAddr(ipStr)
		if err != nil {
			log.Fatalf("ip parse error: %s", err)
		}

		r := cloudip.Lookup(addr)
		if r == nil {
			log.Printf("ip %s not found", addr)
			exitCode = 1
			continue
		}

		switch *format {
		case "text":
			fmt.Printf("Prefix:   %s\n", r.Prefix)
			fmt.Printf("Provider: %s\n", r.Provider)
			if r.Region != "" {
				fmt.Printf("Region:   %s\n", r.Region)
			}
			if len(r.Services) > 0 {
				fmt.Printf("Service:  %s\n", strings.Join(r.Services, ","))
			}
			if len(flag.Args()) > 1 {
				fmt.Println("==================================================")
			}
		case "csv":
			services := strings.Join(r.Services, ";")
			csvOut.Write([]string{r.Prefix.String(), string(r.Provider), r.Region, services})
			csvOut.Flush()
		case "json":
			jsonOut.Encode(r)
		}
	}

	if exitCode != 0 {
		log.Print("Error: Some Lookups failed")
	}
	os.Exit(exitCode)
}
