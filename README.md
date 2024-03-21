- [roll - dice rolling from the command line](#roll---dice-rolling-from-the-command-line)
	- [About](#about)
	- [Build and Install](#build-and-install)
		- [Makefile](#makefile)
		- [Building with Go](#building-with-go)
	- [Example usage](#example-usage)
		- [Roll a single die](#roll-a-single-die)
		- [Rolling multiple different dice](#rolling-multiple-different-dice)
		- [Using modifiers](#using-modifiers)
		- [Summing up multiple dice rolls](#summing-up-multiple-dice-rolls)
		- [Print usage](#print-usage)
	- [Library usage](#library-usage)
	- [License](#license)


# roll - dice rolling from the command line

## About

`roll` is a dice rolling command line toy. Useful for D&D enthusiast who are not shying away from the command line.

## Build and Install

Clone the source, and build/install by using either the supplied makefile or Go from the source directory.

### Makefile

``` sh
sudo make clean install
```

### Building with Go

``` sh
go build
sudo mv ./roll /usr/local/bin
```

## Example usage

### Roll a single die

To roll a die, you will need to use the standard _D&D notation_:

``` text
[number of dice]d<dice sides>[+modifier]
```

For example:

``` sh
roll 5d6
```

The number in front of the "d" sign can be omitted, in which case a single die will be rolled.

More simple examples:
   - d8
   - 1d10
   - 6d6

Dice sides should be of the following list of valid dices:
   - 4
   - 6
   - 8
   - 10
   - 12
   - 20
   - 100

### Rolling multiple different dice

Rolling multiple different dice is as easy as listing them as separate command line arguments, separated with a space. Each roll will be printed to a separate line for easier parsing.

``` sh
roll 2d6 1d4 1d8
```

### Using modifiers

You can also use modifiers when supplying the dice definition. A modifier can either be a positive (`+`) or negative (`-`) value.

``` sh
roll 1d6+3
roll 1d8-1
```

When using negative modifiers, you may want to allow the result to go below 0. You can do this by specifying the `-n` flag.

``` sh
roll 1d4-10 # result will be 0
```

``` sh
roll -n 1d4-10 # result will be negative
```

### Summing up multiple dice rolls

If you want to see the sum of the rolls instead of the individual rolls, you can use the `-s` flag. This can be done with a single roll, or multiple rolls.

``` sh
roll -s 3d4 2d6 1d10
```

When you combine negative modifiers and the `-s` flag, none of the rolls can go below 0. So unless you use the `-n` flag, this will cause the sum to be greater than or equal to 0 as well.

``` sh
roll -s 1d4-10 1d6-20 # result will be 0
```

``` sh
roll -s -n 1d4-10 1d6-20 # result will be negative
```

### Print usage

You can print the usage with the `-h` flag:

``` sh
roll -h
```

```
Usage:
	roll [roll-def] ...
		roll-def: [count]d<sides>[+modifier]
		sides: [4,6,8,10,12,20,100]

Examples:
	d6	rolls one six-sided die
	1d10	rolls one ten-sided die
	1d10+1	rolls one ten-sided die and adds 1 to the result
	1d10-1	rolls one ten-sided die and subtracts 1 from the result
	3d4	rolls three four-sided dice
	d4 d6	rolls a four-sided die and a six-sided die

Flags:
  -n	allow rolls to go below 0 when using negative modifiers
  -s	print a single line summing up all dice
```


## Library usage


Add to your project:

```
go get github.com/gar-r/roll
```

Usage:

```go
package main

import (
	"fmt"

	"github.com/gar-r/roll/dice"
)

func main() {
	fmt.Printf("1d6:    %d\n", dice.D6.Roll())
	fmt.Printf("3d10:   %d\n", dice.D10.RollMany(3))
	fmt.Printf("2d20+1: %d\n", dice.D20.RollSpecial(2, 1))
}
```

## License

This application and the source falls under the _GNU GENERAL PUBLIC LICENSE_.
If you find it useful, please consider referring it to a friend, or contributing to the source.
