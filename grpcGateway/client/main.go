package main

// "cert.pem",
// "c:/workspace/Mkcert/localhost+2.pem"
var certFile = "c:/workspace/Mkcert/localhost+2.pem"

func main() {
	doGrpcRunGet()
	doGrpcRunPost()
	doHttpRunGet()
	doHttpRunPost()
}
