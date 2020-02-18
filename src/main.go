package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/kyokomi/emoji"
)

// define scanner.
var sc = bufio.NewScanner(os.Stdin)

// nextInt checks if input is invalid as the iteration number and returns it.
func nextInt() (int, error) {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil || i <= 0 {
		return -1, errors.New("invalid iteration number")
	}
	return i, nil
}

// nextText returns input text.
func nextText() string {
	sc.Scan()
	return sc.Text()
}

// toSecond convert the formatted time to second.
// For example, 3m30s will be 3*60 + 30.
func toSecond(s string) (int, error) {
	// regexp match
	return 0, nil
}

// timeItoa returns time as 2-digit style.
func timeItoa(i int) string {
	var s string
	if i < 10 {
		s = "0" + strconv.Itoa(i)
	} else {
		s = strconv.Itoa(i)
	}
	return s
}

// tick displays the time for given seconds.
// TODO: change the position of cursor.
func tick(second int) {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	done := make(chan bool)
	go func() {
		time.Sleep(time.Duration(second) * time.Second)
		done <- true
	}()
	var i int
	fmt.Println()
	for {
		select {
		case <-done:
			fmt.Println()
			return
		case t := <-ticker.C:
			emoji.Printf(" :clock%d:%s:%s:%s\r", (i%12)+1, timeItoa(t.Hour()), timeItoa(t.Minute()), timeItoa(t.Second()))
		}
	}
}

// main
func main() {
	emoji.Println(":bellhop::bellhop: Welcome to EasyGoBell ! :bellhop::bellhop:")

	var n int
	gotNum := false
	for {
		if gotNum {
			break
		}
		emoji.Print("\n:pushpin:How many times do you want to ring bell? ")
		ntmp, err := nextInt()
		if err != nil {
			emoji.Println(":dizzy:invalid number: try again!")
		} else {
			n = ntmp
			gotNum = true
		}
	}

	emoji.Println("\n:memo:Please input times you want to ring bell!")
	time.Sleep(500 * time.Millisecond)
	fmt.Println("   [Format] Xminutes&YYseconds -> XmYYs")
	time.Sleep(time.Second)

	var secs []int
	i := 0
	for {
		if i >= n {
			break
		}
		emoji.Printf("\n:hourglass:Time %d: ", i+1)
		timeDescs := nextText()
		sec, err := toSecond(timeDescs)
		if err != nil {
			emoji.Println(":dizzy:invalid format: input again!")
			continue
		}
		secs = append(secs, sec)
		i++
	}

	fmt.Println(secs) //
	// TODO: ticks and ring bell
	tick(10) //
}
