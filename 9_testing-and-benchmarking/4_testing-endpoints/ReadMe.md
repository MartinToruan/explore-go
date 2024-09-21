1. In this module, you can see we can test a handler api without running the service.
2. In this module, you could also see a function ExampleSendJSON() to put an example on the godoc document

```azure
- Run a godoc server
  godoc -http=":3000"

- Open localhost:3000 on the brower
- Navigate to http://localhost:3000/pkg/github.com/MartinToruan/explore-go-in-action/9_testing-and-benchmarking/4_testing-endpoints/handler/#example_SendJSON
```
3. You could also run a test on the example by running the code below. 
```azure
go test -v -run="ExampleSendJSON"
```

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; Go will compare the return from the API with the output written on the test
```azure
    Output:
    {Kristopel kristopel@gmail.com}
```

4. Bechmark
```azure
Run All Benchmark
go test -v -run="none" -bench=. 

Run spesific Benchmark
go test -v -run="none" -bench="BenchmarkSprintf"

Run benchmark with minimum 3s
go test -v -run="none" -bench="BenchmarkSprintf" -benchtime="3s"

Showing memory usage and number of allocation to heap
go test -v -run="non" -bench="BenchmarkSprintf" -benchmem
```