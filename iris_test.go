package iris

import (
	"io"
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
		originEnv := os.Getenv("IRIS_APP_NAME")

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

		os.Setenv("IRIS_APP_NAME", originEnv)
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

func Test_customHostServerLogger_Write(t *testing.T) {
	type fields struct {
		parent     io.Writer
		ignoreLogs [][]byte
	}
	type args struct {
		p []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
	}{
		{"l.ignoreLogs is not empty", fields{os.Stdout, [][]byte{[]byte("abc"), []byte("123")}}, args{[]byte("123\n")}, 0, false},
		{"l.ignoreLogs is empty", fields{os.Stdout, [][]byte{}}, args{[]byte("123\n")}, 4, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &customHostServerLogger{
				parent:     tt.fields.parent,
				ignoreLogs: tt.fields.ignoreLogs,
			}
			got, err := l.Write(tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("customHostServerLogger.Write() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("customHostServerLogger.Write() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestApplication_Configure(t *testing.T) {
	tests := []struct {
		name   string
		fields *Application
		args   []Configurator
		want   *Application
	}{
		{
			"configure with no configurators",
			&Application{},
			[]Configurator{},
			&Application{},
		},
		{
			"configure with nil and non-nil configurators",
			&Application{builded: false},
			[]Configurator{nil, func(a *Application) { a.builded = true }},
			&Application{builded: true},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app := &Application{
				builded: tt.fields.builded,
			}
			if got := app.Configure(tt.args...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Application.Configure() = %v, want %v", got, tt.want)
			}
		})
	}
}
