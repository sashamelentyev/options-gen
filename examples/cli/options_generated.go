// Code generated by options-gen. DO NOT EDIT.
package cli

import (
	"fmt"
	"net/http"

	goplvalidator "github.com/go-playground/validator/v10"
	uniqprefixformultierror "github.com/hashicorp/go-multierror"
)

var _validator461e464ebed9 = goplvalidator.New()

type optOptionsMeta struct {
	setter    func(o *Options)
	validator func(o *Options) error
}

func NewOptions(
	httpClient *http.Client,
	token string,

	options ...optOptionsMeta,
) Options {
	o := Options{}
	o.httpClient = httpClient
	o.token = token

	for i := range options {
		options[i].setter(&o)
	}

	return o
}

func (o *Options) Validate() error {
	var g uniqprefixformultierror.Group

	g.Go(func() error {
		err := _Options_httpClientValidator(o)
		if err != nil {
			return fmt.Errorf("invalid value for option WithHttpClient: %w", err)
		}
		return nil
	})
	g.Go(func() error {
		err := _Options_tokenValidator(o)
		if err != nil {
			return fmt.Errorf("invalid value for option WithToken: %w", err)
		}
		return nil
	})
	return g.Wait()
}

func _Options_httpClientValidator(o *Options) error {

	if err := _validator461e464ebed9.Var(o.httpClient, "required"); err != nil {
		return fmt.Errorf("field `httpClient` did not pass the test: %w", err)
	}

	return nil
}

func _Options_tokenValidator(o *Options) error {

	return nil
}
