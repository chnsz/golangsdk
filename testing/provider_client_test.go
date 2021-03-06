package testing

import (
	"context"
	"fmt"
	"io/ioutil"
	"math"
	"net"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/chnsz/golangsdk"
	th "github.com/chnsz/golangsdk/testhelper"
	"github.com/chnsz/golangsdk/testhelper/client"
)

func TestAuthenticatedHeaders(t *testing.T) {
	p := &golangsdk.ProviderClient{
		TokenID: "1234",
	}
	expected := map[string]string{"X-Auth-Token": "1234"}
	actual := p.AuthenticatedHeaders()
	th.CheckDeepEquals(t, expected, actual)
}

func TestUserAgent(t *testing.T) {
	p := &golangsdk.ProviderClient{}

	p.UserAgent.Prepend("custom-user-agent/2.4.0")
	expected := "custom-user-agent/2.4.0 golangsdk/2.0.0"
	actual := p.UserAgent.Join()
	th.CheckEquals(t, expected, actual)

	p.UserAgent.Prepend("another-custom-user-agent/0.3.0", "a-third-ua/5.9.0")
	expected = "another-custom-user-agent/0.3.0 a-third-ua/5.9.0 custom-user-agent/2.4.0 golangsdk/2.0.0"
	actual = p.UserAgent.Join()
	th.CheckEquals(t, expected, actual)

	p.UserAgent = golangsdk.UserAgent{}
	expected = "golangsdk/2.0.0"
	actual = p.UserAgent.Join()
	th.CheckEquals(t, expected, actual)
}

func TestConcurrentReauth(t *testing.T) {
	var info = struct {
		numreauths int
		mut        *sync.RWMutex
	}{
		0,
		new(sync.RWMutex),
	}

	numconc := 20

	prereauthTok := client.TokenID
	postreauthTok := "12345678"

	p := new(golangsdk.ProviderClient)
	p.UseTokenLock()
	p.SetToken(prereauthTok)
	p.ReauthFunc = func() error {
		time.Sleep(1 * time.Second)
		p.AuthenticatedHeaders()
		info.mut.Lock()
		info.numreauths++
		info.mut.Unlock()
		p.TokenID = postreauthTok
		return nil
	}

	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/route", func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("X-Auth-Token") != postreauthTok {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		info.mut.RLock()
		hasReauthed := info.numreauths != 0
		info.mut.RUnlock()

		if hasReauthed {
			th.CheckEquals(t, p.Token(), postreauthTok)
		}

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, `{}`)
	})

	wg := new(sync.WaitGroup)
	reqopts := new(golangsdk.RequestOpts)
	reqopts.KeepResponseBody = true
	reqopts.MoreHeaders = map[string]string{
		"X-Auth-Token": prereauthTok,
	}

	for i := 0; i < numconc; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			resp, err := p.Request("GET", fmt.Sprintf("%s/route", th.Endpoint()), reqopts)
			th.CheckNoErr(t, err)
			if resp == nil {
				t.Errorf("got a nil response")
				return
			}
			if resp.Body == nil {
				t.Errorf("response body was nil")
				return
			}
			defer resp.Body.Close()
			actual, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				t.Errorf("error reading response body: %s", err)
				return
			}
			th.CheckByteArrayEquals(t, []byte(`{}`), actual)
		}()
	}

	wg.Wait()

	th.AssertEquals(t, 1, info.numreauths)
}

func TestRequestWithContext(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "OK")
	}))
	defer ts.Close()

	ctx, cancel := context.WithCancel(context.Background())
	p := &golangsdk.ProviderClient{Context: ctx}

	res, err := p.Request("GET", ts.URL, &golangsdk.RequestOpts{KeepResponseBody: true})
	th.AssertNoErr(t, err)
	_, err = ioutil.ReadAll(res.Body)
	th.AssertNoErr(t, err)
	err = res.Body.Close()
	th.AssertNoErr(t, err)

	cancel()
	res, err = p.Request("GET", ts.URL, &golangsdk.RequestOpts{})
	if err == nil {
		t.Fatal("expecting error, got nil")
	}
	if !strings.Contains(err.Error(), ctx.Err().Error()) {
		t.Fatalf("expecting error to contain: %q, got %q", ctx.Err().Error(), err.Error())
	}
}

