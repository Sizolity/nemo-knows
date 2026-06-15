---
title: Chunk 6 Notes
kind: topic
sources:
  - raw/web/corpus-2026-05-18/031-effective-go.md
confidence: medium
---

Chunk Context
This chunk covers advanced Go topics including embedding vs. subclassing, concurrency models ("Share by communicating"), goroutine mechanics, channel usage (buffered/unbuffered), parallelization strategies, garbage-collected buffer management via channels, and error handling conventions using the `error` interface and `panic`.

Local Summary
The section explains how embedding types allows method forwarding while keeping the receiver on the embedded type. It contrasts this with subclassing behavior. Concurrency principles are introduced, emphasizing that shared memory should be avoided in favor of communication via channels. Goroutines are described as lightweight functions managed by the Go runtime. Channels provide synchronization and data exchange; unbuffered channels act as rendezvous points, while buffered ones can limit concurrency (semaphores). Parallelization across CPU cores is achieved by launching goroutines for independent tasks and collecting completion signals. Finally, error handling patterns using the `error` interface and panic recovery are discussed.

Key Claims
- Embedding a type makes its methods available on the outer type, but the receiver remains the embedded type.
- Go encourages avoiding shared mutable state in favor of passing data through channels to prevent data races.
- Goroutines are lightweight functions that run concurrently within the same address space; they are multiplexed onto OS threads.
- Unbuffered channels synchronize senders and receivers, while buffered channels can limit throughput or act as semaphores.
- Errors in Go are typically returned as second return values of type `error`, which implements an interface with a single `Error() string` method.
- The `panic` function causes the program to stop immediately unless recovered from; it is intended for unrecoverable errors or impossible conditions.

Entities And Concepts
- Embedding (Go composite literals, field embedding)
- Goroutine (`go` keyword, lightweight threads)
- Channel (`chan`, buffered/unbuffered, send/receive operations)
- Semaphore pattern using channels
- Parallelization via CPU core detection (`runtime.NumCPU`)
- Error handling (`error` interface, `os.PathError`)
- Panic and recovery (though only panic definition is covered here)

Procedures And API Details
```go
type Job struct {
    Command string
    *log.Logger
}

func (job *Job) Printf(format string, args ...interface{}) {
    job.Logger.Printf("%q: %s", job.Command, fmt.Sprintf(format, args...))
}

go func() {
    time.Sleep(delay)
    fmt.Println(message)
}()

c := make(chan int)
go func() {
    list.Sort()
    c <- 1
}()
<-c

func Serve(queue chan *Request) {
    for req := range queue {
        sem <- 1
        go func() {
            process(req)
            <-sem
        }()
    }
}

func handle(queue chan *Request) {
    for r := range queue {
        process(r)
    }
}

func Serve(clientRequests chan *Request, quit chan bool) {
    for i := 0; i < MaxOutstanding; i++ {
        go handle(clientRequests)
    }
    <-quit
}

type Request struct {
    args   []int
    f      func([]int) int
    resultChan chan int
}

func (v Vector) DoSome(i, n int, u Vector, c chan int) {
    for ; i < n; i++ {
        v[i] += u.Op(v[i])
    }
    c <- 1
}

var numCPU = runtime.GOMAXPROCS(0)

func client() {
    for {
        var b *Buffer
        select {
        case b = <-freeList:
            // Got one; nothing more to do.
        default:
            b = new(Buffer)
        }
        load(b)
        serverChan <- b
    }
}

func server() {
    for {
        b := <-serverChan
        process(b)
        select {
        case freeList <- b:
            // Buffer on free list; nothing more to do.
        default:
            // Free list full, just carry on.
        }
    }
}

func CubeRoot(x float64) float64 {
    z := x / 3
    for i := 0; i < 1e6; i++ {
        prevz := z
        z -= (z*z*z - x) / (3*z*z)
        if veryClose(z, prevz) {
            return z
        }
    }
    panic(fmt.Sprintf("CubeRoot(%g) did not converge", x))
}

type error interface {
    Error() string
}

type PathError struct {
    Op   string // "open", "unlink", etc.
    Path string // The associated file.
    Err  error  // Returned by the system call.
}

func (e *PathError) Error() string {
    return e.Op + " " + e.Path + ": " + e.Err.Error()
}
```

Nuance Or Contradictions
- Before Go 1.22, using a loop variable inside a goroutine launched in a loop was buggy because the loop variable is shared across all goroutines unless captured carefully.
- While channels are first-class values and can be passed around, care must be taken not to create unbounded resource consumption when spawning goroutines per request without limiting concurrency.
- The distinction between concurrency (structuring as independent components) and parallelism (executing computations in parallel for efficiency) is emphasized; Go supports concurrency but does not guarantee parallel execution unless explicitly orchestrated.

Candidate Wiki Hints
- Embedding vs Subclassing
- Concurrency Model: Share by Communicating
- Goroutines and Channels Basics
- Channel Types: Buffered vs Unbuffered
- Semaphore Pattern with Channels
- Parallelizing Workloads Using CPU Cores
- Error Handling in Go
- Panic Usage and Recovery
