module github.com/chnsz/golangsdk

go 1.14

require (
	github.com/stretchr/testify v1.10.0
	github.com/tjfoc/gmsm v1.4.1
	go.mongodb.org/mongo-driver v1.17.2
	golang.org/x/crypto v0.26.0
	golang.org/x/net v0.25.0
	gopkg.in/yaml.v2 v2.3.0
)

replace golang.org/x/crypto => golang.org/x/crypto v0.23.0
