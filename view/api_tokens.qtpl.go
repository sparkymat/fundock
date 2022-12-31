// Code generated by qtc from "api_tokens.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line view/api_tokens.qtpl:1
package view

//line view/api_tokens.qtpl:1
import "github.com/sparkymat/fundock/presenter"

//line view/api_tokens.qtpl:3
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line view/api_tokens.qtpl:3
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line view/api_tokens.qtpl:3
func StreamAPITokens(qw422016 *qt422016.Writer, csrfToken string, apiTokens []presenter.APIToken) {
//line view/api_tokens.qtpl:3
	qw422016.N().S(`
  <div class="uk-padding">
    <form action="/api_tokens" method="POST">
      <input type="hidden" name="csrf" value="`)
//line view/api_tokens.qtpl:6
	qw422016.E().S(csrfToken)
//line view/api_tokens.qtpl:6
	qw422016.N().S(`">
      <div class="uk-flex uk-flex-row">
        <div class="uk-form-controls uk-flex-1">
          <input class="uk-input" id="client_name" type="text" name="client_name" placeholder="Client name (e.g. iOS app)" required>
        </div>
        <input type="submit" value="Create API token" class="uk-button uk-button-primary uk-margin-left">
      </div>
    </form>
    <table class="uk-table uk-table-striped">
      <thead>
        <tr>
          <th>Client Name</th>
          <th>Token</th>
          <th>Last used</th>
        </tr>
      </thead>
      <tbody>
        `)
//line view/api_tokens.qtpl:23
	for _, apiToken := range apiTokens {
//line view/api_tokens.qtpl:23
		qw422016.N().S(`
          <tr>
            <td>`)
//line view/api_tokens.qtpl:25
		qw422016.E().S(apiToken.ClientName)
//line view/api_tokens.qtpl:25
		qw422016.N().S(`</td>
            <td>`)
//line view/api_tokens.qtpl:26
		qw422016.E().S(apiToken.Token)
//line view/api_tokens.qtpl:26
		qw422016.N().S(`</td>
            `)
//line view/api_tokens.qtpl:27
		if apiToken.LastUsed == nil {
//line view/api_tokens.qtpl:27
			qw422016.N().S(`
              <td><i>Never</i></td>
            `)
//line view/api_tokens.qtpl:29
		} else {
//line view/api_tokens.qtpl:29
			qw422016.N().S(`
              <td>`)
//line view/api_tokens.qtpl:30
			qw422016.E().S(*apiToken.LastUsed)
//line view/api_tokens.qtpl:30
			qw422016.N().S(`</td>
            `)
//line view/api_tokens.qtpl:31
		}
//line view/api_tokens.qtpl:31
		qw422016.N().S(`
          </tr>
        `)
//line view/api_tokens.qtpl:33
	}
//line view/api_tokens.qtpl:33
	qw422016.N().S(`
      </tbody>
    </table>
  </div>
`)
//line view/api_tokens.qtpl:37
}

//line view/api_tokens.qtpl:37
func WriteAPITokens(qq422016 qtio422016.Writer, csrfToken string, apiTokens []presenter.APIToken) {
//line view/api_tokens.qtpl:37
	qw422016 := qt422016.AcquireWriter(qq422016)
//line view/api_tokens.qtpl:37
	StreamAPITokens(qw422016, csrfToken, apiTokens)
//line view/api_tokens.qtpl:37
	qt422016.ReleaseWriter(qw422016)
//line view/api_tokens.qtpl:37
}

//line view/api_tokens.qtpl:37
func APITokens(csrfToken string, apiTokens []presenter.APIToken) string {
//line view/api_tokens.qtpl:37
	qb422016 := qt422016.AcquireByteBuffer()
//line view/api_tokens.qtpl:37
	WriteAPITokens(qb422016, csrfToken, apiTokens)
//line view/api_tokens.qtpl:37
	qs422016 := string(qb422016.B)
//line view/api_tokens.qtpl:37
	qt422016.ReleaseByteBuffer(qb422016)
//line view/api_tokens.qtpl:37
	return qs422016
//line view/api_tokens.qtpl:37
}