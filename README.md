# eckon.dev
## Information to setup the server

> build server  
`$ sudo go build *.go`

> run server  
`$ sudo ./main`

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
