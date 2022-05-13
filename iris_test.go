package iris

import (
	"io"
	"os"
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
