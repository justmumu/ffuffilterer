# FFUF Filterer

FFuf Filterer simply filters json output.

### Features!

 All of ffuf filter options except regex.

### Installation

Firstly, you should install golang language and you should configure $GOPATH variable.
After that,

``` $ go get github.com/rasity/ffuffilterer ```

### Help Preview
```sh
$ ffuffilterer --help
Usage:
  ffuffilterer [OPTIONS]

Application Options:
  -f, --file= File path should be filtered
      --host= Include host filter. (Can be use multiple times)
      --ws=   Exclude size option. (Can be use multiple times)
      --ww=   Exclude words option. (Can be use multiple times)
      --wl=   Exclude lines option. (Can be use multiple times)
      --wc=   Exclude status code option. (Can be use multiple times)
      --fs=   Include size option. (Can be use multiple times)
      --fw=   Include words option. (Can be use multiple times)
      --fl=   Include lines option. (Can be use multiple times)
      --fc=   Include status code option. (Can be use multiple times)

Help Options:
  -h, --help  Show this help message
```

