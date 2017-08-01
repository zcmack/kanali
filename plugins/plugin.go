// Copyright (c) 2017 Northwestern Mutual.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package plugins

import (
  "fmt"
	"context"
  "net/url"
	"net/http"

	"github.com/northwesternmutual/kanali/controller"
	"github.com/northwesternmutual/kanali/spec"
	"github.com/opentracing/opentracing-go"
)

// Plugin is an interface that is used for every Plugin used by Kanali.
// If external plugins are developed, they also must conform to this interface.
type Plugin interface {
	OnRequest(ctx context.Context, proxy spec.APIProxy, ctlr controller.Controller, req *http.Request, span opentracing.Span) error
	OnResponse(ctx context.Context, proxy spec.APIProxy, ctlr controller.Controller, req *http.Request, resp *http.Response, span opentracing.Span) error
}

type PluginList struct {
  Plugins []PluginItem `json:"plugins"`
}

type PluginItem struct {
  Location  string `json:"location"`
  Version   string `json:"version"`
  Name      string `json:"name"`
}

func (p PluginItem) GetURL() (*url.URL, error) {
  u, err := url.Parse(fmt.Sprintf(
    "https://%s/archive/%s.zip",
    p.Location,
    p.Version,
  ))
  if err != nil {
    return nil, err
  }
  return u, nil
}
