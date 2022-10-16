package ROT13

import (
	"strings"
)

func indexposition(s string, arr []string) int {
	for i, c := range arr {
		if c == s {
			return i
		}
	}
	return -1
}

func decode(answer string) string {

	alphabet := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	alphalen := len(alphabet)

	var message []string

	for _, c := range alphabet {
		strC := string(c)
		if strC == " " {
			message = append(message, " ")
			continue
		}

		isUpper := false
		if strings.ToLower(strC) != strC {
			isUpper = true
			continue
		}

		i := indexposition(strC, alphabet)

		if i == -1 {
			message = append(message, strC)
			continue
		}

		var pos int
		pos = i - 13

		if pos < 0 {
			pos = alphalen + (i - 13)
		}

		if isUpper {
			message = append(message, strings.ToUpper(alphabet[pos]))
			continue
		}
		message = append(message, alphabet[pos])
	}
	return strings.Join(message, "")
}

func encode(message string) string {
	return ""
}

/*func main() {
	conn, err := net.Dial("IRC", "irc.root-me.org:6667")
	if err != nil {
		log.Fatalln(err)
	}

	config := irc.ClientConfig{
		Nick: "i_have_a_nick",
		Pass: "password",
		User: "username",
		Name: "Full Name",
		Handler: irc.HandlerFunc(func(c *irc.Client, m *irc.Message) {
			if m.Command == "001" {
				// 001 is a welcome event, so we join channels there
				c.Write("JOIN #root-me_challenge")
			} else if m.Command == "!ep3" && m.FromChannel() {
				// Create a handler on all messages.
				c.WriteMessage(&irc.Message{
					Command: "!ep3",
					Params: []string{
						m.Params[0],
						m.Trailing(),
					},
				})
			}
		}),
	}

	// Create the client
	client := irc.NewClient(conn, config)
	err = client.Run()
	if err != nil {
		log.Fatalln(err)
	}
}*/
