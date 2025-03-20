WITH RECURSIVE category_path AS 
       (
           -- Nivel base
           SELECT c1.id AS "categoryId", 
                  c1.category_name AS "categoryName", 
                  c1.nomeclature As "nomeclature", 
                  c1.parent_id  AS "parentId",
                  c1.multi_lang_values AS "multiLangCategory", 
                  c1.multi_lang_is_active AS "multiLangIsActive"
           FROM {{schema}}.category c1
           WHERE c1.id = '{{categoryId}}'
   
           UNION ALL
   
           -- Recorre hacia arriba la jerarquía
           SELECT c2.id AS "categoryId", 
                  c2.category_name AS "categoryName", 
                  c2.nomeclature As "nomeclature", 
                  c2.parent_id  AS "parentId",
                  c2.multi_lang_values AS "multiLangCategory", 
                  c2.multi_lang_is_active AS "multiLangIsActive"
           FROM {{schema}}.category c2
           INNER JOIN category_path lp ON c2.id = "parentId"::uuid
       )
SELECT * FROM category_path;