# eckon.rocks
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
`$ sudo certbot certonly --standalone -d eckon.rocks -d www.eckon.rocks`


Congratulations! Your certificate and chain have been saved at:
   /etc/letsencrypt/live/eckon.rocks/fullchain.pem
   Your key file has been saved at:
   /etc/letsencrypt/live/eckon.rocks/privkey.pem
   Your cert will expire on 2019-01-20. To obtain a new or tweaked
   version of this certificate in the future, simply run certbot
   again. To non-interactively renew *all* of your certificates, run
