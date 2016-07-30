package main

import (
	"log"
	"os"

	"gopkg.in/pipe.v2"
)

func main() {
	log.SetFlags(log.Lshortfile)

	// echo "starting" | awk '{print $0; system("sleep 5"); print "done"}'
	p := pipe.Line(
		pipe.Exec("echo", "starting"),
		pipe.Exec("awk", "{print $0; system(\"sleep 5\"); print \"done\"}"),
	)

	outb := &pipe.OutputBuffer{}
	s := pipe.NewState(outb, outb)
	err := p(s)
	if err == nil {
		err = s.RunTasks()
	}
	output := outb.Bytes()
	pipe_err := err

	if pipe_err != nil {
		log.Println(pipe_err)
	}

	os.Stdout.Write(output)
}
