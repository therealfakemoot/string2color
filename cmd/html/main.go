package main

import (
	"bufio"
	"flag"
	"fmt"
	"html/template"
	"log/slog"
	"os"
	"regexp"
	"strings"

	"github.com/therealfakemoot/string2color"
)

var (
	pat          = regexp.MustCompile(`[^\p{L}\p{Z}]+`)
	wordTemplate = template.Must(template.New("word").Parse(`<span style="{{ .Color }}">{{ .Word}}</span>{{ .Punc }} `))
)

func main() {
	var input string

	flag.StringVar(&input, "input", "input.txt", "Input corpus.")

	flag.Parse()

	var err error

	corpus, err := os.Open(input)
	if err != nil {
		slog.With("error", err).Error("could not open corpus")
		return
	}

	s := bufio.NewScanner(corpus)
	s.Split(bufio.ScanWords)

	var builder strings.Builder
	builder.WriteString("<p>")

	for s.Scan() {
		err = s.Err()
		if err != nil {
			break
		}
		t := s.Text()
		punctuation := pat.FindAllString(t, -1)
		var word string
		allPunc := strings.Join(punctuation, "")
		word = t[:len(t)-len(allPunc)]
		r, g, b, _ := string2color.ToRGB(word).RGBA()
		slog.With("word", word, "r", r/256, "b", b/256, "g", g/256).Info("color values")
		if len(allPunc) > 0 {
			slog.With("punctuation", allPunc).Info("punctuation")
		}
		color := fmt.Sprintf("color:#%02X%02X%02X;", r/256, g/256, b/256)
		slog.With("colorHex", color).Info("color hex code")
		data := struct {
			Word  string
			Color template.CSS
			Punc  string
		}{
			Word:  word,
			Punc:  allPunc,
			Color: template.CSS(color),
		}
		err := wordTemplate.Execute(&builder, data)
		if err != nil {
			slog.With("error", err).Error("error executing word template")
			return
		}
	}

	builder.WriteString("</p>")
	fmt.Println(builder.String())

	if err != nil {
		slog.With("error", err).Error("scanner returned error")
		return
	}
}
