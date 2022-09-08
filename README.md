# golang-examples
This is a tutorial project that I use to practice some Golang basics.

## Notes
### Running program as an executable
1. The instructions were taken from https://go.dev/doc/tutorial/compile-install
2. '$ echo $PATH' gives the path to the Shell path of the machine. This is the same as the how we 'add to the path' with Windows.
3. 'Go build' compiles the code into an executable in the current path 
4. To run this from the current directory, run './(name of the main file)' e.g. '$ ./hello'
5. To run this from anywhere on the computer, we need to add the Go install path to the Shell path of the machine.
6. To check the Go install path use '$ go list -f '{{.Target}}'
7. To add this to the Shell path, use '$ export PATH=$PATH:/path/to/your/install/directory', Note that this should be the directory path, and should not include the executable.
8. To verify that the Go install path has been added to the Shell path, run '$ echo $PATH' again. This should now have the Go install path added to it.
9. Running the program now from anywhere in the machine, should execute the executable. e.g. '$ hello'