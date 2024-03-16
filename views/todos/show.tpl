{{ define "todos/show.tpl"}}
    {{ template "layouts/header.tpl" . }}

    <div class="card">
        <div class="card-body">
            <h2 class="card-title">{{ .todo.Title }}</h2>
            <p class="card-text">{{ .todo.Detail }}</p>
            <div class="btn-group" role="group" aria-label="Todo Actions">
                <a href="/todos/edit/{{ .todo.Id }}" class="btn btn-primary" role="button">Edit</a>
                <a class="btn btn-danger" role="button" onclick="deleteTodo()">Delete</a>
            </div>
            <div class="btn-group" role="group" aria-label="Todos Nav">
                <a href="/todos/new" class="btn btn-primary" role="button">New TODO</a>
                <a href="/todos" class="btn btn-info" role="button">TODOs List</a>
            </div>
        </div>
    </div>
    
    <script>
        function deleteTodo() {
        var xhr = new XMLHttpRequest();
        xhr.open("DELETE", "/todos/{{.todo.Id}}", true);
        xhr.onload = function() {
            if (xhr.status === 200) {
            // Redirect the page
            window.location.replace("/todos");
            }
        };
        xhr.send();
        }
    </script>

    {{ template "layouts/footer.tpl" . }}
{{end}}