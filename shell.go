package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func (e *env) read(in *os.File, wc, qc chan<- bool, iq chan<- string) {
	go func() {
		o := true
		for {
			if in == nil {
				in = os.Stdin
			}
			reader := bufio.NewReader(in)
			if o {
				fmt.Print(">>> ")
			} else {
				o = true
			}
			text, err := reader.ReadString('\n')
			if err != nil {
				e.logger("read", "", err)
				cleanDir(e.BldDir)
				qc <- true
				return
			}
			if e.parser.parseLine(text, iq) {
				wc <- true
				o = false
			}
		}
	}()
}

func (e *env) write(ec chan<- bool) {

	go func() {
		f, err := os.OpenFile(e.TmpPath, os.O_WRONLY|os.O_CREATE, 0600)
		if err != nil {
			return
		}
		time.Sleep(time.Microsecond)

		for _, l := range e.parser.convertLines() {
			f.WriteString(l)
		}
		f.Sync()
		if err := f.Close(); err != nil {
			e.logger("writer", "", err)
			return
		}
		ec <- true
		e.parser.main = []string{}
	}()
}

func (e *env) goRun(rc chan<- bool) {

	go func() {
		os.Chdir(e.BldDir)
		cmd := "go"
		args := []string{"run", tmpname}
		if err := runCmd(cmd, args...); err != nil {
			e.logger("build", "", err)
			return
		}
		rc <- true
	}()
}

func (e *env) shell() {

	qc := make(chan bool)
	rc := make(chan bool)
	wc := make(chan bool)
	ec := make(chan bool)
	iq := make(chan string, 10)

	e.read(nil, wc, qc, iq)
	goGet(<-iq)

loop:
	for {
		select {
		case <-wc:
			e.write(ec)
		case <-ec:
			e.goRun(rc)
		case <-qc:
			cleanDir(e.BldDir)
			fmt.Println("[gosh] terminated")
			break loop
		}
	}

	time.Sleep(time.Nanosecond)
	return
}
