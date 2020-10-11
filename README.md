# go 101 project for learning Go Lang


## Characteristics

Go is a compiled, statically typed language. -> compiled to OS specific

go tool can run without compiling

Have static link runtime. -> size is larger because runtime is included. (no JVM like needed like Java)

Runtime is linked directly to every applications.

Go has some OOP features but not all. (Custom Types, Interfaces, Type inheritance)

Go does not have structured exception handling -> Error are more like JavaScript (Error object)

Case sensitivity

variables and packages -> start with lower and then mixed case

exported functions and fields have initial UpperCase

No semicolon needed

Starting braces on the same line

## OOP

Go can implement OOP concept 
 - Encapsulation with Types and Structs (Data structure)
 - Polymorphism with an interface (no Class inheritance)
 
 
 no overloading, overriding 
 
 ## Mac Installation
 
 ```
 brew update && brew install golang 
 ```

```
mkdir -p $HOME/go/{bin,src,pkg}
```
 
 This created folders `bin` and `src`

add this to your `.zshrc`

```
export GOPATH=$HOME/go
export GOROOT="$(brew --prefix golang)/libexec"
export PATH="$PATH:${GOPATH}/bin:${GOROOT}/bin"
```
 
reload settings with 

```
source $HOME/.bashrc
```

then create projects under 

```  
~/go/src
```

