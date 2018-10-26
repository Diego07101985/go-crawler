package api-tests

type ReposTestClient struct {
	Repos []repos.Repo
	Err   error
}

func (c ReposTestClient) Get(string) ([]repos.Repo, error) {
	return c.Repos, c.Err
}

type Client interface {
	Get(string) ([]Repo, error)
}

type ReposClient struct{}

func (c ReposClient) Get(user string) ([]Repo, error) {
	// call github api
}

type App struct {
	repos repos.Client
}