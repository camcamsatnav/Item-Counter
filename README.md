# Item Counter
> A program to count all of a given item written in go.

The program will take an item and search a whole world directory for the given item and return a count

## Installing / Getting started

Download the latest release from the releases page and run in command line.

```shell
.\main.exe -d <directory> -s <item-to-search>
```

You should get a total count of the given item within that world.

## Examples
Type: `String`
Flag: `-d`

The directory to search for the item in.

Type: `String`
Flag: `-s`

The item you are trying to count

```bash
.\main.exe -d C:\Users\camde\AppData\Roaming\.minecraft\saves\world -s redstone
```

If your world directory has spaces in it, you will need to wrap it in quotes.

```bash
.\main.exe -d "C:\Users\camde\AppData\Roaming\.minecraft\saves\my world" -s redstone
```

## Developing

How to develop this project further:

```shell
git clone https://github.com/camcamsatnav/Item-Counter.git
cd Item-Counter
```

### Building

```shell
go mod tidy
go build
```

The first command adds missing dependencies and removes unused ones.
Go build compiles the program into an executable.

## Features

* Searches your world directory for a given item
* Counts super fast

## Links

- Repository: https://github.com/camcamsatnav/Item-Counter
- Issue tracker: https://github.com/camcamsatnav/Item-Counter/issues
- Related projects:
    - Better project: https://github.com/SciCraft/mc-scanner
    - Used for opening chunks: https://github.com/Tnze/go-mc
    - Used for parsing nbt files: https://github.com/midnightfreddie/nbt2json


## Licensing

The code in this project is licensed under MIT license.