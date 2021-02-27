# go-gwyddion-engine
This is a pet project game engine written in Golang, using Go OpenGL wrappers.  This is still very much in the early days, and the existing API will have breaking changes before stabilizing.

## Development
All of the eventual functionality of the game engine will work best when running on your local machine, however this does require some setup to ensure all of the necessary OpenGL and C based development libraries are installed.  To allow people to get up and running quickly, there is a VS Code Remote Container environment defined, which has all ofthe necessary libraries installed in a Docker container.

To gain access to the windows being created by the game engine while running in Remote Containers, visit http://localhost:6080, which will show the virtual desktop being forwarded by a web VNC interface.

## Integration Test
While this repo is intended as a library, there is a functional integration test program that can be ran by executing `make integration-test`.  This will launch a window and cycle through drawing squares to the screen.  As additional functionality gets added to the engine, the plan is to integrate the functionality with this test program, both as an example of implementation, but also to verify functionality.

The file `/cmd/integration/main.go` also provides a simple example application.

## Notes
* There is an error that is thrown when first running an OpenGL application against the remote display.  I think there is some state in the image used for VNC.  If you run into this, simply execute the OpenGL application again and it should work.
* The mouse click functionality of the VNC web interface does not seem to be working correctly, as both left and right clicks bring up a context menu, and neither allows for dragging.  This is something to look into for future functionality.
