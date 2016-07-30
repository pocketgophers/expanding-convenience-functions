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

	// setup the pipeline
	output := &pipe.OutputBuffer{}
	s := pipe.NewState(output, output)
	err := p(s)
	if err != nil {
		log.Fatalln(err)
	}

	// run the pipeline
	err = s.RunTasks()
	if err != nil {
		log.Println(err)
	}

	os.Stdout.Write(output.Bytes())
}
