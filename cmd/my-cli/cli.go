package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/url"
	"os"
	"strconv"
	"time"

	"github.com/go-ping/ping"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "CLI tool that makes a request to a website"
	app.Usage = "Let's you fetch a JSON Object from an API"

	requestFlags := []cli.Flag{
		&cli.StringFlag{
			Name:  "url",
			Value: "https://software_assessment.yasharavindra-wrangler.workers.dev/links",
		},
		&cli.IntFlag{
			Name:  "profile",
			Value: 1,
		},
	}
	// create commands
	app.Commands = []*cli.Command{
		{
			Name:  "runTool",
			Usage: "Provides JSON Object containing links",
			Flags: requestFlags,

			// the action, or code that will be executed when
			// we execute our `runTool` command
			Action: func(c *cli.Context) error {

				//check if flag is `profile`
				if os.Args[2] == "--profile" {
					var success = int64(0)
					var fastest_time = make([]time.Duration, c.Int(os.Args[3]))
					var slowest_time = make([]time.Duration, c.Int(os.Args[3]))
					var mean_time = make([]time.Duration, c.Int(os.Args[3]))

					//ping the url specified in the console
					pinger, err := ping.NewPinger(c.String("url"))
					if err != nil {
						fmt.Printf("ERROR: %s\n", err.Error())
						return nil
					}

					//returns ping statistics on successful ping
					pinger.OnFinish = func(stats *ping.Statistics) {
						success++
						fastest_time = append(fastest_time, stats.MinRtt)
						slowest_time = append(slowest_time, stats.MaxRtt)
						mean_time = append(mean_time, stats.AvgRtt)

					}

					pinger.Count = 1
					pinger.Interval = time.Second
					pinger.Timeout = time.Second * 100000
					pinger.SetPrivileged(true)

					fmt.Printf("PING %s (%s):\n", pinger.Addr(), pinger.IPAddr())

					//positive integer accepting the number of requests is stored in `val`
					val, err := strconv.ParseInt(os.Args[3], 10, 64) //conversion of hex to int

					//pinger runs number of times specified in the command prompt
					for i := int64(0); i < val; i++ {
						err = pinger.Run()
					}

					var max time.Duration = slowest_time[0]
					var min time.Duration = fastest_time[0]
					var total time.Duration = mean_time[0]

					for i := int64(1); i < val; i++ {
						if fastest_time[i] < min {
							min = fastest_time[i]
						}
						if slowest_time[i] > max {
							max = slowest_time[i]
						}
						total = total + mean_time[i]
					}

					fmt.Printf("\n--- Ping Statistics ---\n")
					fmt.Printf("The fastest time is: %v \n", min)
					fmt.Printf("The slowest time is: %v \n", max)
					fmt.Printf("The mean time is: %v \n", time.Duration(int64(total)/val))
					fmt.Printf("Percentage requests that succeeded: %.2f percent", float64(success/(val)*100))

					if err != nil {
						fmt.Printf("Failed to ping target host: %s", err)
					}

					//check if flag is `url`
				} else if os.Args[2] == "--url" {
					s := c.String("url")
					u, err := url.Parse(s)
					if err != nil {
						log.Fatal(err)
					}

					//establishing tcp connection with port 80
					conn, err := net.Dial("tcp", u.Host+":80")
					if err != nil {
						log.Fatal(err)
					}

					//setting http request header
					rt := fmt.Sprintf("GET %v HTTP/1.1\r\n", u.Path)
					rt += fmt.Sprintf("Host: %v\r\n", u.Host)
					rt += fmt.Sprintf("Connection: close\r\n")
					rt += fmt.Sprintf("\r\n")

					_, err = conn.Write([]byte(rt))
					if err != nil {
						log.Fatal(err)
					}

					resp, err := ioutil.ReadAll(conn)
					if err != nil {
						log.Fatal(err)
					}
					fmt.Println(string(resp))

					//close connection
					conn.Close()
					return nil
				}
				return nil
			},
		},
	}

	// start our application
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
