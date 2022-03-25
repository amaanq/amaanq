package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

const (
	Username              = "amaanq"
	TokenPath             = ".accesstoken"
	GitHubURL             = "https://api.github.com/users/" + Username
	GitHubPublicReposURL  = "https://api.github.com/users/" + Username + "/repos"
	GitHubPrivateReposURL = "https://api.github.com/user/repos?access_token="
)

var (
	Token string
)

func main() {
	LoadToken()
	GetGitHubUserStats()
	GetGitHubPublicUserRepoStats()
	GetGitHubPrivateUserRepoStats()
}

func LoadToken() {
	tokenBytes, err := ioutil.ReadFile(TokenPath)
	if err != nil {
		panic(err)
	}

	Token = string(tokenBytes)
}

func GetGitHubUserStats() {
	req, err := http.NewRequest("GET", GitHubURL, nil)
	if err != nil {
		panic(err)
	}
	req.SetBasicAuth(Username, Token)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	ghresp := &GitHubUsersResponse{}
	if err := json.Unmarshal(b, ghresp); err != nil {
		panic(err)
	}
	fmt.Println(ghresp)
}

func GetGitHubPublicUserRepoStats() {
	req, err := http.NewRequest("GET", GitHubPublicReposURL, nil)
	if err != nil {
		panic(err)
	}
	req.SetBasicAuth(Username, Token)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	ghreporesp := make(GitHubPublicRepoResponse, 0)
	if err := json.Unmarshal(b, &ghreporesp); err != nil {
		panic(err)
	}
}

func GetGitHubPrivateUserRepoStats() {
	req, err := http.NewRequest("GET", GitHubPrivateReposURL+Token, nil)
	if err != nil {
		panic(err)
	}
	req.SetBasicAuth(Username, Token)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	ghreporesp := make(GitHubPrivateRepoResponse, 0)
	if err := json.Unmarshal(b, &ghreporesp); err != nil {
		panic(err)
	}
	for _, repo := range ghreporesp {
		if repo.Language != nil {
			fmt.Println(*repo.Language)
		}
	}
}

type GitHubUsersResponse struct {
	Login                   string `json:"login"`
	ID                      int    `json:"id"`
	NodeID                  string `json:"node_id"`
	AvatarURL               string `json:"avatar_url"`
	GravatarID              string `json:"gravatar_id"`
	URL                     string `json:"url"`
	HTMLURL                 string `json:"html_url"`
	FollowersURL            string `json:"followers_url"`
	FollowingURL            string `json:"following_url"`
	GistsURL                string `json:"gists_url"`
	StarredURL              string `json:"starred_url"`
	SubscriptionsURL        string `json:"subscriptions_url"`
	OrganizationsURL        string `json:"organizations_url"`
	ReposURL                string `json:"repos_url"`
	EventsURL               string `json:"events_url"`
	ReceivedEventsURL       string `json:"received_events_url"`
	Type                    string `json:"type"`
	SiteAdmin               bool   `json:"site_admin"`
	Name                    string `json:"name"`
	Company                 any    `json:"company"`
	Blog                    string `json:"blog"`
	Location                any    `json:"location"`
	Email                   any    `json:"email"`
	Hireable                any    `json:"hireable"`
	Bio                     string `json:"bio"`
	TwitterUsername         any    `json:"twitter_username"`
	PublicRepos             int    `json:"public_repos"`
	PublicGists             int    `json:"public_gists"`
	Followers               int    `json:"followers"`
	Following               int    `json:"following"`
	CreatedAt               string `json:"created_at"`
	UpdatedAt               string `json:"updated_at"`
	PrivateGists            int    `json:"private_gists"`
	TotalPrivateRepos       int    `json:"total_private_repos"`
	OwnedPrivateRepos       int    `json:"owned_private_repos"`
	DiskUsage               int    `json:"disk_usage"`
	Collaborators           int    `json:"collaborators"`
	TwoFactorAuthentication bool   `json:"two_factor_authentication"`
	Plan                    Plan   `json:"plan"`
}

type Plan struct {
	Name          string `json:"name"`
	Space         int    `json:"space"`
	Collaborators int    `json:"collaborators"`
	PrivateRepos  int    `json:"private_repos"`
}

type GitHubPublicRepoResponse []GitHubPublicRepo

