clear
export SCRAPP=0 # 1 if to scrapp worker

export IMPORT=1 # 1 if want to import from csv to mongo
export LANG=ar

export ENV=development
export GIN_MODE=debug #debug #release
go run main.go
