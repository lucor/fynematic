package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"text/template"
	"time"
)

const baseURL = "https://raw.githubusercontent.com/google/material-design-icons/%s/src/%s/%s/materialicons%s/24px.svg"

var styles = []string{
	"outlined", "filled", "round", "sharp", "twotone",
}

var categories = []string{
	"action", "alert", "av", "communication", "content", "device", "editor", "file",
	"hardware", "home", "image", "maps", "navigation", "notification", "places", "search",
	"social", "toggle",
}

func main() {
	var style string
	var category string
	var name string
	var pkg string
	var dir string
	var release string

	helpStyle := fmt.Sprintf("icon style. Allowed styles: %v", styles)
	helpCategories := fmt.Sprintf("icon category. Allowed categories: %v", categories)

	flag.StringVar(&category, "category", "action", helpCategories)
	flag.StringVar(&dir, "dir", ".", "output folder")
	flag.StringVar(&name, "name", "info", "icon name. See https://fonts.google.com/icons for the full list")
	flag.StringVar(&pkg, "package", "main", "package to use in header")
	flag.StringVar(&style, "style", "outlined", helpStyle)
	flag.StringVar(&release, "release", "master", "material icon release tag on GitHub")
	flag.Parse()

	if !contains(styles, style) {
		fmt.Printf("Invalid style %q. Allowed: %v\n", style, styles)
		os.Exit(1)
	}

	if !contains(categories, category) {
		fmt.Printf("Invalid category %q. Allowed: %v\n", category, categories)
		os.Exit(1)
	}

	if pkg == "" {
		fmt.Printf("package is required\n")
		os.Exit(1)
	}

	err := os.MkdirAll(dir, 0755)
	if err != nil {
		fmt.Printf("could not create or access the output directory %q: %v\n", dir, err)
		os.Exit(1)
	}

	reqStyle := ""
	if style != "filled" {
		reqStyle = style
	}
	url := fmt.Sprintf(baseURL, release, category, name, reqStyle)
	client := http.DefaultClient
	client.Timeout = 10 * time.Second

	resp, err := client.Get(url)
	if err != nil {
		fmt.Println("could not download icon from the GitHub repo:", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("could not read response from the GitHub repo:", err)
		os.Exit(1)
	}

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("could not download icon from the GitHub repo: %s\n", body)
		os.Exit(1)
	}

	iconLight := string(body)
	iconDark := strings.Replace(iconLight, "<svg ", "<svg fill=\"#ffffff\" ", 1)

	data := tplData{
		Name:      name,
		Style:     style,
		VarName:   makeVarName([]string{name, style}),
		Package:   pkg,
		IconDark:  strconv.Quote(iconDark),
		IconLight: strconv.Quote(iconLight),
		outputDir: dir,
	}

	err = makeThemedResourceFile(data)
	if err != nil {
		fmt.Printf("could not generate the theme resource file: %v\n", err)
		os.Exit(1)
	}

	err = makeIconFile(data)
	if err != nil {
		fmt.Printf("could not generate the icon file: %v\n", err)
		os.Exit(1)
	}
}

type tplData struct {
	Name      string
	Style     string
	VarName   string
	Package   string
	IconDark  string
	IconLight string
	outputDir string
}

func contains(slice []string, value string) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}

func makeIconFile(data tplData) error {
	filename := fmt.Sprintf("%s_%s.go", data.Name, data.Style)
	out := filepath.Join(data.outputDir, filename)
	w, err := os.Create(out)
	if err != nil {
		return err
	}
	defer w.Close()

	t := template.Must(template.New("icons").Parse(bundleTpl))
	return t.Execute(w, data)
}

func makeThemedResourceFile(data tplData) error {
	out := filepath.Join(data.outputDir, "themed_resource.go")
	w, err := os.Create(out)
	if err != nil {
		return err
	}
	defer w.Close()

	t := template.Must(template.New("themedResource").Parse(themedResourceTpl))
	return t.Execute(w, data)
}

func makeVarName(elems []string) string {
	var out []string
	for _, el := range elems {
		parts := strings.Split(el, "_")
		for _, part := range parts {
			out = append(out, strings.Title(part))
		}
	}
	return strings.Join(out, "")
}
