- BUILD 
go build -o "./wind.exe" -ldflags "-H=windowsgui" .
# UI Framework Prototype on Raylib (Go)

**Languages:** [Русский](README.ru.md) | [English](README.md) | [Українська](README.ua.md)

## Project Description

This project is a prototype of a minimal UI framework written in Go using **raylib-go** (`github.com/gen2brain/raylib-go/raylib`). Its main goal is to explore and implement core UI concepts such as visual elements, layering, animations, and user input handling, while maintaining a clean and extensible application architecture.

The project is not intended to be a full-featured UI engine, but rather an experimental and educational foundation for understanding how UI systems can be built from scratch.

## Features

* Rendering and input handling via **raylib-go**
* UI elements based on `Rectangle`
* Basic UI animations
* Layer system:

  * draw order management
  * proper element overlapping
* UI interactions:

  * hover / unhover
  * left / right click
  * drag & drop
* Correct mouse input processing
* Dynamic element creation and removal
* Architecture focused on maintainability and extensibility

## Architecture

The project is centered around the `Application` entity, responsible for:

* application lifecycle (init / update / render / close)
* UI element management
* input processing
* tracking hovered and active elements

All UI components implement a common `Element` interface, allowing safe sorting, interaction handling, and easy extension with new element types.

## Layer System

Each UI element has an associated `Layer`. Elements are sorted in ascending layer order, ensuring that elements with higher layers are drawn last and appear visually on top of others.

## Project Goals

* Learn and practice **raylib-go** usage
* Understand core UI system design principles
* Implement proper user input handling
* Avoid index-based UI interaction logic
* Design a clean and extensible architecture
* Build a solid foundation for future features (buttons, containers, layout systems)


## Author

**mykytaserdiuk**

GitHub: [https://github.com/mykytaserdiuk](https://github.com/mykytaserdiuk)
