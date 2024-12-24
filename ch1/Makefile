# each possible operation is called a target (e.g. fmt, vet, build)

# .DEFAULT_GOAL defines the default target when the Makefile is run when no target is specified
.DEFAULT_GOAL := build

# the .PHONY lines keeps make from getting confused
# if a directory or file in the project has the same name as one of the listed targets
.PHONY:fmt vet build

# the word before the colon (:) is the name of the target
fmt:
# ./... references the content of the current directory and all subdirectories
	go fmt ./...

# any words after the target are other targets that must be run before the specified target runs
vet: fmt
	go vet ./...

build: vet
	go build

# https://pkg.go.dev/cmd/go#hdr-Remove_object_files_and_cached_files
clean:
	go clean

