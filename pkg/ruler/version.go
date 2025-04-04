package ruler

var (
	Githash string
	Version string
)

func Semver() string {
	return Version
}
