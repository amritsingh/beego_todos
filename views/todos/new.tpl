{{ define "todos/new.tpl"}}
    {{ template "layouts/header.tpl" .}}
        <h2>Create a new TODO</h2>
        <form class="row g-3" action="/todos" method="POST">
            <div class="mb-3">
                <label for="title" class="form-label">Title</label>
                <input type="text" class="form-control" id="title" 
                    aria-describedby="titleHelp" name="title">
                <div id="titleHelp" class="form-text">Title of the TODO Item</div>
            </div>
            <div class="mb-3">
                <label for="detail" class="form-label">Detail</label>
                <textarea class="form-control" id="detail" rows="5" name="detail"></textarea>
            </div>
            <div class="mx-auto">
                <button type="submit" class="btn btn-primary mx-auto">Submit</button>
            </div>
        </form>
        <br/>
        <a href="/todos" class="btn btn-info" role="button">TODO List</a>
    {{ template "layouts/footer.tpl" .}}
{{ end }}