package goldshire

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

var (
	_      Store = &GitlabStore{}
	client       = http.DefaultClient
)

type GitlabProjectInfo struct {
	Path    string `json:"path_with_namespace"`
	SSHUrl  string `json:"ssh_url_to_repo"`
	HTTPUrl string `json:"http_url_to_repo"`
}

type GitlabError struct {
	Message string `json:"message"`
}

func (this *GitlabError) Error() string {
	return this.Message
}

func NewGitlabStore(cfg GitlabConfig) *GitlabStore {
	return &GitlabStore{
		cfg: cfg,
	}
}

type GitlabStore struct {
	cfg GitlabConfig
}

func (this *GitlabStore) Get(path string) (*Meta, error) {
	if path == "" {
		return nil, nil
	}

	pieces := strings.SplitN(path, "/", 3)
	if len(pieces) < 2 { // with namespace, without project name
		return nil, nil
	}

	urlStr := fmt.Sprintf("%s/%s", this.cfg.ApiBase, pieces[0]+"%2F"+pieces[1])
	req, err := http.NewRequest("GET", urlStr, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("PRIVATE-TOKEN", this.cfg.PrivateToken)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	// project not found
	if resp.StatusCode == http.StatusNotFound {
		return nil, nil
	}

	// decode error message or project info
	decoder := json.NewDecoder(resp.Body)

	if resp.StatusCode/100 != 2 {
		msg := &GitlabError{}
		if err := decoder.Decode(msg); err != nil {
			return nil, err
		}

		return nil, msg
	}

	info := &GitlabProjectInfo{}

	if err := decoder.Decode(info); err != nil {
		return nil, err
	}

	return &Meta{
		Base: info.Path,
		VCS:  "git",
		Url:  info.SSHUrl,
	}, nil
}
