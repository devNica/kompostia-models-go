SELECT c1.id,
       c1.category_name AS "categoryName",
       c1.nomeclature,
       c1.parent_id AS "parentId",
       c1.multi_lang_values AS "multiLangCategory",
       c1.multi_lang_is_active AS "multiLangIsActive",
       c2.category_name AS "parentName"
FROM {{schema}}.category c1
LEFT JOIN {{schema}}.category c2 ON c2.id = c1.parent_id
WHERE 
    (
        COALESCE('{{key}}', '') = ''
        OR c1.category_name ILIKE '%' || '{{key}}' || '%'
        OR c2.category_name ILIKE '%' || '{{key}}' || '%'
        OR c1.nomeclature ILIKE '%' || '{{key}}' || '%'
    );