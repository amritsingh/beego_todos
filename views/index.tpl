{{ define "index.tpl"}}
    {{ template "layouts/header.tpl" .}}
      <div class="container mt-4">
        <h1>{{.Title}}</h1>
        <h2 class="mb-4">
          <a href="/todos" class="mb-4">TODO List</a>
        </h2>
      </div>
    {{ template "layouts/footer.tpl" .}}
{{ end }}