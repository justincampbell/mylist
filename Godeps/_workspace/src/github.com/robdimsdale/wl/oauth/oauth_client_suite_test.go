package oauth_test

import (
	"github.com/justincampbell/mylist/Godeps/_workspace/src/github.com/robdimsdale/wl"
	"github.com/justincampbell/mylist/Godeps/_workspace/src/github.com/robdimsdale/wl/logger"
	"github.com/justincampbell/mylist/Godeps/_workspace/src/github.com/robdimsdale/wl/oauth"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/ghttp"
	"testing"
)

const (
	dummyAccessToken = "dummyAccessToken"
	dummyClientID    = "dummyClientID"
)

var (
	client wl.Client

	server *ghttp.Server
	apiURL string

	testLogger logger.Logger
)

func TestWL(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "WL Suite")
}

var _ = BeforeEach(func() {
	server = ghttp.NewServer()
	apiURL = server.URL()

	testLogger = logger.NewTestLogger(GinkgoWriter)
	client = oauth.NewClient(
		dummyAccessToken,
		dummyClientID,
		apiURL,
		testLogger,
	)
})

var _ = AfterEach(func() {
	server.Close()
})
