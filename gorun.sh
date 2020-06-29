clear
# export SCRAPP=1
export ENV=development
export GIN_MODE=debug #debug #release
go run main.go -alsologtostderr -v=2 -log_dir="/c/Temp/logs"
