---
title: Go Closure Capture Gotchas
kind: topic
sources:
  - source.md
  - raw/web/corpus-2026-05-18/034-go-faq.md
confidence: medium
---

# Go Closure Capture Gotchas

In Go, closures capture variables from their enclosing scope by reference. This behavior is often intentional but can lead to subtle bugs when the captured variable's value changes after the closure is created and before it is executed. These issues are commonly referred to as "closure capture gotchas." Understanding how the runtime handles these captures is essential for writing safe concurrent code using goroutines.

### The Common Pitfall: Loop Variables in Goroutines

A classic example of a closure capture gotcha occurs when iterating over a range loop and launching a new goroutine in each iteration. If the loop variable is captured directly, all goroutines may end up reading or printing the same final value of that variable.

```go
for i := 0; i < 5; i++ {
    go func() {
        fmt.Println(i) // Prints "5" five times
    }()
}
```

In this snippet, the loop variable `i` is not copied into each goroutine. Instead, all goroutines capture a reference to the same memory location where `i` resides. By the time any of them execute, the loop has finished and `i` holds the value `5`.

### The Solution: Capturing by Value

To avoid this issue, you must create a new variable inside the loop that shadows the outer one. This ensures each goroutine captures its own copy of the value at the moment the closure is created.

```go
for i := 0; i < 5; i++ {
    i := i // Create a new local variable
    go func() {
        fmt.Println(i) // Prints 0, 1, 2, 3, 4 correctly
    }()
}
```

### Underlying Mechanism: Escape Analysis

Go's escape analysis determines whether a variable should be allocated on the stack or the heap. In the first example above, the loop variable `i` escapes the stack frame of the function because it is accessed by closures (the goroutines) that outlive the current execution context. Consequently, it is placed on the heap, and all goroutines point to the same heap-allocated slot.

In the corrected version, the inner `i` does not escape the scope of the specific goroutine's closure, allowing it to reside on the stack or be copied efficiently depending on the compiler's optimization strategy. This distinction is a key aspect of Go's memory management model.

### Advanced Scenarios

Gotchas can also arise with more complex data structures, such as slices and maps. If a slice is captured in a closure, the closure holds a reference to the underlying array. Modifying the slice (e.g., appending elements) after the closure is created but before it runs can lead to unexpected behavior if multiple closures share the same slice backing store.

Similarly, capturing a map requires careful synchronization when writing to it concurrently, as map operations are not atomic by default. While reading from a captured map within a goroutine is generally safe without locks (due to Go's memory model ensuring consistent views), writing requires explicit mutex protection or channel-based coordination.
