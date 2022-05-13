package iris

import (
	"os"
	"reflect"
	"testing"

	"github.com/kataras/golog"
)

func Test_newLogger(t *testing.T) {
	defaultLogger := golog.Default

	type args struct {
		app     *Application
		appName string
	}
	tests := []struct {
		name string
		args args
		want *golog.Logger
	}{
		// test cases
		{
			"create logger with out IRIS_APP_NAME set",
			args{
				app:     &Application{},
				appName: "",
			}, defaultLogger.Clone()},
		{
			"create logger with IRIS_APP_NAME set",
			args{
				app:     &Application{},
				appName: "app",
			},
			defaultLogger.Clone(),
		},
	}
	for _, tt := range tests {
		tt.want.Child(tt.args.app)
		tt.want.SetChildPrefix(tt.args.appName)

		t.Run(tt.name, func(t *testing.T) {
			os.Setenv("IRIS_APP_NAME", tt.args.appName)
			got := newLogger(tt.args.app)

			if !reflect.DeepEqual(got.Prefix, tt.want.Prefix) {
				t.Errorf("Expected: %+v, got: %+v", got.Prefix, tt.want.Prefix)
			}

			if !reflect.DeepEqual(tt.args.app.name, tt.args.appName) {
				t.Errorf("Expected: %+v, got: %+v", tt.args.appName, tt.args.app.name)
			}
		})
	}
}

func TestApplication_SetName(t *testing.T) {
	type args struct {
		appName string
	}
	tests := []struct {
		name string
		app  *Application
		args args
		want *Application
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.app.SetName(tt.args.appName); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Application.SetName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestApplication_String(t *testing.T) {
	tests := []struct {
		name string
		app  *Application
		want string
	}{
		{"app with no name", &Application{}, ""},
		{"app with name 'iris'", &Application{name: "iris"}, "iris"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.app.String(); got != tt.want {
				t.Errorf("Application.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
