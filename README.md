# Item Finder
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
git clone https://github.com/camcamsatnav/mcscanner.git
cd mcscanner
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

Even though this information can be found inside the project on machine-readable
format like in a .json file, it's good to include a summary of most useful
links to humans using your project. You can include links like:

- Project homepage: https://your.github.com/awesome-project/
- Repository: https://github.com/your/awesome-project/
- Issue tracker: https://github.com/your/awesome-project/issues
    - In case of sensitive bugs like security vulnerabilities, please contact
      my@email.com directly instead of using issue tracker. We value your effort
      to improve the security and privacy of this project!
- Related projects:
    - Your other project: https://github.com/your/other-project/
    - Someone else's project: https://github.com/someones/awesome-project/


## Licensing

One really important part: Give your project a proper license. Here you should
state what the license is and how to find the text version of the license.
Something like:

"The code in this project is licensed under MIT license."