package brainfuck

import (
	"errors"
	"io"
	"reflect"
	"strings"
	"testing"
)

func TestInterpret(t *testing.T) {
	type args struct {
		stream io.Reader
	}
	tests := []struct {
		name        string
		args        args
		decodedText string
		wantErr     error
	}{
		{
			name: "HelloWorld",
			args: args{
				stream: strings.NewReader("HelloWorld"),
			},
			decodedText: "",
			wantErr:     nil,
		},
		{
			name: "Empty",
			args: args{
				stream: nil,
			},
			decodedText: "",
			wantErr:     errors.New("nil IO reader"),
		},
		{
			name: "Brainfuck wiki example Hello World!",
			args: args{
				stream: strings.NewReader("++++++++[>++++[>++>+++>+++>+<<<<-]>+>+>->>+[<]<-]>>.>---.+++++++..+++.>>.<-.<.+++.------.--------.>>+.>++."),
			},
			decodedText: " Hello World!",
			wantErr:     nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			decodeText, err := Interpret(tt.args.stream)
			if !reflect.DeepEqual(err, tt.wantErr) {
				t.Errorf("Test Interpret error got = %v, want %v", err, tt.wantErr)
			}
			if strings.Compare(decodeText, tt.decodedText) == -1 {
				t.Errorf("Test Interpret got  = %v, want %v", decodeText, tt.decodedText)
			}
		})
	}
}
