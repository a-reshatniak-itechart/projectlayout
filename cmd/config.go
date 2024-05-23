package main

type Config struct {
	ListenAddr string

	UserRepoType   string
	UserLocalFsDir string
	UserSqliteDsn  string
}
