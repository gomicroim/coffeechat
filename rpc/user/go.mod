module user

go 1.17

// fixed go mod tidy in go work
// see: https://github.com/golang/go/issues/50750
replace github.com/gomicroim/gomicroim/pkg => ../../pkg

require (
	entgo.io/ent v0.11.2
	github.com/go-kratos/kratos/contrib/registry/etcd/v2 v2.0.0-20220927060808-4d950507475c
	github.com/go-kratos/kratos/v2 v2.5.0
	github.com/go-redis/redis/v8 v8.11.5
	github.com/go-sql-driver/mysql v1.6.0
	github.com/gomicroim/gomicroim/pkg v0.0.0-00010101000000-000000000000
	github.com/google/wire v0.5.0
	github.com/stretchr/testify v1.8.0
	go.etcd.io/etcd/client/v3 v3.5.5
	go.uber.org/zap v1.23.0
	google.golang.org/grpc v1.46.2
	google.golang.org/protobuf v1.28.0
)

require (
	ariga.io/atlas v0.5.1-0.20220717122844-8593d7eb1a8e // indirect
	github.com/agext/levenshtein v1.2.1 // indirect
	github.com/apparentlymart/go-textseg/v13 v13.0.0 // indirect
	github.com/cespare/xxhash/v2 v2.1.2 // indirect
	github.com/coreos/go-semver v0.3.0 // indirect
	github.com/coreos/go-systemd/v22 v22.3.2 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/dgrijalva/jwt-go v3.2.0+incompatible // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/fsnotify/fsnotify v1.5.4 // indirect
	github.com/go-logr/logr v1.2.3 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/go-openapi/inflect v0.19.0 // indirect
	github.com/go-playground/form/v4 v4.2.0 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/google/go-cmp v0.5.8 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/gorilla/mux v1.8.0 // indirect
	github.com/hashicorp/hcl/v2 v2.10.0 // indirect
	github.com/imdario/mergo v0.3.12 // indirect
	github.com/mitchellh/go-wordwrap v0.0.0-20150314170334-ad45545899c7 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/twinj/uuid v1.0.0 // indirect
	github.com/zclconf/go-cty v1.8.0 // indirect
	go.etcd.io/etcd/api/v3 v3.5.5 // indirect
	go.etcd.io/etcd/client/pkg/v3 v3.5.5 // indirect
	go.opentelemetry.io/otel v1.8.0 // indirect
	go.opentelemetry.io/otel/trace v1.8.0 // indirect
	go.uber.org/atomic v1.9.0 // indirect
	go.uber.org/multierr v1.6.0 // indirect
	golang.org/x/mod v0.6.0-dev.0.20220419223038-86c51ed26bb4 // indirect
	golang.org/x/net v0.0.0-20220809184613-07c6da5e1ced // indirect
	golang.org/x/sync v0.0.0-20220722155255-886fb9371eb4 // indirect
	golang.org/x/sys v0.0.0-20220728004956-3c1f35247d10 // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/genproto v0.0.0-20220524023933-508584e28198 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
