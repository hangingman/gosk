# Gosk

It's a play-ground

## Build & Run

```
$ go get github.com/pointlander/peg               # only first time
$ git clone https://github.com/HobbyOSs/gosk.git  # clone
$ cd gosk
$ peg gosk.peg                                    # generate go source from PEG
$ go build                                        # build
$ ./gosk xxx.nas                                  # read nask file
```