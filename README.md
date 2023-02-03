# testmemoryleak
## Documentations
### Init
   go mod tidy

### Single Test
- terminal 1

      go run main.go

- terminal 2

      curl --request POST localhost:8090/check --data '{"url": "https://freetestdata.com/wp-content/uploads/2021/09/png-5mb-1.png"}'

- terminal 3

      go tool pprof -http=:8091 localhost:6060/debug/pprof/heap

### Multiple Test
- terminal 1

      go run main.go

- terminal 2

      while true; do 
         curl --request POST localhost:8090/check --data '{"url":"https://freetestdata.com/wp-content/uploads/2021/09/png-5mb-1.png"}'
      sleep 0.5
      done

- terminal 3

      go tool pprof -http=:8091 localhost:6060/debug/pprof/heap