type GitHubPublicRepo struct {
	ID               int           `json:"id"`
	NodeID           string        `json:"node_id"`
	Name             string        `json:"name"`
	FullName         string        `json:"full_name"`
	Private          bool          `json:"private"`
	Owner            Owner         `json:"owner"`
	HTMLURL          string        `json:"html_url"`
	Description      *string       `json:"description"`
	Fork             bool          `json:"fork"`
	URL              string        `json:"url"`
	ForksURL         string        `json:"forks_url"`
	KeysURL          string        `json:"keys_url"`
	CollaboratorsURL string        `json:"collaborators_url"`
	TeamsURL         string        `json:"teams_url"`
	HooksURL         string        `json:"hooks_url"`
	IssueEventsURL   string        `json:"issue_events_url"`
	EventsURL        string        `json:"events_url"`
	AssigneesURL     string        `json:"assignees_url"`
	BranchesURL      string        `json:"branches_url"`
	TagsURL          string        `json:"tags_url"`
	BlobsURL         string        `json:"blobs_url"`
	GitTagsURL       string        `json:"git_tags_url"`
	GitRefsURL       string        `json:"git_refs_url"`
	TreesURL         string        `json:"trees_url"`
	StatusesURL      string        `json:"statuses_url"`
	LanguagesURL     string        `json:"languages_url"`
	StargazersURL    string        `json:"stargazers_url"`
	ContributorsURL  string        `json:"contributors_url"`
	SubscribersURL   string        `json:"subscribers_url"`
	SubscriptionURL  string        `json:"subscription_url"`
	CommitsURL       string        `json:"commits_url"`
	GitCommitsURL    string        `json:"git_commits_url"`
	CommentsURL      string        `json:"comments_url"`
	IssueCommentURL  string        `json:"issue_comment_url"`
	ContentsURL      string        `json:"contents_url"`
	CompareURL       string        `json:"compare_url"`
	MergesURL        string        `json:"merges_url"`
	ArchiveURL       string        `json:"archive_url"`
	DownloadsURL     string        `json:"downloads_url"`
	IssuesURL        string        `json:"issues_url"`
	PullsURL         string        `json:"pulls_url"`
	MilestonesURL    string        `json:"milestones_url"`
	NotificationsURL string        `json:"notifications_url"`
	LabelsURL        string        `json:"labels_url"`
	ReleasesURL      string        `json:"releases_url"`
	DeploymentsURL   string        `json:"deployments_url"`
	CreatedAt        string        `json:"created_at"`
	UpdatedAt        string        `json:"updated_at"`
	PushedAt         string        `json:"pushed_at"`
	GitURL           string        `json:"git_url"`
	SSHURL           string        `json:"ssh_url"`
	CloneURL         string        `json:"clone_url"`
	SvnURL           string        `json:"svn_url"`
	Homepage         any           `json:"homepage"`
	Size             int           `json:"size"`
	StargazersCount  int           `json:"stargazers_count"`
	WatchersCount    int           `json:"watchers_count"`
	Language         *string       `json:"language"`
	HasIssues        bool          `json:"has_issues"`
	HasProjects      bool          `json:"has_projects"`
	HasDownloads     bool          `json:"has_downloads"`
	HasWiki          bool          `json:"has_wiki"`
	HasPages         bool          `json:"has_pages"`
	ForksCount       int           `json:"forks_count"`
	MirrorURL        any           `json:"mirror_url"`
	Archived         bool          `json:"archived"`
	Disabled         bool          `json:"disabled"`
	OpenIssuesCount  int           `json:"open_issues_count"`
	License          *License      `json:"license"`
	AllowForking     bool          `json:"allow_forking"`
	IsTemplate       bool          `json:"is_template"`
	Topics           []any         `json:"topics"`
	Visibility       Visibility    `json:"visibility"`
	Forks            int           `json:"forks"`
	OpenIssues       int           `json:"open_issues"`
	Watchers         int           `json:"watchers"`
	DefaultBranch    DefaultBranch `json:"default_branch"`
	Permissions      Permissions   `json:"permissions"`
}

type License struct {
	Key    string `json:"key"`
	Name   string `json:"name"`
	SpdxID string `json:"spdx_id"`
	URL    string `json:"url"`
	NodeID string `json:"node_id"`
}

type Owner struct {
	Login             Login        `json:"login"`
	ID                int          `json:"id"`
	NodeID            NodeID       `json:"node_id"`
	AvatarURL         string       `json:"avatar_url"`
	GravatarID        string       `json:"gravatar_id"`
	URL               string       `json:"url"`
	HTMLURL           string       `json:"html_url"`
	FollowersURL      string       `json:"followers_url"`
	FollowingURL      FollowingURL `json:"following_url"`
	GistsURL          GistsURL     `json:"gists_url"`
	StarredURL        StarredURL   `json:"starred_url"`
	SubscriptionsURL  string       `json:"subscriptions_url"`
	OrganizationsURL  string       `json:"organizations_url"`
	ReposURL          string       `json:"repos_url"`
	EventsURL         EventsURL    `json:"events_url"`
	ReceivedEventsURL string       `json:"received_events_url"`
	Type              Type         `json:"type"`
	SiteAdmin         bool         `json:"site_admin"`
}

