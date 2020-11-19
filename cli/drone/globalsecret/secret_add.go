package globalsecret

import (
	"io/ioutil"
	"strings"

	"github.com/laszlocph/woodpecker/cli/drone/internal"
	"github.com/laszlocph/woodpecker/drone-go/drone"

	"github.com/urfave/cli"
)

var globalSecretCreateCmd = cli.Command{
	Name:   "add",
	Usage:  "adds a global secret",
	Action: globalSecretCreate,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "name",
			Usage: "global secret name",
		},
		cli.StringFlag{
			Name:  "value",
			Usage: "global secret value",
		},
		cli.StringSliceFlag{
			Name:  "event",
			Usage: "global secret limited to these events",
		},
		cli.StringSliceFlag{
			Name:  "image",
			Usage: "global secret limited to these images",
		},
	},
}

func globalSecretCreate(c *cli.Context) error {
	client, err := internal.NewClient(c)
	if err != nil {
		return err
	}
	secret := &drone.Secret{
		Name:   c.String("name"),
		Value:  c.String("value"),
		Images: c.StringSlice("image"),
		Events: c.StringSlice("event"),
	}
	if len(secret.Events) == 0 {
		secret.Events = defaultSecretEvents
	}
	if strings.HasPrefix(secret.Value, "@") {
		path := strings.TrimPrefix(secret.Value, "@")
		out, ferr := ioutil.ReadFile(path)
		if ferr != nil {
			return ferr
		}
		secret.Value = string(out)
	}
	_, err = client.GlobalSecretCreate(secret)
	return err
}

var defaultSecretEvents = []string{
	drone.EventPush,
	drone.EventTag,
	drone.EventDeploy,
}
