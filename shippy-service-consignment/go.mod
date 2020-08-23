module github.com/mdesson/shippy/shippy-service-consignment

go 1.14

replace github.com/mdesson/shippy/shippy-service-consignment => ../shippy-service/consignment

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

require (
	github.com/golang/protobuf v1.4.2
	github.com/micro/micro v1.18.0 // indirect
	github.com/micro/micro/v2 v2.9.3 // indirect
	github.com/micro/micro/v3 v3.0.0-beta.0.20200823205926-34486b8ea66f // indirect
	github.com/micro/services v0.10.0 // indirect
	golang.org/x/net v0.0.0-20200813134508-3edf25e44fcc // indirect
	golang.org/x/sys v0.0.0-20200820212457-1fb795427249 // indirect
	golang.org/x/text v0.3.3 // indirect
	google.golang.org/genproto v0.0.0-20200815001618-f69a88009b70 // indirect
	google.golang.org/grpc v1.31.0
	google.golang.org/protobuf v1.25.0
)
