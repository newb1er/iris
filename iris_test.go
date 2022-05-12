package iris

import (
	"reflect"
	"testing"

	"github.com/kataras/golog"
)

func Test_newLogger(t *testing.T) {
	type args struct {
		app *Application
	}
	tests := []struct {
		name string
		args args
		want *golog.Logger
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newLogger(tt.args.app); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newLogger() = %v, want %v", got, tt.want)
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
		// TODO: Add test cases.
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
