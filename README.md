Pour compiler : 
1) Installer Go : https://go.dev/doc/install
2) Exécuter `go build -o detect-enc detect-enc.go`

Pour compiler pour Windows depuis Linux : `GOOS=windows GOARCH=amd64 go build -o detect-enc.exe detect-enc.go`

Usage : `./detect-enc <fichier>`
