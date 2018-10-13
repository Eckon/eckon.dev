# eckon.rocks
## Information to setup the server

> build server
```bash
$ sudo go build *.go
```

> run server
```bash
$ sudo ./main
```

> open new screen
```bash
$ screen -S "name"
```

> detach screen
```
cntr-a and cntr-d
```

> alternative
```bash
$ code & disown
```

> see all the screens
```bash
$ screen -list
```

> reatach screen
```bash
$ screen -r "name"
```

> kill process on port 80 (fuck apache)
```bash
$ sudo kill $(sudo lsof -t -i:80)
```