

# Dummy Domain

This is a dummy domain module

It is a resource for 

(Working Locally WorkSpaces or replace)[]

The objective for WorkSpaces

You can replace the Workspace config for another directory and so change
the dependant behaviour. This can also be achieved through replace


---

This repo also shows the use of 

* [Fuzzing](https://go.dev/doc/security/fuzz/)
  see the makefile
    
    make fuzz_test

* (Benchmarks)[https://pkg.go.dev/testing#hdr-Benchmarks]
  there is a tool called [benchstat](https://pkg.go.dev/golang.org/x/perf/cmd/benchstat)
  where one can run two bench commands and compare them with bench stat
  see
    
    make bench_test

  install bench stat, run the command from a dir with no go.mod
    
    go install golang.org/x/perf/cmd/benchstat@latest

  Then you can run bench tests and save the output eg

    go test -bench=. ./stringer > origional.txt

    do some code changes and run again

    go test -bench=. ./stringer > new.txt

    Then us benchstat to comepare

    benchstat origional.txt new.txt
    



---

## Linting coverage
  
use [colangci-lint](https://github.com/golangci/golangci-lint?tab=readme-ov-file)
There is a ./.golangci.yml file that configures this tool. 

## coverage 

run
  
  make test_cover

note you can alter the -func in the makefile to -html to see the code which is covered or not

