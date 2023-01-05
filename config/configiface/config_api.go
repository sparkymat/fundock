package configiface

type ConfigAPI interface {
	DBConnectionString() string
	SingleUser() bool
	JWTSecret() string
	SessionSecret() string
}