func TestRequestConnectionReuse(t *testing.T) {
	ts := httptest.NewUnstartedServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "OK")
	}))

	// an amount of iterations
	var iter = 10000
	// connections tracks an amount of connections made
	var connections int64

	ts.Config.ConnState = func(_ net.Conn, s http.ConnState) {
		// track an amount of connections
		if s == http.StateNew {
			atomic.AddInt64(&connections, 1)
		}
	}
	ts.Start()
	defer ts.Close()

	p := &golangsdk.ProviderClient{}
	reqopts := new(golangsdk.RequestOpts)

	for i := 0; i < iter; i++ {
		_, err := p.Request("GET", ts.URL, reqopts)
		th.AssertNoErr(t, err)
	}

	th.AssertEquals(t, int64(1), connections)
}

func TestRequestConnectionClose(t *testing.T) {
	ts := httptest.NewUnstartedServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "OK")
	}))

	// an amount of iterations
	var iter = 10
	// connections tracks an amount of connections made
	var connections int64

	ts.Config.ConnState = func(_ net.Conn, s http.ConnState) {
		// track an amount of connections
		if s == http.StateNew {
			atomic.AddInt64(&connections, 1)
		}
	}
	ts.Start()
	defer ts.Close()

	p := &golangsdk.ProviderClient{}
	reqopts := new(golangsdk.RequestOpts)
	reqopts.KeepResponseBody = true

	for i := 0; i < iter; i++ {
		_, err := p.Request("GET", ts.URL, reqopts)
		th.AssertNoErr(t, err)
	}

	th.AssertEquals(t, int64(iter), connections)
}

func retryTest(retryCounter *uint, t *testing.T) golangsdk.RetryFunc {
	return func(ctx context.Context, respErr *golangsdk.ErrUnexpectedResponseCode, e error, retries uint) error {
		seconds := math.Pow(2, float64(retries))
		if seconds > 60 { // won't wait more than 60 seconds
			seconds = 60
		}

		sleep := time.Duration(seconds) * time.Second

		if ctx != nil {
			t.Logf("Context sleeping for %d seconds, retry number %d", int(seconds), retries)
			select {
			case <-time.After(sleep):
				t.Log("sleep is over")
			case <-ctx.Done():
				t.Log(ctx.Err())
				return e
			}
		} else {
			t.Logf("Sleeping for %d seconds, retry number %d", int(seconds), retries)
			time.Sleep(sleep)
			t.Log("sleep is over")
		}

		*retryCounter = *retryCounter + 1

		return nil
	}
}

func TestRequestRetry(t *testing.T) {
	var retryCounter uint

	p := &golangsdk.ProviderClient{}
	p.UseTokenLock()
	p.SetToken(client.TokenID)
	p.MaxBackoffRetries = 3

	p.RetryBackoffFunc = retryTest(&retryCounter, t)

	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/route", func(w http.ResponseWriter, r *http.Request) {
		//always reply 429
		http.Error(w, "retry later", http.StatusTooManyRequests)
	})

	_, err := p.Request("GET", th.Endpoint()+"/route", &golangsdk.RequestOpts{})
	if err == nil {
		t.Fatal("expecting error, got nil")
	}
	t.Logf("error message: %s", err)
	th.AssertEquals(t, retryCounter, p.MaxBackoffRetries)
}

func TestRequestRetrySuccess(t *testing.T) {
	var retryCounter uint

	p := &golangsdk.ProviderClient{}
	p.UseTokenLock()
	p.SetToken(client.TokenID)
	p.MaxBackoffRetries = 3

	p.RetryBackoffFunc = retryTest(&retryCounter, t)

	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/route", func(w http.ResponseWriter, r *http.Request) {
		//always reply 200
		http.Error(w, "retry later", http.StatusOK)
	})

	_, err := p.Request("GET", th.Endpoint()+"/route", &golangsdk.RequestOpts{})
	if err != nil {
		t.Fatal(err)
	}
	th.AssertEquals(t, retryCounter, uint(0))
}

func TestRequestRetryContext(t *testing.T) {
	var retryCounter uint

	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		sleep := 5 * time.Second
		time.Sleep(sleep)
		cancel()
	}()

	p := &golangsdk.ProviderClient{
		Context: ctx,
	}
	p.UseTokenLock()
	p.SetToken(client.TokenID)
	p.MaxBackoffRetries = 3

	p.RetryBackoffFunc = retryTest(&retryCounter, t)

	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/route", func(w http.ResponseWriter, r *http.Request) {
		//always reply 429
		http.Error(w, "retry later", http.StatusTooManyRequests)
	})

	_, err := p.Request("GET", th.Endpoint()+"/route", &golangsdk.RequestOpts{})
	if err == nil {
		t.Fatal("expecting error, got nil")
	}
	t.Logf("retryCounter: %d, p.MaxBackoffRetries: %d", retryCounter, p.MaxBackoffRetries)
	th.AssertEquals(t, retryCounter, p.MaxBackoffRetries-1)
}
