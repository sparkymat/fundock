// Code generated by qtc from "layout.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line view/layout.qtpl:1
package view

//line view/layout.qtpl:1
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line view/layout.qtpl:1
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line view/layout.qtpl:1
func StreamLayout(qw422016 *qt422016.Writer, title string, content string) {
//line view/layout.qtpl:1
	qw422016.N().S(`
  <!DOCTYPE html>
  <html>
    <head>
      <meta charset="utf-8">
      <meta name="viewport" content="width=device-width,initial-scale=1">
      <title>`)
//line view/layout.qtpl:7
	qw422016.E().S(title)
//line view/layout.qtpl:7
	qw422016.N().S(`</title>
      <link rel="stylesheet" type="text/css" href="/css/uikit.min.css">
      <link rel="stylesheet" type="text/css" href="/css/style.css">
    </head>
    <body>
      <nav class="uk-navbar-container" uk-navbar>
        <div class="uk-navbar-left">
            <a href="/" class="uk-navbar-item uk-logo uk-margin-small-left">fundock</a>
            <ul class="uk-navbar-nav">
                <li class="uk-parent"><a href="/functions">Functions</a></li>
                <li class="uk-parent"><a href="/api_tokens">API Tokens</a></li>
            </ul>
        </div>
      </nav>
      `)
//line view/layout.qtpl:21
	qw422016.N().S(content)
//line view/layout.qtpl:21
	qw422016.N().S(`
      <script src="/js/uikit.min.js"></script>
      <script src="/js/uikit-icons.min.js"></script>
      <script src="/js/index.js"></script>
    </body>
  </html>
`)
//line view/layout.qtpl:27
}

//line view/layout.qtpl:27
func WriteLayout(qq422016 qtio422016.Writer, title string, content string) {
//line view/layout.qtpl:27
	qw422016 := qt422016.AcquireWriter(qq422016)
//line view/layout.qtpl:27
	StreamLayout(qw422016, title, content)
//line view/layout.qtpl:27
	qt422016.ReleaseWriter(qw422016)
//line view/layout.qtpl:27
}

//line view/layout.qtpl:27
func Layout(title string, content string) string {
//line view/layout.qtpl:27
	qb422016 := qt422016.AcquireByteBuffer()
//line view/layout.qtpl:27
	WriteLayout(qb422016, title, content)
//line view/layout.qtpl:27
	qs422016 := string(qb422016.B)
//line view/layout.qtpl:27
	qt422016.ReleaseByteBuffer(qb422016)
//line view/layout.qtpl:27
	return qs422016
//line view/layout.qtpl:27
}
