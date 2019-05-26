# eckon.dev
## Information to setup the server

> build and start server
`$ go build -o bin/server && sudo ./bin/server`

> open new screen  
`$ screen -S "name"`

> detach screen  
`cntr-a and cntr-d`

> alternative  
`$ code & disown`

> see all the screens  
`$ screen -list`

> reatach screen  
`$ screen -r "name"`

> kill process on port 80 (fuck apache)  
`$ sudo kill $(sudo lsof -t -i:80)`

> certificate (certbot)
`$ sudo certbot certonly --standalone -d eckon.dev -d www.eckon.dev`

> using the new modules in go
`$ go mod init`
`$ go mod tidy`

> using modules to import different packages from local
`replace eckon.dev/src => ./src`
`require (eckon.dev/src v0.0.0)`
