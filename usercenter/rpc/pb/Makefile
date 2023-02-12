RED  =  "\e[31;1m"
GREEN = "\e[32;1m"
YELLOW = "\e[33;1m"
BLUE  = "\e[34;1m"
PURPLE = "\e[35;1m"
CYAN  = "\e[36;1m"

gen-rpc:
	goctl rpc protoc *.proto --go_out=../ --go-grpc_out=../ --zrpc_out=../ -style go_zero --multiple
	@printf $(GREEN)"[SUCCESS] generate rpc successfully"