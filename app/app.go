package app

// import "github.com/monnand/goredis"
// import "github.com/garyburd/redigo/redis"

// App entry
type App interface {
	Name() string
}

// Register app here
func NewApp(name string) App {
	switch name {
	case "maker":
		return &Maker{AppName: name}
	case "lab":
		return &Lab{AppName: name}
	default:
		return nil
	}
}
