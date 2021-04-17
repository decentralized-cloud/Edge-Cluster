module github.com/decentralized-cloud/edge-cluster

go 1.16

require (
	github.com/asaskevich/govalidator v0.0.0-20210307081110-f21760c49a8d // indirect
	github.com/go-kit/kit v0.10.0
	github.com/go-ozzo/ozzo-validation v3.6.0+incompatible
	github.com/gofrs/flock v0.8.0
	github.com/golang/mock v1.5.0
	github.com/lucsky/cuid v1.0.2
	github.com/micro-business/go-core v0.6.1
	github.com/onsi/ginkgo v1.15.2
	github.com/onsi/gomega v1.11.0
	github.com/pkg/errors v0.9.1
	github.com/prometheus/client_golang v1.10.0
	github.com/robfig/cron/v3 v3.0.1
	github.com/savsgio/atreugo/v11 v11.6.3
	github.com/savsgio/go-logger v1.0.0
	github.com/spf13/cobra v1.1.3
	github.com/thoas/go-funk v0.8.0
	go.mongodb.org/mongo-driver v1.5.0
	go.uber.org/zap v1.16.0
	google.golang.org/grpc v1.36.1
	google.golang.org/protobuf v1.25.0
	gopkg.in/yaml.v2 v2.4.0
	helm.sh/helm/v3 v3.5.4
	k8s.io/api v0.20.5
	k8s.io/apimachinery v0.20.5
	k8s.io/client-go v0.20.5
)

replace (
	github.com/docker/distribution => github.com/docker/distribution v0.0.0-20191216044856-a8371794149d
	github.com/docker/docker => github.com/moby/moby v17.12.0-ce-rc1.0.20200618181300-9dc6525e6118+incompatible
)
