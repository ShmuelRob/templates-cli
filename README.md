project-generator/
├── cmd/
│   └── gen/
│       └── main.go
├── internal/
│   ├── cli/
│   │   └── cli.go
│   ├── templates/
│   │   ├── template.go
│   │   └── etl.go
│   └── utils/
│       └── utils.go
├── templates/
│   └── etl/
│       ├── extract/
│       │   └── extract.go.tmpl
│       ├── transform/
│       │   └── transform.go.tmpl
│       ├── load/
│       │   └── load.go.tmpl
│       └── main.go.tmpl
├── go.mod
└── README.md