type Permissions struct {
	Admin    bool `json:"admin"`
	Maintain bool `json:"maintain"`
	Push     bool `json:"push"`
	Triage   bool `json:"triage"`
	Pull     bool `json:"pull"`
}

type DefaultBranch string

type EventsURL string

type FollowingURL string

type GistsURL string

type Login string

type NodeID string

type StarredURL string

type Type string

type Visibility string

type GitHubPrivateRepoResponse []GitHubPrivateRepo

type GitHubPrivateRepo struct {
	ID               int           `json:"id"`
	NodeID           string        `json:"node_id"`
	Name             string        `json:"name"`
	FullName         string        `json:"full_name"`
	Private          bool          `json:"private"`
	Owner            Owner         `json:"owner"`
	HTMLURL          string        `json:"html_url"`
	Description      *string       `json:"description"`
	Fork             bool          `json:"fork"`
	URL              string        `json:"url"`
	ForksURL         string        `json:"forks_url"`
	KeysURL          string        `json:"keys_url"`
	CollaboratorsURL string        `json:"collaborators_url"`
	TeamsURL         string        `json:"teams_url"`
	HooksURL         string        `json:"hooks_url"`
	IssueEventsURL   string        `json:"issue_events_url"`
	EventsURL        string        `json:"events_url"`
	AssigneesURL     string        `json:"assignees_url"`
	BranchesURL      string        `json:"branches_url"`
	TagsURL          string        `json:"tags_url"`
	BlobsURL         string        `json:"blobs_url"`
	GitTagsURL       string        `json:"git_tags_url"`
	GitRefsURL       string        `json:"git_refs_url"`
	TreesURL         string        `json:"trees_url"`
	StatusesURL      string        `json:"statuses_url"`
	LanguagesURL     string        `json:"languages_url"`
	StargazersURL    string        `json:"stargazers_url"`
	ContributorsURL  string        `json:"contributors_url"`
	SubscribersURL   string        `json:"subscribers_url"`
	SubscriptionURL  string        `json:"subscription_url"`
	CommitsURL       string        `json:"commits_url"`
	GitCommitsURL    string        `json:"git_commits_url"`
	CommentsURL      string        `json:"comments_url"`
	IssueCommentURL  string        `json:"issue_comment_url"`
	ContentsURL      string        `json:"contents_url"`
	CompareURL       string        `json:"compare_url"`
	MergesURL        string        `json:"merges_url"`
	ArchiveURL       string        `json:"archive_url"`
	DownloadsURL     string        `json:"downloads_url"`
	IssuesURL        string        `json:"issues_url"`
	PullsURL         string        `json:"pulls_url"`
	MilestonesURL    string        `json:"milestones_url"`
	NotificationsURL string        `json:"notifications_url"`
	LabelsURL        string        `json:"labels_url"`
	ReleasesURL      string        `json:"releases_url"`
	DeploymentsURL   string        `json:"deployments_url"`
	CreatedAt        string        `json:"created_at"`
	UpdatedAt        string        `json:"updated_at"`
	PushedAt         string        `json:"pushed_at"`
	GitURL           string        `json:"git_url"`
	SSHURL           string        `json:"ssh_url"`
	CloneURL         string        `json:"clone_url"`
	SvnURL           string        `json:"svn_url"`
	Homepage         *string       `json:"homepage"`
	Size             int           `json:"size"`
	StargazersCount  int           `json:"stargazers_count"`
	WatchersCount    int           `json:"watchers_count"`
	Language         *string       `json:"language"`
	HasIssues        bool          `json:"has_issues"`
	HasProjects      bool          `json:"has_projects"`
	HasDownloads     bool          `json:"has_downloads"`
	HasWiki          bool          `json:"has_wiki"`
	HasPages         bool          `json:"has_pages"`
	ForksCount       int           `json:"forks_count"`
	MirrorURL        any           `json:"mirror_url"`
	Archived         bool          `json:"archived"`
	Disabled         bool          `json:"disabled"`
	OpenIssuesCount  int           `json:"open_issues_count"`
	License          *License      `json:"license"`
	AllowForking     bool          `json:"allow_forking"`
	IsTemplate       bool          `json:"is_template"`
	Topics           []any         `json:"topics"`
	Visibility       Visibility    `json:"visibility"`
	Forks            int           `json:"forks"`
	OpenIssues       int           `json:"open_issues"`
	Watchers         int           `json:"watchers"`
	DefaultBranch    DefaultBranch `json:"default_branch"`
	Permissions      Permissions   `json:"permissions"`
}
