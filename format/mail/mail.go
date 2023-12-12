package mail

// https://datatracker.ietf.org/doc/html/rfc5322#section-3.5
// https://datatracker.ietf.org/doc/html/rfc6532
// https://github.com/golang/go/blob/master/src/net/mail/message.go

import (
	"embed"

	"github.com/wader/fq/format"
	"github.com/wader/fq/pkg/decode"
	"github.com/wader/fq/pkg/interp"
	// "net/mail"
)

//go:embed mail.md
var mailFS embed.FS

func init() {
	interp.RegisterFormat(
		format.EML,
		&decode.Format{
			Description: "Electronic Mail Format",
			DecodeFn:    mailDecode,
		})
	interp.RegisterFS(mailFS)
}

func mailDecode(d *decode.D) any {
	return nil
}
