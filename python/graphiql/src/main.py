from fastapi import FastAPI
from fastapi.responses import HTMLResponse

app = FastAPI()


@app.get("/graphiql/", response_class=HTMLResponse)
async def graphiql():
    return """
    <!DOCTYPE html>
    <html>
      <head>
        <style>
          body {
            height: 100%;
            margin: 0;
            width: 100%;
            overflow: hidden;
          }
          #graphiql {
            height: 100vh;
          }
          .jwt-token {
            background: linear-gradient(#f7f7f7, #e2e2e2);
            border-bottom: 1px solid #d0d0d0;
            font-family: system, -apple-system, 'San Francisco', '.SFNSDisplay-Regular', 'Segoe UI', Segoe, 'Segoe WP', 'Helvetica Neue', helvetica, 'Lucida Grande', arial, sans-serif;
            padding: 7px 14px 6px;
            font-size: 14px;
          }
        </style>
        <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/graphiql/0.10.2/graphiql.css" />
        <script src="https://cdnjs.cloudflare.com/ajax/libs/fetch/2.0.3/fetch.min.js"></script>
        <script src="https://cdnjs.cloudflare.com/ajax/libs/react/15.5.4/react.min.js"></script>
        <script src="https://cdnjs.cloudflare.com/ajax/libs/react/15.5.4/react-dom.min.js"></script>
        <script src="https://cdnjs.cloudflare.com/ajax/libs/graphiql/0.10.2/graphiql.js"></script>
      </head>
      <body>
        <div class="jwt-token">
          JWT Token <input id="jwt-token" placeholder="JWT Token goes here">
          GraphQL Url <input id="gql-url" placeholder="GraphQl Url">
        </div>
        <div id="graphiql">Loading...</div>
        <script>
        var search = window.location.search;
        var parameters = {};
        document.getElementById('jwt-token').value = localStorage.getItem('graphiql:jwtToken');
        document.getElementById('gql-url').value = localStorage.getItem('graphiql:graphqlUrl');
        function graphQLFetcher(graphQLParams) {
          const jwtToken = document.getElementById('jwt-token').value;
          const graphqlUrl = document.getElementById('gql-url').value;
          let headers = {
            'Accept': 'application/json',
            'Content-Type': 'application/json'
          };
          if (jwtToken) {
            localStorage.setItem('graphiql:jwtToken', jwtToken);
            headers = {
              'Accept': 'application/json',
              'Content-Type': 'application/json',
              'Authorization': jwtToken ? `Bearer ${jwtToken}` : null
            };
          }
          if (graphqlUrl) {
            localStorage.setItem('graphiql:graphqlUrl', graphqlUrl);
          }
          return fetch(`${graphqlUrl}`, {
            method: 'post',
            headers,
            body: JSON.stringify(graphQLParams)
          }).then(function (response) {
            return response.text();
          }).then(function (responseBody) {
            try {
              return JSON.parse(responseBody);
            } catch (error) {
              return responseBody;
            }
          });
        }
        // Render <GraphiQL /> into the body.
        // See the README in the top level of this module to learn more about
        // how you can customize GraphiQL by providing different values or
        // additional child elements.
        ReactDOM.render(
          React.createElement(GraphiQL, {
            fetcher: graphQLFetcher,
          }),
          document.getElementById('graphiql')
        );
        </script>
      </body>
    </html>
    """
