create-dir: manifests

manifests: a.go b.go
	mkdir manifests
	mv a.go manifests
	mv b.go manifests

a.go:
	echo "package main" > a.go
	echo "import \"fmt\"" >> a.go
	echo "func main() {" >> a.go
	echo "fmt.Println(\"Hello from a.go\")" >> a.go
	echo "}" >> a.go

b.go:
	echo "package main" > b.go
	echo "import \"fmt\"" >> b.go
	echo "func main() {" >> b.go
	echo "fmt.Println(\"Hello from b.go\")" >> b.go
	echo "}" >> b.go

# .PHONY is a special command that treats the target as a command and not a file
# Without .PHONY, create will be treated as a file and hence if we have a file called "create", running
# "make create" will not run the command in the target create

.PHONY: remove-all-files
remove:
	rm -rf manifests