package integration

import (
	"testing"

	"github.com/caddyserver/caddy/v2/caddytest"
)

func TestHttpOnly(t *testing.T) {

	// arrange
	caddytest.InitServer(t, ` 
  {
    http_port     9080
    https_port    9443
  }
  
  a.caddy.local:9080 {
    respond /version 200 {
      body "hello from a.caddy.local"
    }	
    }
  `, "caddyfile")

	// act and assert
	caddytest.AssertGetResponse(t, "http://a.caddy.local:9080/version", 200, "hello from a.caddy.local")
}

func TestRespond(t *testing.T) {

	// arrange
	caddytest.InitServer(t, ` 
  {
    http_port     9080
    https_port    9443
  }
  
  a.caddy.local:9443 {
    tls /caddy.local.crt /caddy.local.key {
    }
    respond /version 200 {
      body "hello from a.caddy.local"
    }	
  }
  `, "caddyfile")

	// act and assert
	caddytest.AssertGetResponse(t, "https://a.caddy.local:9443/version", 200, "hello from a.caddy.local")
}

func xTestRedirect(t *testing.T) {

	// arrange
	caddytest.InitServer(t, `
  {
    http_port     9080
    https_port    9443
  }
  
  b.caddy.local:9443 {
    tls /caddy.local.crt /caddy.local.key {
    }

    redir / https://b.caddy.local:9443/hello 301
    
    respond /hello 200 {
      body "hello from b.caddy.local"
    }	
    }
  `, "caddyfile")

	// act and assert
	caddytest.AssertRedirect(t, "https://b.caddy.local:9443/", "https://b.caddy.local:9443/hello", 301)

	// follow redirect
	caddytest.AssertGetResponse(t, "https://b.caddy.local:9443/", 200, "hello from b.caddy.local")
}

func xTest2Hosts(t *testing.T) {

	// arrange
	caddytest.InitServer(t, `
  {
    http_port     9080
    https_port    9443
  }
  
  a.caddy.local:9443 {
    tls /caddy.local.crt /caddy.local.key {
    }

    respond /hello 200 {
      body "hello from a.caddy.local"
    }	
  }

  b.caddy.local:9443 {
    tls /caddy.local.crt /caddy.local.key {
    }

    respond /hello 200 {
      body "hello from b.caddy.local"
    }	
    }
  `, "caddyfile")

	// act and assert
	caddytest.AssertGetResponse(t, "https://a.caddy.local:9443/hello", 200, "hello from a.caddy.local")
	caddytest.AssertGetResponse(t, "https://b.caddy.local:9443/hello", 200, "hello from b.caddy.local")
}

func xTest2HostsAndOneStaticIP(t *testing.T) {

	// arrange
	caddytest.InitServer(t, `
  {
    http_port     9080
    https_port    9443
  }
  
  a.caddy.local:9443, 127.0.0.1:9080 {
    tls /caddy.local.crt /caddy.local.key {
    }

    respond /hello 200 {
      body "hello from a.caddy.local"
    }	
  }

  b.caddy.local:9443 {
    tls /caddy.local.crt /caddy.local.key {
    }

    respond /hello 200 {
      body "hello from b.caddy.local"
    }	
    }
  `, "caddyfile")

	// act and assert
	caddytest.AssertGetResponse(t, "http://127.0.0.1:9080/hello", 200, "hello from a.caddy.local")
	caddytest.AssertGetResponse(t, "https://a.caddy.local:9443/hello", 200, "hello from a.caddy.local")
	caddytest.AssertGetResponse(t, "https://b.caddy.local:9443/hello", 200, "hello from b.caddy.local")
}
