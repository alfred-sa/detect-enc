Pour compiler : 
1) Installer Go : https://go.dev/doc/install
2) Installer les modules : `go mod download`
3) Exécuter `go build -o detect-enc detect-enc.go`

Pour compiler pour Windows : `GOOS=windows GOARCH=amd64 go build -o detect-enc.exe detect-enc.go`
Pour compiler pour Linux : `GOOS=linux GOARCH=amd64 go build -o detect-enc.exe detect-enc.go`

Usage : `./detect-enc <fichier>`
