.PHONE: generated
.FORCE:

generated: .FORCE | bin/generator
	./bin/generator spec.yaml .

bin/generator: .FORCE
	go build -o bin/generator generator.go

bin/active: .FORCE | generated
	go build -o bin/active active.go