package utils

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
)

const (
	// See http://man7.org/linux/man-pages/man8/route.8.html
	route = "/proc/net/route"
)

func ResolveGateway() (iface string, ip net.IP, err error) {
	f, err := os.Open(route)
	if err != nil {
		return "", nil, err
	}
	defer f.Close()
	bytes, err := ioutil.ReadAll(f)
	if err != nil {
		return "", nil, err
	}
	return parseLinuxProcNetRoute(bytes)
}

func ResolveDefaultIP() string {
	defaultIface, _, err := ResolveGateway()
	if err != nil {
		log.Print(err)
		return ""
	}
	iface, err := net.InterfaceByName(defaultIface)
	if err != nil {
		log.Print(err)
		return ""
	}
	addrs, _ := iface.Addrs()
	//TODO handle err
	for _, addr := range addrs {
		// check the address type and if it is not a loopback the display it
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}

func ResolveLocalIPs() []map[string]string {
	var ips []map[string]string
	itfs, _ := net.Interfaces()
	for _, itf := range itfs {
		name := itf.Name
		ip := ResolveInterfaceIP(name)
		if ip != "" {
			ips = append(ips, map[string]string{
				name: ip,
			})
		}
	}
	return ips
}

func ResolveHostIP(host string) string {
	addrs, _ := net.LookupIP(host)
	var ip net.IP
	for _, addr := range addrs {
		if !addr.IsLoopback() && addr.To4() != nil {
			ip = addr
		}
	}
	if ip != nil {
		return ip.String()
	} else {
		return ""
	}
}

func ResolveInterfaceIP(eth string) string {
	iface, err := net.InterfaceByName(eth) //here your interface
	if err != nil {
		log.Print(err)
		return ""
	}
	addrs, _ := iface.Addrs()
	var ip net.IP
	for _, addr := range addrs {
		switch v := addr.(type) {
		case *net.IPNet:
			if !v.IP.IsLoopback() {
				if v.IP.To4() != nil { //Verify if IP is IPV4
					ip = v.IP
				}
			}
		}
	}
	if ip != nil {
		return ip.String()
	} else {
		return ""
	}
}

func parseLinuxProcNetRoute(f []byte) (iface string, gateway net.IP, err error) {
	const (
		sep              = "\t" // field separator
		ifaceField       = 0    // field containing iface name
		destinationField = 1    // field containing hex destination address
		gatewayField     = 2    // field containing hex gateway address
	)
	scanner := bufio.NewScanner(bytes.NewReader(f))
	// Skip header line
	if !scanner.Scan() {
		return "", nil, errors.New("invalid linux route file")
	}
	for scanner.Scan() {
		row := scanner.Text()
		tokens := strings.Split(row, sep)
		iface := tokens[ifaceField]
		if len(tokens) <= gatewayField {
			return "", nil, fmt.Errorf("invalid row '%s' in route file", row)
		}
		// Cast hex destination address to int
		destinationHex := "0x" + tokens[destinationField]
		destination, err := strconv.ParseInt(destinationHex, 0, 64)
		if err != nil {
			return "", nil, fmt.Errorf(
				"parsing destination field hex '%s' in row '%s': %w",
				destinationHex,
				row,
				err,
			)
		}
		// The default interface is the one that's 0
		if destination != 0 {
			continue
		}
		gatewayHex := "0x" + tokens[gatewayField]
		// cast hex address to uint32
		d, err := strconv.ParseInt(gatewayHex, 0, 64)
		if err != nil {
			return "", nil, fmt.Errorf(
				"parsing default interface address field hex '%s' in row '%s': %w",
				destinationHex,
				row,
				err,
			)
		}
		d32 := uint32(d)
		// make net.IP address from uint32
		ipd32 := make(net.IP, 4)
		binary.LittleEndian.PutUint32(ipd32, d32)
		// format net.IP to dotted ipV4 string
		return iface, net.IP(ipd32), nil
	}
	return "", nil, errors.New("interface with default destination not found")
}
