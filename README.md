# <img align="right" src="https://user-images.githubusercontent.com/19890545/147268423-d643c63a-96d2-40d1-9791-6cd842dc5647.png" alt="ebiten-bunny-mark" title="ebiten-bunny-mark" /> Ebiten Bunny Mark

This is an implementation of the popular graphics benchmark written on [Ebiten](https://ebiten.org/).

The initial benchmark was created by [Ian Lobb (code)](http://blog.iainlobb.com/2010/11/display-list-vs-blitting-results.html)
and [Amanda Lobb (art)](http://amandalobb.com/)

### Contents

- [Preview](#preview)
- [Usage](#usage)
- [Perfomance](#perfomance)

### Preview

<img src="https://user-images.githubusercontent.com/19890545/147268942-4c939aee-1c30-42d8-b792-39021fd62568.gif">
<img src="https://user-images.githubusercontent.com/19890545/147268946-e6ff7293-9715-472c-aba1-5dd04690d79c.gif">

### Usage

To run benchmark please do the following command.

```
go run github.com/sedyh/ebiten-bunny-mark@latest
```

Controls:

- Press the left mouse button to add some amount of gophers 
- Adjust the number of gophers that appear at a time with the mouse wheel
- Press the right mouse button to color the gophers, this will greatly increase the load
- Increase the number of gophers until the FPS starts dropping below 60 to find out your result.
- To understand that the drop in performance is not a one-off - use the graphs on the right, they show TPS, FPS and the number of objects over a certain time.


### Perfomance

| Software           | Hardware            | Ebiten Version | Resolution | Maximum number of objects at stable 60 fps |
|--------------------|---------------------|----------------|------------|--------------------------------------------|
| Native Linux       | R5 4600, RX 5700 XT | 2.2.3          | 800x600    | 33800                                      |
| Wasm Linux Chrome  | R5 4600, RX 5700 XT | 2.2.3          | 800x600    | 8150                                       |
| Wasm Linux Firefox | R5 4600, RX 5700 XT | 2.2.3          | 800x600    | 6200                                       |
