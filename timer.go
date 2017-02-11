package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"time"
)

func printf(fmtstr string, args ...interface{}) (int, error) {
	return fmt.Printf("\a"+time.Now().Format("2006-01-02 03:04:05.000000000")+" "+fmtstr, args...)
}

func main() {
	var (
		err        error
		resp       string
		resp_bytes []byte
		delay      time.Duration
		reader     *bufio.Reader
	)

	reader = bufio.NewReader(os.Stdin)
	for {
		printf("resume the suffering? ")
		resp_bytes, _, err = reader.ReadLine()
		if err != nil {
			if err != io.EOF {
				printf("stdin.ReadLine: error: %v\n", err)
			}
			break
		}
		resp = string(resp_bytes)
		if len(resp) > 2 && resp[0] == '=' {
			var delaystr string
			fmt.Sscanf(resp, "=%s", &delaystr)
			delay, err = time.ParseDuration(delaystr)
			if err != nil {
				printf("error parsing new delay %q: %v\n", delaystr, err)
				delay = time.Duration(0)
			} else {
				printf("set delay to %s\n", delay)
			}
		}
		time.Sleep(delay)
	}
}
