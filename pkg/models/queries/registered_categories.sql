SELECT c1.id "id",
       c1.category_name "categoryName",
       c1.nomeclature "nomeclature",
       c1.parent_id "parentId",
       COALESCE(c1.multi_lang_values, '[]') "multiLangCategory",
       c1.multi_lang_is_active "multiLangIsActive",
       c2.category_name "parentName"
FROM {{schema}}.category c1
LEFT JOIN {{schema}}.category c2 ON c2.id = c1.parent_id
WHERE 
    (
        COALESCE('{{key}}', '') = ''
        OR c1.category_name ILIKE '%' || '{{key}}' || '%'
        OR c2.category_name ILIKE '%' || '{{key}}' || '%'
        OR c1.nomeclature ILIKE '%' || '{{key}}' || '%'
    );