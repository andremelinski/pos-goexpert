```
go install github.com/spf13/cobra-cli@latest
export PATH="~/go/bin:$PATH"
go mod init <MODNAME>
cobra-cli init
cobra-cli add <COMMAND_NAME>
go run main.go <COMMAND_NAME>
cobra-cli add create -p 'categoryCmd'
go run main.go category create
```
