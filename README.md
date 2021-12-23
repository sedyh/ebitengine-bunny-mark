# <img align="right" src="https://user-images.githubusercontent.com/19890545/147268423-d643c63a-96d2-40d1-9791-6cd842dc5647.png" alt="ebiten-bunny-mark" title="ebiten-bunny-mark" /> Ebiten Bunny Mark

This is an implementation of the popular graphics benchmark written on [Ebiten](https://ebiten.org/).

The initial benchmark was created by [Ian Lobb (code)](http://blog.iainlobb.com/2010/11/display-list-vs-blitting-results.html)
and [Amanda Lobb (art)](http://amandalobb.com/)

### Contents

- [Preview](#preview)
- [Running](#running)
- [Instructions](#instructions)
- [Perfomance](#perfomance)
- [Contributing](#contributing)

### Preview

<img src="https://user-images.githubusercontent.com/19890545/147268942-4c939aee-1c30-42d8-b792-39021fd62568.gif">
<img src="https://user-images.githubusercontent.com/19890545/147268946-e6ff7293-9715-472c-aba1-5dd04690d79c.gif">

### Running

To run the benchmark from ready binary, visit the [latest releases](https://github.com/sedyh/ebiten-bunny-mark/releases) page.

To run the benchmark from sources do the following command:

```
go run github.com/sedyh/ebiten-bunny-mark@master
```
<sub>Please remember that @master only works since Go 17.</sub>

### Instructions

- Close all other programs for more accurate results.
- Press the left mouse button to add some amount of gophers.
- Adjust the number of gophers that appear at a time with the mouse wheel.
- Increase the number of gophers until the FPS starts dropping below 60 to find out your result.
- To understand that the drop in performance is not a one-off - use the graphs on the right, they show TPS, FPS and the number of objects over a certain time.
- Press the right mouse button to color the gophers, this will greatly increase the load, but keep in mind that all measurements were taken without coloring.

### Perfomance

| Software                                     | Hardware             | Ebiten Version | Resolution | Maximum number of objects at stable 60 fps |
|----------------------------------------------|----------------------|----------------|------------|--------------------------------------------|
| Native, MacOS Big Sur 11.1                   | M1 2020              | 2.2.3          | 800x600    | 65400                                      |
| Native, Windows 10 Pro 19.0                  | i7 3770, GTX 1050 TI | 2.2.3          | 800x600    | 39000                                      |
| Native, Linux Mint 20.2 Cinnamon             | R5 4600, RX 5700 XT  | 2.2.3          | 800x600    | 36000                                      |
| Wasm, Linux Mint 20.2 Cinnamon, Chrome 96.0  | R5 4600, RX 5700 XT  | 2.2.3          | 800x600    | 8150                                       |
| Wasm, MacOS Big Sur 11.1, Safari 14.0        | M1 2020              | 2.2.3          | 800x600    | 8100                                       |
| Wasm, Linux Mint 20.2 Cinnamon, Firefox 95.0 | R5 4600, RX 5700 XT  | 2.2.3          | 800x600    | 6200                                       |
| Wasm, Windows 10 Pro 19.0, Chrome 96.0       | i7 3770, GTX 1050 TI | 2.2.3          | 800x600    | 5700                                       |
| Wasm, Windows 10 Pro 19.0, Firefox 95.0      | i7 3770, GTX 1050 TI | 2.2.3          | 800x600    | 4700                                       |
| Wasm, iPadOS 14.6, Safari 605.1              | A10X                 | 2.2.3          | 800x600    | 3200                                       |

### Contributing

To add, remove, or change things on the perfomance table: send a message in [this](https://github.com/sedyh/ebiten-bunny-mark/issues/2) issue.

To add new functionality or fix something: submit a pull request.
