SET @schema = '{
         "$schema": "http://json-schema.org/draft-07/schema#",
         "definitions": {
             "todo": {
                 "definitions": {
                     "id": {
                         "description": "TODOã®ID",
                         "type": "integer"
                     },
                     "naiyou": {
                         "description": "TODOã®åå®¹",
                         "type": "string"
                     }
                 }
             }
         },
         "type": "object",
         "properties": {
             "todos": {
                 "type": "array",
                 "items": {
                     "type": "object",
                     "properties": {
                         "id": {
                             "$ref": "#/definitions/todo/definitions/id"
                         },
                         "naiyou": {
                             "$ref": "#/definitions/todo/definitions/naiyou"
                         }
                     }
                 }
             }
         }
     }';


SELECT json_schema_valid(@schema, '{"todos":[{"id":1, "naiyou":"è²·ãç©"},{"id":2, "naiyou":"æ ç»"}]}') a;
