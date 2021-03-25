package mime

import (
	"io"
	"strings"

	"github.com/jhillyerd/enmime"
)

func ExtractContentFromEmail(reader io.Reader) ( /*subject*/ string /*body*/, string, error) {
	envelope, err := enmime.ReadEnvelope(reader)
	if err != nil {
		return "", "", "", err
	}

	sender := envelope.GetHeader("From")
	subject := envelope.GetHeader("Subject")
	body := strings.TrimSpace(envelope.Text)
	return sender, subject, body, nil
}
