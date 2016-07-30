package main

import (
	"log"
	"os"
	"time"

	"gopkg.in/pipe.v2"
)

func main() {
	log.SetFlags(log.Lshortfile)

	// echo "starting" | awk '{print $0; system("sleep 5"); print "done"}'
	p := pipe.Line(
		pipe.Exec("echo", "starting"),
		pipe.Exec("awk", "{print $0; system(\"sleep 5\"); print \"done\"}"),
	)

	// setup the pipeline
	output := &pipe.OutputBuffer{}
	s := pipe.NewState(output, output)
	err := p(s)
	if err != nil {
		log.Fatalln(err)
	}

	// start the pipeline
	done := make(chan error)
	go func() {
		done <- s.RunTasks()
	}()

	// cancel the pipeline
	time.Sleep(6 * time.Second)
	s.Kill()

	// wait for the pipeline to finish
	err = <-done
	if err != nil {
		log.Println(err)
	}

	os.Stdout.Write(output.Bytes())
}
