package step

import (
	"fmt"
	"net/http"
)

type openLink struct {
	ctx Context
}


func (s *openLink) Execute(ctx Context) error {
	response, err := http.Get(ctx.Params()["url"]);
	if err != nil {
		return err;
	}

	fmt.Println(fmt.Sprintf("Opening url: %s status: %d", ctx.Params()["url"], response.StatusCode))
	return nil;
}

func (*openLink) Code() string {
	return "openLink"
}

func (*openLink) Name() string {
	return "Open Link"
}

func (*openLink) Description() string {
	return "open link to warm up cache of website"
}

func (*openLink) Params() map[string]string {
	params := map[string]string{}
	params["url"] = "link to open"

	return params
}
func (openlink *openLink) Validate(ctx Context) error {
	url := ctx.Params()["url"]

	if url == "" {
		return NewError(openlink.Name(), "url param must be provided")
	}

	return nil
}