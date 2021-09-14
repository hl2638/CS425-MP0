NOTES:

- run logger like this:
go run ./MP0-go/logger.go [<log_file_name>]
// if you don't specify log file name, it takes time.now() as file name.

- run node like this:
python -u ./MP0-py/generator.py [<frequency>] | go run ./MP0-go/node.go [<node_name> <address> <port>]

Logger receives messages and prints them on stdout. It also logs msg delays and sizes to the log file.


TODO:

- deploy on vm

- analyze in json, draw graph (Python)

- write instructions

- write report