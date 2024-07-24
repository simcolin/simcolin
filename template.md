<h1 align="center">Hi 👋, I'm Simon</h1>
<h3 align="center">Frontend developer for work, game developer as a hobby</h3>

Tech used recently
<div align="center">
{{ range .UsedTech }}<img src="https://img.shields.io/badge/{{ urlquery .Name }}-{{ .Color }}?style=for-the-badge&logo={{ .Slug }}&logoColor={{ if .Black }}black{{ else }}white{{ end }}" alt="{{ .Name }}"/>
{{ end }}
</div>

Langages used recently
<div align="center">
{{ range .UsedLang }}<img src="https://img.shields.io/badge/{{ urlquery .Name }}-{{ .Color }}?style=for-the-badge&logo={{ .Slug }}&logoColor={{ if .Black }}black{{ else }}white{{ end }}" alt="{{ .Name }}"/>
{{ end }}
</div>

Known tech (used in the past)
<div align="center">
{{ range .KnownTech }}<img src="https://img.shields.io/badge/{{ urlquery .Name }}-{{ .Color }}?style=for-the-badge&logo={{ .Slug }}&logoColor={{ if .Black }}black{{ else }}white{{ end }}" alt="{{ .Name }}"/>
{{ end }}
</div>

Known langages
<div align="center">
{{ range .KnownLang }}<img src="https://img.shields.io/badge/{{ urlquery .Name }}-{{ .Color }}?style=for-the-badge&logo={{ .Slug }}&logoColor={{ if .Black }}black{{ else }}white{{ end }}" alt="{{ .Name }}"/>
{{ end }}
</div>

Tech I look forward to learning
<div align="center">
{{ range .WtlTech }}<img src="https://img.shields.io/badge/{{ urlquery .Name }}-{{ .Color }}?style=for-the-badge&logo={{ .Slug }}&logoColor={{ if .Black }}black{{ else }}white{{ end }}" alt="{{ .Name }}"/>
{{ end }}
</div>

Langages I look forward to learning
<div align="center">
{{ range .WtlLang }}<img src="https://img.shields.io/badge/{{ urlquery .Name }}-{{ .Color }}?style=for-the-badge&logo={{ .Slug }}&logoColor={{ if .Black }}black{{ else }}white{{ end }}" alt="{{ .Name }}"/>
{{ end }}
</div>