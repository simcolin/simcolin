package main

import (
	"encoding/base64"
	"os"
	"text/template"
)

type Tech struct {
	Name  string
	Slug  string
	Icon  string
	Color string
	Black bool
}

type Knownledge struct {
	UsedTech  []Tech
	UsedLang  []Tech
	KnownTech []Tech
	KnownLang []Tech
	WtlTech   []Tech
	WtlLang   []Tech
}

func svgToDataUri(svg string) string {
	return "data:image/svg+xml;base64," + base64.StdEncoding.EncodeToString([]byte(svg))
}

func missingLogo(filename string) string {
	logoBytes, err := os.ReadFile(filename)
	if err != nil {
		return ""
	}
	return svgToDataUri(string(logoBytes))
}

func main() {
	knowledge := Knownledge{
		UsedTech: []Tech{
			{Name: "Svelte", Slug: "svelte", Color: "FF3E00"},
			{Name: "pnpm", Slug: "pnpm", Color: "F69220"},
			{Name: "Git", Slug: "git", Color: "F05032"},
			{Name: "Visual_Studio_Code", Slug: missingLogo("missing_logos/vscode.svg"), Color: "007ACC"},
			{Name: "Node.js", Slug: "nodedotjs", Color: "339933"},
			{Name: "Turso", Slug: "turso", Color: "4FF8D2", Black: true},
			{Name: "SQLite", Slug: "sqlite", Color: "003B57"},
			{Name: "Fastify", Slug: "fastify", Color: "000000"},
			{Name: "Hono", Slug: "hono", Color: "E36002"},
			{Name: "Tailwind_CSS", Slug: "tailwindcss", Color: "06B6D4"},
			{Name: "Post_CSS", Slug: "postcss", Color: "DD3A0A"},
			{Name: "Godot_Engine", Slug: "godotengine", Color: "478CBF"},
			{Name: "Steamworks", Slug: "steamworks", Color: "1E1E1E"},
			{Name: "Bun", Slug: "bun", Color: "14151A"},
			{Name: "Cloudflare_Pages", Slug: "cloudflarepages", Color: "F38020"},
		},
		UsedLang: []Tech{
			{Name: "Typescript", Slug: "typescript", Color: "3178C6"},
			{Name: "Javascript", Slug: "javascript", Color: "F7DF1E", Black: true},
			{Name: "CSS", Slug: "css3", Color: "1572B6"},
			{Name: "HTML", Slug: "html5", Color: "E34F26"},
			{Name: "GDScript", Slug: "godotengine", Color: "478CBF"},
			{Name: "Go", Slug: "go", Color: "00ADD8"},
			{Name: "Odin", Slug: missingLogo("missing_logos/odin.svg"), Color: "1E5085"},
		},
		KnownTech: []Tech{
			{Name: "npm", Slug: "npm", Color: "CB3837"},
			{Name: "Angular", Slug: "angular", Color: "0F0F11"},
			{Name: "Firebase", Slug: "firebase", Color: "FFCA28", Black: true},
			{Name: "ChartJS", Slug: "chartdotjs", Color: "FF6384"},
			{Name: "MongoDB", Slug: "chartdotjs", Color: "47A248"},
			{Name: "ChartJS", Slug: "chartdotjs", Color: "FF6384"},
			{Name: "Processing", Slug: "processingfoundation", Color: "006699"},
			{Name: "p5.js", Slug: "p5dotjs", Color: "ED225D"},
			{Name: "Unity_Engine", Slug: "unity", Color: "FFFFFF", Black: true},
			{Name: "Express", Slug: "express", Color: "000000"},
			{Name: "SCSS", Slug: "sass", Color: "CC6699"},
		},
		KnownLang: []Tech{
			{Name: "Lua", Slug: "lua", Color: "2C2D72"},
			{Name: "CPP", Slug: "cplusplus", Color: "00599C"},
			{Name: "Java", Slug: missingLogo("missing_logos/java.svg"), Color: "5382a1"},
			{Name: "C#", Slug: missingLogo("missing_logos/csharp.svg"), Color: "9b5297"},
			{Name: "Python", Slug: "python", Color: "3776AB"},
			{Name: "PHP", Slug: "php", Color: "777BB4"},
		},
		WtlTech: []Tech{
			{Name: "HTMX", Slug: "htmx", Color: "3366CC"},
			{Name: "Vue.js", Slug: "vuedotjs", Color: "4FC08D"},
		},
		WtlLang: []Tech{
			{Name: "Kotlin", Slug: "kotlin", Color: "7F52FF"},
			{Name: "Zig", Slug: "zig", Color: "F7A41D"},
			{Name: "Swift", Slug: "swift", Color: "F05138"},
			{Name: "Jai", Slug: "jai", Color: "3178C6"},
		},
	}

	templateBytes, err := os.ReadFile("template.md")
	if err != nil {
		panic(err)
	}

	tmpl, err := template.New("README").Parse(string(templateBytes))
	if err != nil {
		panic(err)
	}

	readme, err := os.OpenFile("README.md", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		panic(err)
	}
	defer readme.Close()

	err = tmpl.Execute(readme, knowledge)
	if err != nil {
		panic(err)
	}
}
