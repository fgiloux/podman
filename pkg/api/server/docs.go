// Package api Provides an API for the  Libpod library
//
// This documentation describes the Podman v2.0 RESTful API.
// It replaces the Podman v1.0 API and was initially delivered
// along with Podman v2.0.  It consists of a Docker-compatible
// API and a Libpod API providing support for Podman’s unique
// features such as pods.
//
// To start the service and keep it running for 5,000 seconds (-t 0 runs forever):
//
//    podman system service -t 5000 &
//
// You can then use cURL on the socket using requests documented below.
//
// NOTE: if you install the package podman-docker, it will create a symbolic
// link for /run/docker.sock to /run/podman/podman.sock
//
// See podman-service(1) for more information.
//
//  Quick Examples:
//
//   'podman info'
//
//      curl --unix-socket /run/podman/podman.sock http://d/v3.0.0/libpod/info
//
//   'podman pull quay.io/containers/podman'
//
//      curl -XPOST --unix-socket /run/podman/podman.sock -v 'http://d/v3.0.0/images/create?fromImage=quay.io%2Fcontainers%2Fpodman'
//
//   'podman list images'
//
//      curl --unix-socket /run/podman/podman.sock -v 'http://d/v3.0.0/libpod/images/json' | jq
//
// Terms Of Service:
//
//     Schemes: http, https
//     Host: podman.io
//     BasePath: /
//     Version: 3.2.0
//     License: Apache-2.0 https://opensource.org/licenses/Apache-2.0
//     Contact: Podman <podman@lists.podman.io> https://podman.io/community/
//
//     InfoExtensions:
//     x-logo:
//       - url: https://raw.githubusercontent.com/containers/libpod/main/logo/podman-logo.png
//       - altText: "Podman logo"
//
//     Produces:
//     - application/json
//     - application/octet-stream
//     - text/plain
//
//     Consumes:
//     - application/json
//     - application/x-tar
// swagger:meta
package server
