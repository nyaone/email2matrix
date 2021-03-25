package matrix

import "fmt"

func GenerateMessage(
	subject string,
	sender string,
	body string,
	ignoreSubject, ignoreBody, skipMarkdown bool,
) string {
	if skipMarkdown {
		if ignoreBody || body == "" {
			if subject == "" {
				return ""
			}
			return fmt.Sprintf("%s\n%s", subject, sender)
		}

		if ignoreSubject || subject == "" {
			return body
		}

		return fmt.Sprintf("%s\n%s\n\n%s", subject, sender, body)
	}

	if ignoreBody || body == "" {
		if subject == "" {
			return ""
		}
		return fmt.Sprintf("# %s\n> %s", subject, sender)
	}

	if ignoreSubject || subject == "" {
		return body
	}

	return fmt.Sprintf("# %s\n> %s\n\n%s", subject, sender, body)
}
