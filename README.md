# truck
TCP Proxy to UDP services.

## Install
```
go get [-u] github.com/microwaves/truck
```

## Use
```
$ truck --target target.udp.host:1337 --port 7777
```

`localhost:7777` should now accept TCP connections and proxy to your UDP service. :unicorn:

## Authors

Stephano Zanzin - [@microwaves](https://github.com/microwaves)

## License
See [LICENSE](LICENSE)
