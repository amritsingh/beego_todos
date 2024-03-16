{{ define "todos/edit.tpl"}}
    {{ template "layouts/header.tpl" . }}
        <h1>Edit TODO</h1>

        <form action="/todos/{{ .todo.Id }}" method="POST">
            <div class="form-group">
                <label for="title">Title:</label>
                <input type="text" class="form-control" id="title" name="title" value="{{ .todo.Title }}">
            </div>

            <div class="form-group">
                <label for="detail">Content:</label>
                <textarea class="form-control" id="detail" name="detail" rows="10">{{ .todo.Detail }}</textarea>
            </div>

            <div class="mb-2"></div>

            <button type="submit" class="btn btn-primary">Submit</button>
        </form>

    {{ template "layouts/footer.tpl" . }}
{{ end }}
