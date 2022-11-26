SET @schema = '{
  "type": "object",
  "properties": {
    "prefecture_id": {
      "type": "integer",
        "minimum": 1,
        "maximum": 47
     }
   },
   "required": [
     "prefecture_id"
   ]
}';

select json_schema_valid(@schema, '{"prefecture_id": 14}');
