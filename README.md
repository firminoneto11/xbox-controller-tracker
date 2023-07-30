# Xbox Controller Tracker

## A little program written in Go that tracks whenever the 'Guide' button from a Xbox Controller is pressed and then presses a key from keyboard

## Installing

- Download this repository

- Make sure you have Go 1.20 or above

- Install a C Compiler for Windows: [GCC](https://jmeubank.github.io/tdm-gcc/download/)

- Install SDL (Check the instructions bellow)

- Install ZLIB (Check the instructions bellow)

## Installing SDL

In this project, theres a `sdl.zip` file. After you have successfully installed the TDM GCC for windows, unzip this `sdl.zip` file into the `C:\TDM-GCC-64\x86_64-w64-mingw32` folder. That's it! SDL is installed.

Keep in mind that the `C:\TDM-GCC-64` folder should be the path where you installed TDM GCC.

## Installing ZLIB

The process of installing zlib is similar to the SDL's. In the project's root directory theres a file named `zlib.zip`. Just unzip it into the `C:\TDM-GCC-64\x86_64-w64-mingw32`.

Keep in mind that the `C:\TDM-GCC-64` folder should be the path where you installed TDM GCC.

### The same code has also been written in python. Check the ./python directory if you'd like!
