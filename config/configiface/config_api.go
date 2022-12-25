package configiface

type ConfigAPI interface {
	DBConnectionString() string
}
