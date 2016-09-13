# chole
A elegant tunnels of exposing private network to the internet.

## Install

`go get -u github.com/imeoer/chole`

## Usage

**# run on server:**

`chole -s`

**# configure rule in config.yaml:**

```
rules:
  web:
    out: 80
    in: 8000
  ssh:
    out: 22
    in: 192.168.1.2:22
```

**# run on local:**

`chole`