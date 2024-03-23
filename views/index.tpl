{{ define "index.tpl"}}
    {{ template "layouts/header.tpl" .}}
      <div class="container mt-4">
        <h1>{{.Title}}</h1>
      </div>
    {{ template "layouts/footer.tpl" .}}
{{ end }}