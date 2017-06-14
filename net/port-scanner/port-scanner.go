package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"regexp"
	"strconv"
	"strings"
	"sync"
)

var port int
var portRange string

var parallelCounts int

func init() {

	flag.IntVar(&port, "p", 80, "port")
	flag.StringVar(&portRange, "r", "", "range ports. format is <from>~<to>. eg. 100~200")
	flag.IntVar(&parallelCounts, "n", 1, "parallel counts")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "\nUsage: %s [Options] <IP>\n\nOptions:\n\n", os.Args[0])
		flag.PrintDefaults()
	}

	flag.Parse()

}

func printOpeningPort(port int) {

	fmt.Println("port " + strconv.Itoa(port) + " is opening")

}

func checkPort(ip net.IP, port int, wg *sync.WaitGroup, parallelChan *chan int) {

	defer wg.Done()

	tcpAddr := net.TCPAddr{
		IP:   ip,
		Port: port,
	}

	conn, err := net.DialTCP("tcp", nil, &tcpAddr)

	if err == nil {
		printOpeningPort(port)
		conn.Close()

	}

	<-*parallelChan

}

func main() {

	args := flag.Args()

	if len(args) != 1 {
		flag.Usage()
	} else {

		ip := net.ParseIP(flag.Arg(0))

		wg := sync.WaitGroup{}

		if portRange != "" {

			matched, _ := regexp.Match(`^\d+~\d+$`, []byte(portRange))

			if !matched {

				flag.Usage()

			} else {

				portSecs := strings.Split(portRange, "~")

				startPort, err1 := strconv.Atoi(portSecs[0])
				endPort, err2 := strconv.Atoi(portSecs[1])

				if err1 != nil || err2 != nil || startPort < 1 || endPort < 2 || endPort <= startPort || parallelCounts < 1 {
					flag.Usage()
				} else {

					wg.Add(endPort - startPort + 1)

					parallelChan := make(chan int, parallelCounts)

					for i := startPort; i <= endPort; i++ {

						parallelChan <- 1

						go checkPort(ip, i, &wg, &parallelChan)

					}

					wg.Wait()

				}

			}

		} else {

			wg.Add(1)

			parallelChan := make(chan int)

			go func() {
				parallelChan <- 1
			}()

			go checkPort(ip, port, &wg, &parallelChan)

			wg.Wait()

		}

	}

}
