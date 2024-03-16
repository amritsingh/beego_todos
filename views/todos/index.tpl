{{ define "todos/index.tpl"}}
    {{ template "layouts/header.tpl" .}}
        <h1>Todos</h1>

        <br/>

        <script type="text/javascript">
            function todoStateChange(element) {
                var xhttp = new XMLHttpRequest();
                id = parseInt(element.id.split("-")[1])
                xhttp.onreadystatechange = function() {};
                xhttp.open("POST", "/todos/" + id + "/complete", true);
                params = {
                    id: id,
                    state: element.checked
                }
                xhttp.setRequestHeader("Accept", "application/json")
                xhttp.setRequestHeader("Content-Type", "application/json")
                xhttp.send(JSON.stringify(params))
            }
        </script>

        <a class="btn btn-outline-primary" href="/todos/new" role="button">New Todo</a>

        <br/>
        <br/>
        
        <div class="accordion" id="accordionTodos">
            {{ with .todos }}
                {{ range . }}
                    <div class="accordion-item">
                        <h2 class="accordion-header" 
                            id="panelsStayOpen-heading-{{.Id}}" 
                            data-bs-toggle="collapse" 
                            data-bs-target="#panelsStayOpen-collapse-{{.Id}}" 
                            aria-expanded="false" aria-controls="panelsStayOpen-collapse-{{.Id}}">
                            <div class="flex-parent-element-{{.Id}} accordion-button collapsed">
                                <input class="flex-child-element-{{.Id}} form-check-input me-1" 
                                    type="checkbox" value="" 
                                    id="checkbox-{{.Id}}" 
                                    name="checkbox-{{.Id}}" onClick="todoStateChange(this)"
                                   {{if .CompletedAt}}
                                        checked
                                   {{end}}
                                ><strong>{{.Title}}</strong></input>
                            </div>
                        </h2>
                        <div id="panelsStayOpen-collapse-{{.Id}}" 
                            class="accordion-collapse collapse" 
                            aria-labelledby="panelsStayOpen-heading-{{.Id}}">
                            <div class="accordion-body">
                                {{.Detail}}
                                <a class="btn btn-outline-primary" href="/todos/{{.Id}}" role="button">View</a>
                            </div>
                        </div>
                    </div>
                {{end}}
            {{end}}
        </div>
    {{ template "layouts/footer.tpl" .}}
{{ end }}