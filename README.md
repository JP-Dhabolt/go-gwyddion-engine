# go-gwyddion-engine
This is a pet project game engine written in Golang, using Go OpenGL wrappers.  This is still very much in the early days, and the existing API will have breaking changes before stabilizing.

## Development
All of the eventual functionality of the game engine will work best when running on your local machine, however this does require some setup to ensure all of the necessary OpenGL and C based development libraries are installed.  To allow people to get up and running quickly, there is a VS Code Remote Container environment defined, which has all ofthe necessary libraries installed in a Docker container.

To gain access to the windows being created by the game engine while running in Remote Containers, visit http://localhost:6080, which will show the virtual desktop being forwarded by a web VNC interface.

## Integration Test
While this repo is intended as a library, there is a functional integration test program that can be ran by executing `make integration-test`.  This will launch a window and cycle through drawing squares to the screen.  As additional functionality gets added to the engine, the plan is to integrate the functionality with this test program, both as an example of implementation, but also to verify functionality.

The file [cmd/integration/main.go](cmd/integration/main.go) also provides a simple example application.

## Game of Life Example
The file [cmd/life/main.go](cmd/life/main.go), along with the contents of the package `internal/life` provide a more complicated example application.  You can also run this with `make run-life`, which will start up the simulation.

In the simulation, the living cells start off with a Red color, and slowly shift towards a Green color the more generations they are alive.  This provides some visual feedback as to how stable the current simulation is.

You can also look at the [cmd/life/main.go](cmd/life/main.go) for command line flags you can pass in to modify some of the starting parameters.  As the engine adds functionality, I anticipate expanding the interactivity of the Game of Life.

## Attributions
The Game of Life example (and the origins of the engine itself) originated with a [great blog series from Kyle Banks](https://github.com/KyleBanks/conways-gol).

## Notes
* There is an error that is thrown when first running an OpenGL application against the remote display.  I think there is some state in the image used for VNC.  If you run into this, simply execute the OpenGL application again and it should work.
* The mouse click functionality of the VNC web interface does not seem to be working correctly, as both left and right clicks bring up a context menu, and neither allows for dragging.  This is something to look into for future functionality.
