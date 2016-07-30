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
	output, pipe_err := pipe.CombinedOutput(p)

	if pipe_err != nil {
		log.Println(pipe_err)
	}

	os.Stdout.Write(output)
}
