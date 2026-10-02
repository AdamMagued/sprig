# Network Functions

Sprig network manipulation functions.

## getHostByName

The `getHostByName` receives a domain name and returns the ip address.

```
getHostByName "www.google.com" would return the corresponding ip address of www.google.com
```

## joinHostPort

The `joinHostPort` combines host and port into a network address (`host:port` or `[host]:port` for IPv6).

```
joinHostPort "127.0.0.1" 8080
```

The above returns: `127.0.0.1:8080`

```
joinHostPort "::1" 443
```

The above returns: `[::1]:443`

