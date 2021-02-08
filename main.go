package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"rogchap.com/v8go"
)

func main() {
	// https://github.com/rogchap/v8go/pull/68
	iso, _ := v8go.NewIsolate()
	global, _ := v8go.NewObjectTemplate(iso)
	sleep, _ := v8go.NewFunctionTemplate(iso, func(info *v8go.FunctionCallbackInfo) *v8go.Value {
		fmt.Printf("[sleep] %+v\n", info.Args())
		time.Sleep(100 * time.Millisecond)
		return nil
	})
	global.Set("sleep", sleep, v8go.ReadOnly)
	println, _ := v8go.NewFunctionTemplate(iso, func(info *v8go.FunctionCallbackInfo) *v8go.Value {
		fmt.Printf("%+v\n", info.Args())
		return nil
	})
	global.Set("println", println, v8go.ReadOnly)

	ctx, err := v8go.NewContext(iso, global)
	if err != nil {
		log.Fatalln("v8go.NewContext", err)
	}

	//
	//
	//

	mainJs, err := ioutil.ReadFile("main.js")
	if err != nil {
		log.Fatalln("ioutil.ReadFile", err)
	}
	if _, err := ctx.RunScript(string(mainJs), "runtime.js"); err != nil {
		log.Fatalln("ctx.RunScript 1", err)
	}

	//
	//
	//

	vals := make(chan *v8go.Value, 1)
	errs := make(chan error, 1)
	go func() {
		indexJs, err := ioutil.ReadFile("index.js")
		if err != nil {
			log.Fatalln("ioutil.ReadFile", err)
		}
		val, err := ctx.RunScript(string(indexJs), "runtime.js")
		if err != nil {
			errs <- err
			return
		}
		vals <- val
	}()
	select {
	case val := <-vals:
		log.Printf("val: %s, %#v", val, val)
	case err := <-errs:
		log.Fatalln("ctx.RunScript", err)
	case <-time.After(10 * time.Second):
		vm, _ := ctx.Isolate()  // get the Isolate from the context
		vm.TerminateExecution() // terminate the execution
		err := <-errs           // will get a termination error back from the running script
		log.Fatalln("ctx.RunScript", err)
	}
}
