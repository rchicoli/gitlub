package gitcmd

import (
	"os/user"

	"github.com/libgit2/git2go"
)

var CloneOptions = &git.CloneOptions{
	FetchOptions: &FetchOptions,
	Bare:         false,
}

var remoteCallBacks = git.RemoteCallbacks{
	CredentialsCallback:      credentialsCallback,
	CertificateCheckCallback: certificateCheckCallback,
}

var FetchOptions = git.FetchOptions{
	RemoteCallbacks: remoteCallBacks,
}

type Commands struct {
}

func credentialsCallback(url string, username string, allowedTypes git.CredType) (git.ErrorCode, *git.Cred) {

	currentUser, _ := user.Current()
	ret, cred := git.NewCredSshKey("git", currentUser.HomeDir+"/.ssh/id_rsa.pub", currentUser.HomeDir+"/.ssh/id_rsa", "")
	return git.ErrorCode(ret), &cred
}

// Made this one just return 0 during troubleshooting...
func certificateCheckCallback(cert *git.Certificate, valid bool, hostname string) git.ErrorCode {
	// if hostname != "github.com" {
	// 	return git.ErrUser
	// }
	return 0
}
