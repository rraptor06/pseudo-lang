# ⌨️️ Pseudo-Lang 🗣️

## Table of contents 📑
- [Description](https://github.com/toro-nicolas/pseudo-lang/blob/main/README.md#description-)
- [Usage](https://github.com/toro-nicolas/pseudo-lang/blob/main/README.md#usage-%EF%B8%8F)
- [Compilation](https://github.com/toro-nicolas/pseudo-lang/blob/main/README.md#compilation-%EF%B8%8F)
- [Documentation](https://github.com/toro-nicolas/pseudo-lang/blob/main/README.md#documentation-)
- [Credits](https://github.com/toro-nicolas/pseudo-lang/blob/main/README.md#credits-)


## Description 📝
**Pseudo-Lang** is a programming language. Its aim is to **write fast programs easily**, **accessible to all** types of audience. It is based on the principle of the "[**pseudo-code**](https://en.wikipedia.org/wiki/Pseudocode)", writing code in an almost natural language. Here, our language is based on a kind of **pseudo-code in French** (if required, we can add other languages).  
Like every programming language, this one includes a **compiler**: **PLC (Pseudo-Lang Compiler)**. The compiler will convert the **Pseudo-Lang project** into a **C project** that can be accessed directly by the client.  

This github repository contains the **PLC source code** and all the **information linked to Pseudo-Lang**.  
For more information on the project, please visit the [github project](https://github.com/users/toro-nicolas/projects/2) or the [project website](https://toro-nicolas.github.io/pseudo-lang/).


## Usage ⚔️
Here's an example of a **PSL** (Pseudo-Lang) file :
```addition.psl```
```psl
afficher("Hello, World!\n")

fonction: additionner(entier a, entier b) -> entier
    afficher(a, " + ", b, " = ")
    retourner a + b
    
a <- 5
b <- 3
afficher(additionner(a, b), "\n")
```

You can run plc like this :
```sh
./plc addition.psl -o addition -r
```

For more information, please see the help section.
```sh
USAGE
        plc [OPTIONS] [file1.pl file2.pl ...]
DESCRIPTION
        The official Pseudo-Lang Compiler
OPTIONS
        -c              Only convert Pseudo-Lang to C, no compilation
        -d [directory]  Specify the code output directory
        -o [name]       Specify the name of the executable (default is main.out)
        -r              Run the program after compilation*
        -w              Disable warning display
```


## Compilation 🛠️
You can compile the **PLC project** with this command :
```sh
make
```

If you want clean the **PLC project**, you can run this command :
```sh
make clean
```

You can clean and compile the **PLC project** with ```make re``` .


## Documentation 📚
The documentation to the **Pseudo-Lang** is accessible [here](https://toro-nicolas.github.io/pseudo-lang/pseudo-lang/).

To open the **PLC documentation**, run the command :
```sh
make doc
```


## Credits 👤

This project, **Pseudo-Lang** and **its compiler**, was totally **imagined and built** by [**Nicolas TORO**](https://github.com/toro-nicolas) as part of a **HUB project** for his first year in [**EPITECH's Grande Ecole program**](https://www.epitech.eu/programme-grande-ecole-informatique/).
