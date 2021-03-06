/*

 Copyright 2019 The KubeSphere Authors.

 Licensed under the Apache License, Version 2.0 (the "License");
 you may not use this file except in compliance with the License.
 You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

 Unless required by applicable law or agreed to in writing, software
 distributed under the License is distributed on an "AS IS" BASIS,
 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 See the License for the specific language governing permissions and
 limitations under the License.

*/

package filter

import (
	"k8s.io/klog"
	"net"
	"strings"
	"time"

	"github.com/emicklei/go-restful"
)

func Logging(req *restful.Request, resp *restful.Response, chain *restful.FilterChain) {
	start := time.Now()
	chain.ProcessFilter(req, resp)
	klog.V(4).Infof("%s - \"%s %s %s\" %d %d %dms",
		getRequestIP(req),
		req.Request.Method,
		req.Request.RequestURI,
		req.Request.Proto,
		resp.StatusCode(),
		resp.ContentLength(),
		time.Since(start)/time.Millisecond,
	)
}

func getRequestIP(req *restful.Request) string {
	address := strings.Trim(req.Request.Header.Get("X-Real-Ip"), " ")
	if address != "" {
		return address
	}

	address = strings.Trim(req.Request.Header.Get("X-Forwarded-For"), " ")
	if address != "" {
		return address
	}

	address, _, err := net.SplitHostPort(req.Request.RemoteAddr)
	if err != nil {
		return req.Request.RemoteAddr
	}

	return address
}
