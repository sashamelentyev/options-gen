package generator

import (
	"fmt"
	"testing"

	// test named imports.
	req "github.com/stretchr/testify/require"
)

const gofile = "generator_test.go"

func TestGetImports(t *testing.T) {
	imports, err := GetFileImports(gofile)
	req.NoError(t, err)
	requiredImports := []string{
		`"fmt"`,
		`"testing"`,
		`req "github.com/stretchr/testify/require"`,
	}
	req.Equal(t, requiredImports, imports)
}

func TestGetOptionSpec(t *testing.T) {
	data, err := GetOptionSpec(gofile, "TestOptions")
	req.NoError(t, err)
	req.Equal(t, []OptionMeta{
		{
			Name:  "Stringer",
			Field: "stringer",
			Type:  "fmt.Stringer",
			TagOption: TagOption{
				IsRequired: true,
				IsNotEmpty: true,
			},
		},
		{
			Name:  "Str",
			Field: "str",
			Type:  "string",
			TagOption: TagOption{
				IsRequired: false,
				IsNotEmpty: true,
			},
		},
		{
			Name:  "NoValidation",
			Field: "noValidation",
			Type:  "string",
			TagOption: TagOption{
				IsRequired: false,
				IsNotEmpty: false,
			},
		},
	}, data)
}

func TestRenderOptions(t *testing.T) {
	data, err := GetOptionSpec(gofile, "TestOptions")
	req.NoError(t, err)

	imports, err := GetFileImports(gofile)
	req.NoError(t, err)

	res, err := RenderOptions("generator", "TestOptions", imports, data)
	req.NoError(t, err)

	// NOTE(a.telyshev): При наличии ошибок компиляции в файле импорты могут плохо
	// мержиться, сортироваться и пр.
	req.Equal(t, testStructGenerated, string(res), "generated file:\n%v", string(res))
}

func TestToStopCIFromComplaining(t *testing.T) {
	s := TestOptions{
		stringer:     nil,
		str:          "123",
		noValidation: "qwe",
	}
	req.Equal(t, TestOptions{
		stringer:     nil,
		str:          "123",
		noValidation: "qwe",
	}, s)
}

type TestOptions struct {
	stringer     fmt.Stringer `option:"required,not-empty"`
	str          string       `option:"not-empty"`
	noValidation string
}

const testStructGenerated = `// Code generated by options-gen. DO NOT EDIT.
package generator

import (
	"fmt"

	"github.com/kazhuravlev/options-gen/validator"
)

type optTestOptionsMeta struct {
	setter    func(o *TestOptions)
	validator func(o *TestOptions) error
}

func _TestOptions_stringerValidator(o *TestOptions) error {
	if validator.IsNil(o.stringer) {
		return fmt.Errorf("%w: stringer must be set (type fmt.Stringer)", ErrInvalidOption)
	}
	return nil
}

func _TestOptions_strValidator(o *TestOptions) error {
	if validator.IsNil(o.str) {
		return fmt.Errorf("%w: str must be set (type string)", ErrInvalidOption)
	}
	return nil
}

func WithStr(opt string) optTestOptionsMeta {
	return optTestOptionsMeta{
		setter:    func(o *TestOptions) { o.str = opt },
		validator: _TestOptions_strValidator,
	}
}

func _TestOptions_noValidationValidator(o *TestOptions) error {

	return nil
}

func WithNoValidation(opt string) optTestOptionsMeta {
	return optTestOptionsMeta{
		setter:    func(o *TestOptions) { o.noValidation = opt },
		validator: _TestOptions_noValidationValidator,
	}
}

func NewTestOptions(
	stringer fmt.Stringer,

	options ...optTestOptionsMeta,
) TestOptions {
	o := TestOptions{}
	o.stringer = stringer

	for i := range options {
		options[i].setter(&o)
	}

	return o
}

func (o *TestOptions) Validate() error {
	if err := _TestOptions_stringerValidator(o); err != nil {
		return fmt.Errorf("%w: invalid value for option WithStringer", err)
	}

	if err := _TestOptions_strValidator(o); err != nil {
		return fmt.Errorf("%w: invalid value for option WithStr", err)
	}

	if err := _TestOptions_noValidationValidator(o); err != nil {
		return fmt.Errorf("%w: invalid value for option WithNoValidation", err)
	}

	return nil
}
`
