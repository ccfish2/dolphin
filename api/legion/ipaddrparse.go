package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {

	fileName := "github.com/dolphin/api/legion/ip-ranges.json"

	fileData, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal("failed opening intput data %v", err)
		return
	}

	var ipranges map[string]interface{}
	if err := json.Unmarshal(fileData, &ipranges); err != nil {
		log.Fatal("Parse JSON file failure <%v>", err)
		return
	}

	mp := make(map[string]int)
	for key, value := range ipranges {
		if key == "prefixes" {
			for _, v := range value.([]interface{}) {
				for key, value := range v.(map[string]interface{}) {
					if key == "region" && strings.Compare(value.(string), "us-east-1") == 0 {
						mp[v.(map[string]interface{})["ip_prefix"].(string)] += 1
					}
				}
			}
		}
	}

	mostIPranges := "" //IPv4 ranges
	mostIPRangecnt := 0

	mostfreqsubnets := "" // most common subnets
	numberofHosts := 0

	for cidr, cnt := range mp {
		if cnt > mostIPRangecnt {
			mostIPRangecnt = cnt
			mostIPranges = cidr
		}

		ip, netip, err := net.ParseCIDR(cidr)
		if err != nil {
			log.Fatal("failed to parse CIDR")
			return
		}
		ones, bits := netip.Mask.Size()
		totalhosts := 1 << (bits - ones)

		if totalhosts*cnt > numberofHosts {
			numberofHosts = totalhosts * cnt
			mostfreqsubnets = fmt.Sprintf("%d", ones)
		}

		fmt.Println(" CIDR <", cidr, "> Range ", ip, " netip IP <", netip.IP, "> mask <", netip.Mask, "> total hosts <", totalhosts, ">", " subnets %d", ones)
	}

	fmt.Println(" CIDR range <", mostIPranges, "> and its count <", mostIPRangecnt, "> most common subnets <", mostfreqsubnets, " total hosts ", numberofHosts, " mostIPRangecnt ", mostIPRangecnt)
}
